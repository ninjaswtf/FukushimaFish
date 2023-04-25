pragma solidity ^0.8.9;

contract FukushimaFishData {


    error TokenNotInitiated();

    // 0.054 $AURA / day
    uint256 constant NONE = 0.054 ether;

    // 0.304 $AURA / day
    uint256 constant LOW = 0.304 ether;

    // 0.75 $AURA / day
    uint256 constant MED = 0.75 ether;

    // 3 $AURA / day
    uint256 constant HIGH = 3 ether;

    // 10 $AURA / day
    uint256 constant OVERFLOWING = 10 ether;

    // 20 $AURA / day
    uint256 constant REACTOR = 20 ether;

    address public owner;
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

    constructor() {
        owner = msg.sender;
        _admin[msg.sender] = true;
    }

    uint256 constant REACTOR_MODIFIER = 0x01;
    uint256[] public YIELD_ARRAY = [NONE, LOW, MED, HIGH, OVERFLOWING];

    bytes32 public rootHash;

    // this maps a tokenId to a token yield
    mapping(uint256 => uint256) tokenYield;

    function importData(bytes32 root) external onlyAdmin {
        rootHash = root;
    }

    function getRadiationYieldForToken(uint256 tokenId) external view returns (uint256) {
        if (tokenYield[tokenId] == 0) revert TokenNotInitiated();
        return tokenYield[tokenId];
    }

    function initTokenData(
        uint256 tokenId,
        uint256 level,
        uint256 path,
        bytes32[] calldata proof
    ) external returns (uint256)  {
        require(level >= 0 && level <= 9, "nope.");

        // validates the merkle tree
        bytes32 leaf = keccak256(abi.encode(tokenId, level));

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


        uint256 finalTokenYield = 0;
        // the level is zero, radiation level is none
        if (level == 0) {
            // NONE
            finalTokenYield += NONE;

            // the fish has no radiation level but has a reactor background
        } else if (level == REACTOR_MODIFIER) {
            finalTokenYield += REACTOR;
        } else {
            bool hasReactor = (level % 2 != 0);
            uint256 levelIndex = (level - (hasReactor ? REACTOR_MODIFIER : 0)) /
                2;

            finalTokenYield += YIELD_ARRAY[levelIndex];

            if (hasReactor) {
                finalTokenYield += REACTOR;
            }
        }

        tokenYield[tokenId] = finalTokenYield;

        return finalTokenYield;
    }
}
