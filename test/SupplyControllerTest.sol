// SPDX-License-Identifier: SEE LICENSE IN LICENSE
pragma solidity ^0.8.19;



import {Bootstrapper} from "./Bootstrapper.sol";

import {FukushimaFishNFT} from "../contracts/FukushimaFishNFT.sol";
import {AuraToken} from "../contracts/AuraToken.sol";
import {AuraSupplyControllerV1} from "../contracts/AuraSupplyController.sol";
import {FukushimaFishData} from "../contracts/FukushimaFishData.sol";


contract SupplyControllerTest is Bootstrapper {
    function setUp() public override virtual {
        super.setUp();
    }


    function testGas() external {

        vm.deal(address(msg.sender), 200 ether);

        fukushimaFishNft.setMintStatus(FukushimaFishNFT.MintStatus.PUBLIC);

        
       (bool ok,) = address(fukushimaFishNft).call{value: 20 ether}(abi.encodeWithSignature("publicMint(uint256)", 20));

       require(ok);
    }


    function testInitData() external {

        // Import the merkle tree hash
        data.importData(0x3362d291c96e9e31e53eaa382ba10247aae99ab8e9d8e950041d5a0e4a97d643);

        
        bytes32[] memory proof = new bytes32[](12);

        proof[0] = bytes32(0xad4526fd3a02914e81050e8cc9cb9fa0cd4f3ab503027cc5783aa1c93a47d93a);
        proof[1] = bytes32(0xba7b00b420a962a373b9986ded6b006e7556011a9043a5bceab9f05409ab86b2);
        proof[2] = bytes32(0x336bfef6ffe9b81a0b7cde1edc5b45e00b36587e61c89072857872f21eb861c9);
        proof[3] = bytes32(0x3e8a11e0640cf9b030bf6f823c551a7f5ad5236bbea22f662bead372a0e26859);
        proof[4] = bytes32(0x35573adae924c7d0883805d3682e5b044922b9a9b6cee2087dd2a5138ae0a79b);
        proof[5] = bytes32(0x5c90084edfd16613f374399b15d1ac3ea48ff16876bcd3f663f29cc69ff9b1fc);
        proof[6] = bytes32(0x20b2d47998a18b75b6c8c5055fb79042015691fcec5c73bc634687ebc1759f15);
        proof[7] = bytes32(0x5f4264e0d1d46876472989982f719f66f2020ac3b3291028f251606e6c3ba21a);
        proof[8] = bytes32(0x1c9f92bc6d4c0c5e6309543da2b4870e092e94562bb75d7edd1f78438688baa6);
        proof[9] = bytes32(0x1feab68656e27fde64f3fce6a70e65720e00317ccf2b4f27cfbdc16d0ed44ada);
        proof[10] = bytes32(0x7571059c2d17f2ca00e7c05a7c3466d8492ee9351cfddd9738297464b740cd80);
        proof[11] = bytes32(0x2fa17c296f3730beb9fef2dc000ea87572224219428749b0ce87dda3acda5588);
             

        uint256 encodedToken = 0x0000000000000000000000000000000000000000000000000000000100030042;
        uint256 path = 4030;

        uint256 tokenId = 66;

        data.initTokenData(encodedToken, path, proof);  

        require(data.getAuraYieldForToken(tokenId) == 3 ether);
        require(data.isCelestial(tokenId));
    }



    function testInitData2() external {

        // Import the merkle tree hash
        data.importData(0x3362d291c96e9e31e53eaa382ba10247aae99ab8e9d8e950041d5a0e4a97d643);

        bytes32[] memory proof = new bytes32[](12);

        proof[0] = bytes32(0xb11b999f6fba5941926e35b4d09fb54f7925d46844349198d2be2975104cacc7);
        proof[1] = bytes32(0x3b17ee4825bfb6f9e93d28d3c1db83b11a0a932f168f53aad84a2e35b4371c4d);
        proof[2] = bytes32(0x90cd1557b7fb06a20901a67fc4c8442b6f70e9b8491875e0a1a7845135efe4db);
        proof[3] = bytes32(0x558c6b0c15692458ea52e60f1d6658ca2d1e5ee5ebcecbf1ce92cc47fc61d83b);
        proof[4] = bytes32(0xd74d11b8a91ba68b719532370abc8dd9803f8b699a65b832fd168db312f6e4b3);
        proof[5] = bytes32(0x47a815d78e1561d1f0c89cda0c90c50589d9aadea5eabb17848b34a4cfad04d0);
        proof[6] = bytes32(0x20b2d47998a18b75b6c8c5055fb79042015691fcec5c73bc634687ebc1759f15);
        proof[7] = bytes32(0x5f4264e0d1d46876472989982f719f66f2020ac3b3291028f251606e6c3ba21a);
        proof[8] = bytes32(0x1c9f92bc6d4c0c5e6309543da2b4870e092e94562bb75d7edd1f78438688baa6);
        proof[9] = bytes32(0x1feab68656e27fde64f3fce6a70e65720e00317ccf2b4f27cfbdc16d0ed44ada);
        proof[10] = bytes32(0x7571059c2d17f2ca00e7c05a7c3466d8492ee9351cfddd9738297464b740cd80);
        proof[11] = bytes32(0x2fa17c296f3730beb9fef2dc000ea87572224219428749b0ce87dda3acda5588);
             
        uint256 encodedToken = 0x0000000000000000000000000000000000000000000000000000000000040074;
        uint256 path = 3980;

        uint256 tokenId = 116;

        data.initTokenData(encodedToken, path, proof);  

        require(data.getAuraYieldForToken(tokenId) == 10 ether);
        require(!data.isCelestial(tokenId));
    }







    function encode(uint16 token, uint64 level, uint16 flags) internal pure returns(uint256) {
        uint256 encoded = uint256(uint16(flags));

        encoded <<= 16;
        encoded |= level;

        encoded <<= 16;
        encoded |= token;
        
        return encoded;
    }


    function testManipulation() external {

        // Import the merkle tree hash
        data.importData(0x3362d291c96e9e31e53eaa382ba10247aae99ab8e9d8e950041d5a0e4a97d643);

        
        bytes32[] memory proof = new bytes32[](12);

        proof[0] = bytes32(0xf61307ef97fd50c28217ff83c38146e46e7cb54a28445e5805135d803ad29b4e);
        proof[1] = bytes32(0x8af81de08eeeebcb1203f30123da6c6b6f7b1d62aca762b8acb1ff9a236c3d59);
        proof[2] = bytes32(0x35447f947ced343a8168c2d5909dd47fb04ba7012a6ff86f39f7274b731d1cf2);
        proof[3] = bytes32(0xb07663d517f478bd53d5298c70cf7d7267495fd7483346fee4372f23fba081fd);
        proof[4] = bytes32(0x35573adae924c7d0883805d3682e5b044922b9a9b6cee2087dd2a5138ae0a79b);
        proof[5] = bytes32(0x5c90084edfd16613f374399b15d1ac3ea48ff16876bcd3f663f29cc69ff9b1fc);
        proof[6] = bytes32(0x20b2d47998a18b75b6c8c5055fb79042015691fcec5c73bc634687ebc1759f15);
        proof[7] = bytes32(0x5f4264e0d1d46876472989982f719f66f2020ac3b3291028f251606e6c3ba21a);
        proof[8] = bytes32(0x1c9f92bc6d4c0c5e6309543da2b4870e092e94562bb75d7edd1f78438688baa6);
        proof[9] = bytes32(0x1feab68656e27fde64f3fce6a70e65720e00317ccf2b4f27cfbdc16d0ed44ada);
        proof[10] = bytes32(0x7571059c2d17f2ca00e7c05a7c3466d8492ee9351cfddd9738297464b740cd80);
        proof[11] = bytes32(0x2fa17c296f3730beb9fef2dc000ea87572224219428749b0ce87dda3acda5588);
             

        uint256 encodedToken = encode(75, 4, 0x01);
        uint256 path = 4021;

        vm.expectRevert();
        data.initTokenData(encodedToken, path, proof);  
    }


    function testClaimMainnet() external {

        vm.createSelectFork(vm.envString("MAINNET_RPC_URL"));

        // Test claiming with a mainnet configuration
        address oldMainnetAddress = 0x5667B16275eFc836B5e3298409cc9c949fA38970;
        address currentMainnetAddress = 0x5667B16275eFc836B5e3298409cc9c949fA38970;

        FukushimaFishNFT currentNft = FukushimaFishNFT(currentMainnetAddress);
        // Deploy the data contract
        FukushimaFishData data = new FukushimaFishData();

        // Deploy the aura token
        AuraToken auraToken = new AuraToken();
        auraToken.initialize("1.0.0");

        AuraSupplyControllerV1 supplyController = new AuraSupplyControllerV1(
            currentNft,
            data,
            auraToken
        );






        // Import the merkle tree hash
        data.importData(0x3362d291c96e9e31e53eaa382ba10247aae99ab8e9d8e950041d5a0e4a97d643);

        
        bytes32[] memory proof = new bytes32[](12);

        proof[0] = bytes32(0xad4526fd3a02914e81050e8cc9cb9fa0cd4f3ab503027cc5783aa1c93a47d93a);
        proof[1] = bytes32(0xba7b00b420a962a373b9986ded6b006e7556011a9043a5bceab9f05409ab86b2);
        proof[2] = bytes32(0x336bfef6ffe9b81a0b7cde1edc5b45e00b36587e61c89072857872f21eb861c9);
        proof[3] = bytes32(0x3e8a11e0640cf9b030bf6f823c551a7f5ad5236bbea22f662bead372a0e26859);
        proof[4] = bytes32(0x35573adae924c7d0883805d3682e5b044922b9a9b6cee2087dd2a5138ae0a79b);
        proof[5] = bytes32(0x5c90084edfd16613f374399b15d1ac3ea48ff16876bcd3f663f29cc69ff9b1fc);
        proof[6] = bytes32(0x20b2d47998a18b75b6c8c5055fb79042015691fcec5c73bc634687ebc1759f15);
        proof[7] = bytes32(0x5f4264e0d1d46876472989982f719f66f2020ac3b3291028f251606e6c3ba21a);
        proof[8] = bytes32(0x1c9f92bc6d4c0c5e6309543da2b4870e092e94562bb75d7edd1f78438688baa6);
        proof[9] = bytes32(0x1feab68656e27fde64f3fce6a70e65720e00317ccf2b4f27cfbdc16d0ed44ada);
        proof[10] = bytes32(0x7571059c2d17f2ca00e7c05a7c3466d8492ee9351cfddd9738297464b740cd80);
        proof[11] = bytes32(0x2fa17c296f3730beb9fef2dc000ea87572224219428749b0ce87dda3acda5588);
             

        uint256 encodedToken = 0x0000000000000000000000000000000000000000000000000000000100030042;
        uint256 path = 4030;

        uint256 tokenId = 66;

        data.initTokenData(encodedToken, path, proof);  



        // set the legacy contract address
        supplyController.setLegacyContract(oldMainnetAddress);
        // set the supply controller so the supply controller can mint & burn
        auraToken.setSupplyController(supplyController);


        supplyController.setClaimingOpen(true);

        vm.startPrank(0xF42D1c0c0165AF5625b2ecD5027c5C5554e5b039);


        vm.expectRevert(); // revert because i did not initialize the token.
        supplyController.claim(75);

        vm.expectRevert(); // revert because i do not own this token
        supplyController.claim(66);




        
        proof = new bytes32[](12);

        proof[0] = bytes32(0xf61307ef97fd50c28217ff83c38146e46e7cb54a28445e5805135d803ad29b4e);
        proof[1] = bytes32(0x8af81de08eeeebcb1203f30123da6c6b6f7b1d62aca762b8acb1ff9a236c3d59);
        proof[2] = bytes32(0x35447f947ced343a8168c2d5909dd47fb04ba7012a6ff86f39f7274b731d1cf2);
        proof[3] = bytes32(0xb07663d517f478bd53d5298c70cf7d7267495fd7483346fee4372f23fba081fd);
        proof[4] = bytes32(0x35573adae924c7d0883805d3682e5b044922b9a9b6cee2087dd2a5138ae0a79b);
        proof[5] = bytes32(0x5c90084edfd16613f374399b15d1ac3ea48ff16876bcd3f663f29cc69ff9b1fc);
        proof[6] = bytes32(0x20b2d47998a18b75b6c8c5055fb79042015691fcec5c73bc634687ebc1759f15);
        proof[7] = bytes32(0x5f4264e0d1d46876472989982f719f66f2020ac3b3291028f251606e6c3ba21a);
        proof[8] = bytes32(0x1c9f92bc6d4c0c5e6309543da2b4870e092e94562bb75d7edd1f78438688baa6);
        proof[9] = bytes32(0x1feab68656e27fde64f3fce6a70e65720e00317ccf2b4f27cfbdc16d0ed44ada);
        proof[10] = bytes32(0x7571059c2d17f2ca00e7c05a7c3466d8492ee9351cfddd9738297464b740cd80);
        proof[11] = bytes32(0x2fa17c296f3730beb9fef2dc000ea87572224219428749b0ce87dda3acda5588);
             

        encodedToken = 0x000000000000000000000000000000000000000000000000000000000002004b;
        path = 4021;

        
        data.initTokenData(encodedToken, path, proof);
        supplyController.claim(75);


        // expect to fail because i do not have tokens to claim right now.
        vm.expectRevert();
        supplyController.claim(75);


        vm.stopPrank();



     }
}