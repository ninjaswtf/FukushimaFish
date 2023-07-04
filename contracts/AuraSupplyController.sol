pragma solidity ^0.8.9;


import "./SupplyController.sol";
import "solmate/src/auth/Owned.sol";
import "./AuraToken.sol";
import "./FukushimaFishNFT.sol";
import "./FukushimaFishData.sol";

contract AuraSupplyControllerV1 is SupplyController, Owned(msg.sender) {
    
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


    struct ClaimData {
        uint256 lastClaimedTime;
        bool claimedBefore;
    }
    
    // token id => time last transferred 
    mapping(uint256 => ClaimData) _lastClaimed;



    uint256 constant LEGACY_CONTRACT_SUPPLY_CUTOFF = 323;

    FukushimaFishNFT public oldFukushimaFish;
    FukushimaFishNFT public fukushimaFish;
    FukushimaFishData public data;
    AuraToken public token;

    bool public claimingOpen;

    constructor(FukushimaFishNFT _fukushimaFish, FukushimaFishData _data, AuraToken _token) {
        fukushimaFish = _fukushimaFish;
        token = _token;
        data = _data;
    }


    function setClaimingOpen(bool b) external onlyOwner {
        claimingOpen = b;
    }

    /**
      Claims the tokens
     */
    function claim(uint256 tokenId) external override {
        require(fukushimaFish.ownerOf(tokenId) == msg.sender, "you do not own this NFT");
        uint256 amountClaimable = this.getClaimableTokens(address(0), tokenId);

        require(amountClaimable + token.totalSupply() <= MAX_SUPPLY, "");

        token.mint(msg.sender, amountClaimable);

        if (!_lastClaimed[tokenId].claimedBefore) {
            _lastClaimed[tokenId] = ClaimData(block.timestamp, true);
        } else {
            _lastClaimed[tokenId].lastClaimedTime = block.timestamp;
        }

        emit Claimed(msg.sender, tokenId, amountClaimable);
    }

    function getClaimableTokensID(uint256 tokenId) external view returns (uint256) {
        return this.getClaimableTokens(address(0), tokenId);
    }


    /**
       Gets the claimable tokens available for a given NFT
     */
    function getClaimableTokens(address _ignored, uint256 tokenId) external override view returns (uint256) {
        // if the supply ever reaches the maximum supply we want to disable token generation
        if (token.totalSupply() >= MAX_SUPPLY || !claimingOpen || !fukushimaFish.exists(tokenId)) {
            return 0;
        }

        uint256 timeSinceLastClaim = 0;

        if (!_lastClaimed[tokenId].claimedBefore) {
            timeSinceLastClaim = getMintTime(tokenId);
        } else {
            timeSinceLastClaim = _lastClaimed[tokenId].lastClaimedTime;
        }

        uint256 daysSince = (block.timestamp - timeSinceLastClaim) / 1 days;
        return data.getRadiationYieldForToken(tokenId) * daysSince;
    }


    function getMintTime(uint256 tokenId) internal view returns (uint256) {
        if (tokenId <= LEGACY_CONTRACT_SUPPLY_CUTOFF && address(oldFukushimaFish) != address(0)) {
            return oldFukushimaFish.getMintTime(tokenId);
        }
        return fukushimaFish.getMintTime(tokenId);
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




    function onPostTransfer(address from,
        address to,
        uint256 startTokenId,
        uint256 quantity) external override{
            require(msg.sender == address(fukushimaFish), "NOT ALLOWED");
        }

}