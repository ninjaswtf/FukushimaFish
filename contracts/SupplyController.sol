pragma solidity ^0.8.9;

abstract contract SupplyController {
    
    /**
        Called when a claim is successfully issued
     */
    event Claimed(address claimer, uint256 tokenId, uint256 tokensClaimed);


    event Burned(address burner, uint256 tokensBurned);


    /**
      Claims the tokens
     */
    function claim(uint256 tokenId) external virtual;


    /**
       Gets the claimable tokens available for a given NFT
     */
    function getClaimableTokens(address a, uint256 tokenId) external virtual returns (uint256); 


    /**
        Determines based on certain criteria if minting (anything that is able to generate new tokens) is allowed.

        e.g. We have a fixed supply and the current token count would surpass the max amount
     */
    function isMintingAllowed() external virtual returns (bool);


    /**
       Determines based on certain criteria if burning is allowed.  
     */     
    function isBurningAllowed() external virtual returns (bool);



    /**
      Events
     */
     function onPreTransfer(address from,
        address to,
        uint256 startTokenId,
        uint256 quantity) external virtual{}


    function onPostTransfer(address from,
        address to,
        uint256 startTokenId,
        uint256 quantity) external virtual{}

}
