pragma solidity ^0.8.19;

import "forge-std/Script.sol";

import {FukushimaFishNFT} from "../contracts/FukushimaFishNFT.sol";
import {AuraToken} from "../contracts/AuraToken.sol";
import {AuraSupplyControllerV1} from "../contracts/AuraSupplyController.sol";
import {FukushimaFishData} from "../contracts/FukushimaFishData.sol";

contract TestnetAirdrop is Script {

    function run() external {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        vm.startBroadcast(deployerPrivateKey);

        FukushimaFishNFT currentNft = FukushimaFishNFT(0x5667B16275eFc836B5e3298409cc9c949fA38970);
        address receiver = vm.envAddress("AIRDROP_ADDRESS");

        currentNft.setAirdropper(vm.addr(deployerPrivateKey));
        currentNft.mintAirdropTeamMint(address(receiver), 20);

        vm.stopBroadcast();
    }
}