pragma solidity ^0.8.9;

import "erc721a/contracts/ERC721A.sol";
import "solmate/src/auth/Owned.sol";
/* 

OK so we have requirements gathered for the first Koi NFT Contract:

3888 Fukushima Fish MAX MINT
5 can be minted at a time MINT PER WALLET
WL Price 0.07899
Public Price 0.0999 (subject to change)

The NFTs will produce RAD (Radiation based on their traits.

RAD will be used on a separate contract to make their dragon


*/
interface RadiationMinter {
    function radiate() external;
}
// TODO: Mint function
// TODO: Token integration
contract FukushimaFishNFT is ERC721A("Fukushima Fish", "FISH"), Owned(msg.sender) {

    enum FukushimaMetadataIndex {
      BACKGROUND
    }

    enum MintStatus {
        // 0 = closed
        CLOSED,
        // 1 = open for whitelisted addresses
        WHITELIST,
        // 2 = open for the public
        PUBLIC
    }

    uint256 constant WHITELIST_MINT_COST = 0.07899 ether;
    // Subject to change
    uint256 constant PUBLIC_MINT_COST = 0.0999 ether;

    uint256 constant MAX_MINT_PER_WALLET = 5;


    // The default mint status is CLOSED
    MintStatus public mintStatus = MintStatus.CLOSED;

    string _baseTokenURI = "";
    string _unrevealedURI = "";

    bytes32 whitelistMerkleProofRoot = bytes32(0);
    
    function setMerkleRoot(bytes32 merkleRoot) external onlyOwner {
        whitelistMerkleProofRoot = merkleRoot;
    }


    function setBaseTokenURI(string calldata uri) external onlyOwner {
         _baseTokenURI = uri;
    } 

    function setUnrevealedTokenURI(string calldata uri) external onlyOwner {
         _unrevealedURI = uri;
    }

    // This data is stored off-chain but we can populate it after mint & reveal.
    mapping(uint256 => uint256) _packedMetadata;

    // sets the required off-chain metadata for the given token. Called post-mint
    function setTokenMetadata( uint16[] calldata tokenIds, uint16[] calldata values ) external onlyOwner {
        for (uint i; i < tokenIds.length; i++) {
            _packedMetadata[tokenIds[i]] = values[i];
        }
    }

    // Validate checks if the given address and proof result in the merkle tree root.
    // if the proof & the hashed address resolves to the provided proof, then the address
    // is within the whitelist.
    function validate(address addr, bytes32[] calldata proof, uint256 path) internal view returns (bool) {
        bytes32 hash = keccak256(abi.encodePacked(addr));

        for (uint i; i < proof.length; i++) {
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


    function tokenURI(uint256 id) public view override returns(string memory) {
        // baseTokenURI is empty, assume the token is unrevealed, and default to the unrevealed URI
        // else concatenate the base URI with the token ID and the JSON URI
        return bytes(_baseTokenURI).length == 0 ? _unrevealedURI :
                   string(abi.encodePacked(_baseTokenURI, id, ".json"));
    }

    function _startTokenId() internal pure override returns (uint256) {
        return 1;
    }

}