pragma solidity ^0.8.9;

import "./SupplyController.sol";

import "https://raw.githubusercontent.com/ninjaswtf/ERC721A/main/contracts/ERC721A.sol";
import "solmate/src/auth/Owned.sol";

contract FukushimaFishNFT is
    ERC721A("Fukushima Fish", "KOI"),
    Owned(msg.sender)
{
    enum MintStatus {
        // 0 = closed
        CLOSED,
        // 1 = open for whitelisted addresses
        WHITELIST,
        // 2 = open for the public
        PUBLIC
    }

    uint256 public WHITELIST_MINT_COST = 0.05 ether;
    // Subject to change
    uint256 public PUBLIC_MINT_COST = 0.0777 ether;

    uint256 constant MAX_PUBLIC_MINT_PER_WALLET = 20;

    uint256 constant MAX_SUPPLY = 3888;

    string constant NO_MINTS_REMAINING = "You have no mints remaining";

    // The default mint status is CLOSED
    MintStatus public mintStatus = MintStatus.CLOSED;

    string _baseTokenURI = "";
    string _unrevealedURI = "";

    string public termsOfServiceURI;
    string public readMeURI;

    bytes32 whitelistMerkleProofRoot = bytes32(0);

    bool _updateOnTransfer = false;


    SupplyController public controller;

    function setSupplyController(SupplyController _controller) external onlyOwner {
        controller = _controller;
    }

    function getMintTime(uint256 tokenId) external view returns (uint256) {
        TokenOwnership memory ownership = _ownershipOf(tokenId);
        return ownership.startTimestamp;
    }

    function setUpdateOnTransfer(bool status) external onlyOwner {
        _updateOnTransfer = status;
    }


    function _updateTimestampOnTransfer() internal virtual override returns(bool)  {
        return _updateOnTransfer;
    }

    function exists(uint256 id) external view returns (bool) {
        return _exists(id);
    }

    function minted(address _addr) external view returns (uint256) {
        return _numberMinted(_addr);
    }


    function setTermsOfServiceURI(string calldata uri) external onlyOwner {
        termsOfServiceURI = uri;
    }

    function setReadMeURI(string calldata uri) external onlyOwner {
        readMeURI = uri;
    }

    function setMintStatus(MintStatus status) external onlyOwner {
        mintStatus = status;
    }

    function setMerkleRoot(bytes32 merkleRoot) external onlyOwner {
        whitelistMerkleProofRoot = merkleRoot;
    }

    function setBaseTokenURI(string calldata uri) external onlyOwner {
        _baseTokenURI = uri;
    }

    function setUnrevealedTokenURI(string calldata uri) external onlyOwner {
        _unrevealedURI = uri;
    }

    function setWhitelistMintPrice(uint256 cost) external onlyOwner {
        WHITELIST_MINT_COST = cost;
    }

    function setPublicMintPrice(uint256 cost) external onlyOwner {
        PUBLIC_MINT_COST = cost;
    }

    function ownerMint(address to, uint256 amount) external onlyOwner {
        // supply limit checks
        require(_totalMinted() < MAX_SUPPLY, "minted out.");
        require(
            _totalMinted() + amount <= MAX_SUPPLY,
            "mint amount would be out of range."
        );
        _mint(to, amount);
    }

    // Validate checks if the given address and proof result in the merkle tree root.
    // if the proof & the hashed address resolves to the provided proof, then the address
    // is within the whitelist.
    function validate(
        address addr,
        uint256 limit,
        bytes32[] calldata proof,
        uint256 path
    ) public view returns (bool) {
        bytes32 hash = keccak256(abi.encode(addr, limit));

        for (uint256 i; i < proof.length; i++) {
            // check if the path is odd and inverse the hash
            if (path & 1 == 1) {
                hash = keccak256(abi.encodePacked(hash, proof[i]));
            } else {
                hash = keccak256(abi.encodePacked(proof[i], hash));
            }

            // this divides the path by 2 lol bitwise ops > division
            path >>= 1;
        }

        return hash == whitelistMerkleProofRoot;
    }

    function publicMint(uint256 amount) external payable {
        uint256 currentSupply = _totalMinted();
        // supply limit checks
        require(
            msg.sender == tx.origin &&
                amount > 0 &&
                amount <= MAX_PUBLIC_MINT_PER_WALLET &&
                currentSupply < MAX_SUPPLY &&
                currentSupply + amount <= MAX_SUPPLY &&
                mintStatus == MintStatus.PUBLIC
        );

        uint256 minimumPayment = amount * PUBLIC_MINT_COST;

        require(msg.value >= minimumPayment, "not enough ether sent for mint!");

        _mint(msg.sender, amount);

        if (msg.value > minimumPayment) {
            // refund if the user somehow overpaid
            uint256 refund = msg.value - minimumPayment;
            (bool ok, ) = payable(msg.sender).call{value: refund}("");
            require(ok);
        }
    }

    function whitelistMint(
        uint256 amount,
        uint256 limit,
        bytes32[] calldata proof,
        uint256 path
    ) external payable {
        address msgSender = msg.sender;

        uint256 currentSupply = _totalMinted();

        // supply & sanity checks
        require(
            amount > 0 &&
                amount <= limit &&
                currentSupply < MAX_SUPPLY &&
                currentSupply + amount <= MAX_SUPPLY &&
                msg.sender == tx.origin &&
                mintStatus != MintStatus.CLOSED
        );

        // account mint limit checks
        uint256 _minted = _numberMinted(msgSender);
        require(proof.length > 0 && validate(msgSender, limit, proof, path));
        require(_minted < limit, "Mint limit reached!");
        uint256 remaining = limit - _minted;
        require(amount > 0 && amount <= remaining);

        // payment checks

        uint256 minimumPayment = amount * WHITELIST_MINT_COST;

        require(msg.value >= minimumPayment, "not enough ether sent for mint!");

        _mint(msgSender, amount);

        if (msg.value > minimumPayment) {
            // refund if the user somehow overpaid
            uint256 refund = msg.value - minimumPayment;
            (bool ok, ) = payable(msgSender).call{value: refund}("");
            require(ok);
        }
    }

    function tokenURI(uint256 id) public view override returns (string memory) {
        // baseTokenURI is empty, assume the token is unrevealed, and default to the unrevealed URI
        // else concatenate the base URI with the token ID and the JSON URI
        return
            bytes(_baseTokenURI).length == 0
                ? _unrevealedURI
                : string(abi.encodePacked(_baseTokenURI, id, ".json"));
    }

    function _startTokenId() internal pure override returns (uint256) {
        return 1;
    }

    function withdraw(address to) external onlyOwner {
        (bool ok, ) = payable(to).call{value: address(this).balance}("");
        require(ok);
    }

    function _beforeTokenTransfers(
        address from,
        address to,
        uint256 startTokenId,
        uint256 quantity
    ) internal override {
        if (address(controller) != address(0)) {
            controller.onPreTransfer(from, to, startTokenId, quantity);
        }
    }

    function _afterTokenTransfers(
        address from,
        address to,
        uint256 startTokenId,
        uint256 quantity
    ) internal override {
        if (address(controller) != address(0)) {
            controller.onPostTransfer(from, to, startTokenId, quantity);
        }
    }
}
