pragma solidity ^0.8.9;

import "solmate/src/tokens/ERC20.sol";


contract RadiationToken is ERC20("Radiation", "RAD", 18) {

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


    // The max supply of the token is 5 million (ether, 18 decimals)
    uint256 constant MAX_SUPPLY = 5_000_000 ether;
}