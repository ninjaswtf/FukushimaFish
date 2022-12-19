pragma solidity ^0.8.9;


import "./SupplyController.sol";
import "solmate/src/auth/Owned.sol";
import "./RadiationToken.sol";



interface ERC721 {

    function ownerOf(uint256 token) external returns(address);
}

contract RadiationSupplyControllerV1 is SupplyController, Owned(msg.sender) {



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

    
    // token id => time last transferred 
    mapping(uint256 => uint256) _timeHeld;
   

    ERC721 public fukushimaFish;

    RadiationToken public token;



    bool public claimingOpen;

    constructor(ERC721 _fukushimaFish, RadiationToken _token) {
        fukushimaFish = _fukushimaFish;
        token = _token;
    }


    function setClaimingOpen(bool b) external onlyOwner {
        claimingOpen = b;
    }


    /**
      Claims the tokens
     */
    function claim(uint256 tokenId) external override {
        require(fukushimaFish.ownerOf(tokenId) == msg.sender, "you do not own this NFT");
        uint256 amountClaimable = this.getClaimableTokens(tokenId);

        require(amountClaimable + token.totalSupply() <= MAX_SUPPLY, "");

        token.mint(msg.sender, amountClaimable);
    }


    /**
       Gets the claimable tokens available for a given NFT
     */
    function getClaimableTokens(uint256 tokenId) external override returns (uint256) {
        if (token.totalSupply() >= MAX_SUPPLY) {
            return 0;
        }
        return 0;
    }


    /**
        Determines based on certain criteria if minting (anything that is able to generate new tokens) is allowed.

        e.g. We have a fixed supply and the current token count would surpass the max amount
     */
    function isMintingAllowed() external override returns (bool) {
        return claimingOpen && token.totalSupply() < MAX_SUPPLY;
    }


    /**
       Determines based on certain criteria if burning is allowed.  
     */     
    function isBurningAllowed() external override returns (bool) {
        return false;
    }


}