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


    // Radiation levels measured in ether
    mapping(uint256 => uint256) tokenRadiationLevels;


    function setRadiationLevelForToken(uint256 token, uint256 value) external onlyAdmin {
        tokenRadiationLevels[token] = value;
    }

    function getRadiationLevelForToken(uint256 token) external view returns (uint256) {
        return tokenRadiationLevels[token];
    }
}