pragma solidity ^0.8.19;

import "forge-std/Script.sol";

import {FukushimaFishNFT} from "../contracts/FukushimaFishNFT.sol";
import {AuraToken} from "../contracts/AuraToken.sol";
import {AuraSupplyControllerV1} from "../contracts/AuraSupplyController.sol";
import {FukushimaFishData} from "../contracts/FukushimaFishData.sol";
import {ERC1967Proxy} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import {SupplyController} from "../contracts/SupplyController.sol";

contract FreshDeploy is Script {
    function run() external {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        vm.startBroadcast(deployerPrivateKey);

        FukushimaFishNFT currentNft = new FukushimaFishNFT();

        // Deploy the data contract
        FukushimaFishData data = new FukushimaFishData();

        // Import the merkle tree hash
        data.importData(
            0x3362d291c96e9e31e53eaa382ba10247aae99ab8e9d8e950041d5a0e4a97d643
        );

        // Deploy the aura token
        AuraToken auraToken = new AuraToken();

        ERC1967Proxy auraTokenProxy = new ERC1967Proxy(
            address(auraToken),
            abi.encodeWithSelector(auraToken.initialize.selector, "1.0.0")
        );

        AuraSupplyControllerV1 supplyController = new AuraSupplyControllerV1();

        ERC1967Proxy supplyControllerProxy = new ERC1967Proxy(
            address(supplyController),
            abi.encodeWithSelector(
                supplyController.initialize.selector,
                address(currentNft),
                address(data),
                address(auraToken)
            )
        );

        // set the supply controller so the supply controller can mint & burn
        AuraToken(address(auraTokenProxy)).setSupplyController(SupplyController(address(supplyControllerProxy)));

        // set claiming open
        AuraSupplyControllerV1(address(supplyControllerProxy)).setClaimingOpen(true);

        vm.stopBroadcast();
    }
}
