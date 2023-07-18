pragma solidity ^0.8.9;

contract FukushimaFishData {

    error TokenNotInitiated();

    enum AuraLevel {
        NONE,
        LOW,
        MED,
        HIGH,
        OVERFLOWING,
        YIELD_BONUS // certain fish may have a yield bonus that will give them extra AURA for daily yield.
    }


    mapping(AuraLevel => uint256) auraDailyYields;

    // // 0.054 $AURA / day
    // uint256 constant NONE = 0.054 ether;

    // // 0.304 $AURA / day
    // uint256 constant LOW = 0.304 ether;

    // // 0.75 $AURA / day
    // uint256 constant MED = 0.75 ether;

    // // 3 $AURA / day
    // uint256 constant HIGH = 3 ether;

    // // 10 $AURA / day
    // uint256 constant OVERFLOWING = 10 ether;

    // // 20 $AURA / day
    uint256 public BONUS = 20 ether;

    constructor() {
        auraDailyYields[AuraLevel.NONE] = 0.054 ether;
        auraDailyYields[AuraLevel.LOW] = 0.304 ether;
        auraDailyYields[AuraLevel.MED] = 0.75 ether;
        auraDailyYields[AuraLevel.HIGH] = 3 ether;
        auraDailyYields[AuraLevel.OVERFLOWING] = 10 ether; 



        owner = msg.sender;
        _admin[msg.sender] = true;
    }

    address public owner;


    // Uint256 encoded Token Metadata
    mapping(uint256 => uint256) _metadata;
    mapping(address => bool) _admin;

    modifier onlyAdmin() {
        require(_admin[msg.sender]);
        _;
    }

    modifier onlyOwner() {
        require(msg.sender == owner);
        _;
    }

    function transferOwnership(address to) external onlyOwner {
        owner = to;
    }

    function setAdmin(address addr, bool status) external onlyOwner {
        _admin[addr] = status;
    }

    function setBonus(uint256 bonusAmount) external onlyOwner {
        BONUS = bonusAmount;
    }

    /**
     * Updates a given aura level
     * 
     * @param auraLevel the level to update
     * @param amount  the amount for the given level
     */
    function updateYield(AuraLevel auraLevel, uint256 amount) external onlyOwner {

    }

    uint256 constant REACTOR_MODIFIER = 0x01;

    bytes32 public rootHash;


    function importData(bytes32 root) external onlyAdmin {
        rootHash = root;
    }

    function getAuraYieldForToken(uint256 tokenId) external view returns (uint256) {
        // if (tokenYield[tokenId] == 0) revert TokenNotInitiated();
        // return tokenYield[tokenId];

        return 0;
    }


    function decode(uint256 encoded) internal pure returns(uint16 token, uint16 level, uint16 flags) {
        token = uint16(encoded);
        level = uint16 (encoded >> 16);
        flags = uint16(encoded >> 32);
    }

    function initTokenData(
        uint256 encoded,
        uint256 path,
        bytes32[] calldata proof
    ) internal  {

        // token = the token ID, level = the Aura level w/ a modifier
        (uint16 token, uint16 level, uint16 flags) = decode(encoded);

        require(level >= 0 && level <= 9, "nope.");

        // validates the merkle tree
        bytes32 leaf = keccak256(abi.encode(encoded));

        for (uint256 i; i < proof.length; i++) {
            // check if the path is odd and inverse the hash
            if (path & 1 == 1) {
                leaf = keccak256(abi.encodePacked(leaf, proof[i]));
            } else {
                leaf = keccak256(abi.encodePacked(proof[i], leaf));
            }

            path /= 2;
        }
        
        require(leaf == rootHash, "invalid proof.");

        // after verifying the encoded information is legit, set it
    }
}
