pragma solidity ^0.8.9;

import "./SupplyController.sol";

contract NoOpSupplyController is SupplyController {
    error TransferDisabled();

    function claim(uint256 tokenId) external virtual override {}

    function getClaimableTokens(
        address a,
        uint256 tokenId
    ) external view virtual override returns (uint256) {
        return 0;
    }

    function isMintingAllowed() external view virtual override returns (bool) {
        return false;
    }

    function isBurningAllowed() external view virtual override returns (bool) {
        return false;
    }

    function onPreTransfer(
        address from,
        address to,
        uint256 startTokenId,
        uint256 quantity
    ) external virtual override {
        revert TransferDisabled();
    }

    function getMaxSupply() external view virtual override returns (uint256) {
        return 0;
    }
}