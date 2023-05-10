pragma solidity ^0.8.9;

import "./SupplyController.sol";

import "erc721a/contracts/ERC721A.sol";
import "solmate/src/auth/Owned.sol";

/* 

OK so we have requirements gathered for the first Koi NFT Contract:

3888 Fukushima Fish MAX MINT
5 can be minted at a time MINT PER WALLET
WL Price 0.05
Public Price 0.0777 (subject to change)
*/

// TODO: Mint function
// TODO: Token integration
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

    mapping(uint256 => uint256) public getMintTime;
    mapping(address => uint256) public whitelistMints;

    function exists(uint256 id) external view returns (bool) {
        return _exists(id);
    }

    SupplyController controller;

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

    function mintsRemaining(
        uint256 remaining
    ) internal pure returns (string memory) {
        return
            remaining == 0
                ? NO_MINTS_REMAINING
                : string(
                    abi.encodePacked(
                        "You have ",
                        remaining,
                        " mint(s) remaining"
                    )
                );
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

    function mint(
        uint256 amount,
        uint256 limit,
        bytes32[] calldata proof,
        uint256 path
    ) external payable {
        address msgSender = msg.sender;

        require(amount > 0);

        uint256 currentSupply = _totalMinted();
        // supply limit checks
        require(currentSupply < MAX_SUPPLY, "minted out.");
        require(
            currentSupply + amount <= MAX_SUPPLY,
            "mint amount would be out of range."
        );

        require(amount <= MAX_PUBLIC_MINT_PER_WALLET, "cannot mint more than 10 at a time");

        // botting/contract check (quick externally owned account check)
        require(msgSender == tx.origin);

        uint256 minted = _numberMinted(msgSender);

        if (mintStatus == MintStatus.WHITELIST) {
            require(proof.length > 0, "no proof provided");
            require(
                validate(msgSender, limit, proof, path),
                "invalid merkle proof. you are not whitelisted!"
            );

            // limits only apply to WHITELIST
            // wallet limit checks
            require(minted < limit, "Mint limit reached!");

            uint256 remaining = limit - minted;
            require(
                amount > 0 && amount <= remaining,
                mintsRemaining(remaining)
            );
        }

        // payment checks
        uint256 costPerItem = mintStatus == MintStatus.PUBLIC
            ? PUBLIC_MINT_COST
            : WHITELIST_MINT_COST;
        uint256 minimumPayment = amount * costPerItem;


        uint256 value = msg.value;

        require(value >= minimumPayment, "not enough ether sent for mint!");

        _mint(msgSender, amount);

        if (value > minimumPayment) {
            // refund if the user somehow overpaid
            uint256 refund = value - minimumPayment;
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

    /** 
    
     */

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

        if (from == address(0)) {
            for (uint256 i = startTokenId; i < startTokenId + quantity;) {
                getMintTime[i] = block.timestamp;
                unchecked { i++; }
            }
        }
    }
}
