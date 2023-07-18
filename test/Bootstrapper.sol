// SPDX-License-Identifier: SEE LICENSE IN LICENSE
pragma solidity ^0.8.19;

import {Test} from "forge-std/Test.sol";

import {FukushimaFishNFT} from "../contracts/FukushimaFishNFT.sol";
import {AuraToken} from "../contracts/AuraToken.sol";
import {AuraSupplyControllerV1} from "../contracts/AuraSupplyController.sol";
import {FukushimaFishData} from "../contracts/FukushimaFishData.sol";

contract Bootstrapper is Test {

    FukushimaFishNFT public fukushimaFishNft;
    AuraToken public auraToken;
    FukushimaFishData public data;
    AuraSupplyControllerV1 public supplyController;

    function setUp() public virtual {
        // Deploy a new Fukushima Fish NFT contract
        fukushimaFishNft = new FukushimaFishNFT();


        // Deploy the data contract
        data = new FukushimaFishData();

        // Deploy the aura token
        auraToken = new AuraToken();
        auraToken.initialize("1.0.0");

        supplyController = new AuraSupplyControllerV1(
            fukushimaFishNft,
            data,
            auraToken
        );
    }
}