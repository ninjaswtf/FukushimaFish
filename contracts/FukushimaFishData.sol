pragma solidity ^0.8.9;


contract FukushimaFishData  {

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

    // 0.054 $RAD / day
    uint256 constant NONE = 0.054 ether; 

    // 0.304 $RAD / day
    uint256 constant LOW = 0.304 ether;

    // 0.75 $RAD / day
    uint256 constant MED = 0.75 ether;

    // 3 $RAD / day
    uint256 constant HIGH = 3 ether;

    // 10 $RAD / day
    uint256 constant DANGER = 10 ether;
    
    // 20 $RAD / day
    uint256 constant REACTOR = 20 ether;


    // Radiation levels measured in ether
    mapping(uint256 => uint256) tokenRadiationLevels;




    function setRadiationLevelForToken(uint256 token, uint256 value) external onlyAdmin {
        tokenRadiationLevels[token] = value;
    }

    function getRadiationLevelForToken(uint256 token) external view returns (uint256) {
        return tokenRadiationLevels[token];
    }
}