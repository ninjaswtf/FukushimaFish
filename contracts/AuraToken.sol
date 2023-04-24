pragma solidity ^0.8.9;

import "solmate/src/tokens/ERC20.sol";
import "solmate/src/auth/Owned.sol";
import "./SupplyController.sol";

contract AuraToken is ERC20("Aura", "AURA", 18), Owned(msg.sender) {
    SupplyController public supplyController;

    function setSupplyController(SupplyController _controller) external onlyOwner {
        supplyController = _controller;
    }


    modifier onlyController() {
        require(msg.sender == address(supplyController));
        _;
    }


    function mint(address to, uint256 amount) external onlyController {
        require(supplyController.isMintingAllowed(), "minting is not allowed");
        _mint(to, amount);
    } 


    function burn(address from, uint256 amount) external onlyController {
        require(supplyController.isBurningAllowed(), "burning is not allowed");
        _burn(from, amount);
    }
}