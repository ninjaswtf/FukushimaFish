pragma solidity ^0.8.19;

import "./FukushimaFishNFT.sol";

contract Airdropper {
    FukushimaFishNFT public nft;
    address public owner;

    constructor(address _nft) {
        nft = FukushimaFishNFT(_nft);
        owner = msg.sender;
    }

    function airdropToHolders(address[] calldata uniqueHolders) external {
        require(msg.sender == owner);
        
        for (uint i; i < uniqueHolders.length; ) {
            address holder = uniqueHolders[i];
            nft.mintAirdropTeamMint(holder, nft.balanceOf(holder));
            unchecked {
                i++;
            }
        }
    }
}
