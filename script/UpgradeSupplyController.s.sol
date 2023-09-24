pragma solidity ^0.8.19;

import "forge-std/Script.sol";

import {FukushimaFishNFT} from "../contracts/FukushimaFishNFT.sol";
import {AuraToken} from "../contracts/AuraToken.sol";
import {AuraSupplyControllerV1} from "../contracts/AuraSupplyController.sol";
import {FukushimaFishData} from "../contracts/FukushimaFishData.sol";

contract UpgradeSupplyController is Script {
    function run() external {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        vm.startBroadcast(deployerPrivateKey);

        AuraSupplyControllerV1 newSupplyController = new AuraSupplyControllerV1();

        address newFukushimaFish = 0xEe259D1aae364bCb2D28333c712d394F052984c3; // this is current/new/fixed Fukushima Fish NFT
        address fukushimaFishData = 0x6071c02BF8F2775d30cA6D94e1D87FC45542e1c1; // This is the current Fukushima Fish Data
        address auraToken = 0x8C4b835BfB8aFED829374fB62Bc8598E324d7C86; // This is the current aura token
        address legacyFukushimaFishNFT = 0x5667B16275eFc836B5e3298409cc9c949fA38970;
        address oldSupplyController = 0x4C3c31518D525d1e1657DBef952A98AA807bC19a;

       // Initialize the token
        newSupplyController.initialize(
            FukushimaFishNFT(newFukushimaFish),
            FukushimaFishData(fukushimaFishData),
            AuraToken(auraToken)
        );


        newSupplyController.setLegacyContract(legacyFukushimaFishNFT);

        newSupplyController.setOldSupplyController(oldSupplyController);

        newSupplyController.setClaimingOpen(true);

        // set the supply controller so the supply controller can mint & burn
        AuraToken(auraToken).setSupplyController(newSupplyController);

        // set claiming open
        newSupplyController.setClaimingOpen(true);

        vm.stopBroadcast();
    }
}
