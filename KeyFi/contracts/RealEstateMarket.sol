pragma solidity ^0.8.0;

contract RealEstateMarket {
    enum State { Pending, DownPayment, FullPricePaid, RecorderConfirmation, Complete }

    struct HomeListing {
        address owner;
        string detailsLink;
        string streetAddress;
        string gps;
        uint listingPrice;
        uint256[] openContracts;
    }

    struct HomeSaleContract {
        uint256 listingId;
        address buyer;
        address seller;
        address officialRecorder;
        uint downPayment;
        uint price;
        bool buyerApproval;
        bool sellerApproval;
        State currentState;
        uint256 expiry;
    }

    uint256 public listingId;
    uint256 public saleContractId;
    mapping(uint256 => HomeListing) public homeListings;
    mapping(uint256 => HomeSaleContract) public saleContracts;
    // mapping(address => uint256[]) public ownerListings;
    // mapping(address => uint256[]) public outboundContracts;
    // mapping(address => uint256[]) public inboundContracts;

    modifier inState(State _state) {
        require(saleContracts[saleContractId].currentState == _state, "Invalid state transition");
        _;
    }

    modifier beforeFullPricePaid() {
        require(saleContracts[saleContractId].currentState == State.Pending 
            || saleContracts[saleContractId].currentState == State.DownPayment, "Before the full price is paid");
        _;
    }

    event SaleContractInitiated();
    event DownPaymentMade(uint amount);
    event FullPricePaid(uint amount);
    event RecorderConfirmation();
    event Complete();
    event BuyerApproved();
    event BuyerRevoked();
    event SellerApproved();
    event SellerRevoked();
    event Withdrawal();
    event HomeListingAdded(address owner, string detailsLink, string streetAddress, uint listingPrice);
    event HomeListingRemoved(uint256 listingId);

    function addHomeListing(string memory _detailsLink, string memory _streetAddress, string memory _gps, uint _listingPrice) external returns (uint256) {
        uint256[] memory openContracts;
        HomeListing memory newListing = HomeListing(msg.sender, _detailsLink, _streetAddress, _gps, _listingPrice, openContracts);
        homeListings[listingId] = newListing;

        // uint length = ownerListings[msg.sender].length;
        // ownerListings[msg.sender].push();
        // ownerListings[msg.sender][length] = listingId;
        emit HomeListingAdded(msg.sender, _detailsLink, _streetAddress, _listingPrice);

        listingId += 1;
        return listingId;
    }

    function removeHomeListing(uint256 _listingId) external {
        require(msg.sender == homeListings[_listingId].owner, "only the owner of listing can remove it");
        require(homeListings[_listingId].openContracts.length > 0, "there are still open contracts");
        require(_listingId < listingId, "listing does not exist");

        delete homeListings[_listingId];
        emit HomeListingRemoved(_listingId);
    }

    function startSaleContract(uint256 _listingId, address _officialRecorder, uint _price) external returns (uint256) { 
        HomeSaleContract memory saleContract = HomeSaleContract(_listingId, msg.sender, homeListings[_listingId].owner, _officialRecorder, 0, 
            _price, false, false, State.Pending, type(uint256).max);
        saleContracts[saleContractId] = saleContract;
        homeListings[_listingId].openContracts.push(saleContractId);

        // uint outboundContractsLength = outboundContracts[msg.sender].length;
        // outboundContracts[msg.sender].push();
        // outboundContracts[msg.sender][outboundContractsLength] = _listingId;

        // uint inboundContractsLength = inboundContracts[homeListings[_listingId].owner].length;
        // inboundContracts[homeListings[_listingId].owner].push();
        // inboundContracts[homeListings[_listingId].owner][inboundContractsLength] = _listingId;
        emit SaleContractInitiated();

        uint256 currId = saleContractId;
        saleContractId += 1;
        return currId;
    }

    function makeDownPayment(uint256 _saleContractId) external payable inState(State.Pending) {
        emit DownPaymentMade(msg.value);
        saleContracts[_saleContractId].currentState = State.DownPayment;
    }

    function makePayment(uint256 _saleContractId) external payable inState(State.DownPayment) {
        require(saleContracts[_saleContractId].buyerApproval == true 
            && saleContracts[_saleContractId].sellerApproval == true, 
            "Requires buyer and seller approval before moving forward");
        require(msg.value + saleContracts[_saleContractId].downPayment == saleContracts[_saleContractId].price, 
            "Incorrect amount for full price payment");
        saleContracts[_saleContractId].currentState = State.FullPricePaid;
        saleContracts[_saleContractId].expiry = block.timestamp + 30 days;

        emit FullPricePaid(msg.value);
        
    }

    function confirmOwnershipTransfer(uint256 _saleContractId) external inState(State.FullPricePaid) {
        require(saleContracts[_saleContractId].officialRecorder == msg.sender, "can only be confirmed by recorder");

        saleContracts[_saleContractId].currentState = State.RecorderConfirmation;
        emit RecorderConfirmation();

        _completeTransaction(_saleContractId);
    }

    function _completeTransaction(uint256 _saleContractId) internal inState(State.RecorderConfirmation) {
        payable(saleContracts[_saleContractId].seller).transfer(address(this).balance);
        saleContracts[_saleContractId].currentState = State.Complete;
        emit Complete();

        delete homeListings[saleContracts[_saleContractId].listingId];
        delete saleContracts[_saleContractId];
    }

    function approveBuyer(uint256 _saleContractId) external beforeFullPricePaid {
        require(saleContracts[_saleContractId].buyer == msg.sender, "contract must be approved by buyer");

        saleContracts[_saleContractId].buyerApproval = true;
        emit BuyerApproved();
    }

    function revokeBuyerApproval(uint256 _saleContractId) external beforeFullPricePaid {
        require(saleContracts[_saleContractId].buyer == msg.sender, "contract must be revoked by buyer");

        saleContracts[_saleContractId].buyerApproval = false;
        emit BuyerRevoked();
    }

    function approveSeller(uint256 _saleContractId) external beforeFullPricePaid {
        require(saleContracts[_saleContractId].seller == msg.sender, "contract must be approved by seller");

        saleContracts[_saleContractId].sellerApproval = true;
        emit SellerApproved();
    }

    function revokeSellerApproval(uint256 _saleContractId) external beforeFullPricePaid {
        require(saleContracts[_saleContractId].seller == msg.sender, "contract must be revoked by seller");

        saleContracts[_saleContractId].sellerApproval = false;
        emit SellerRevoked();
    }

    function withdrawAfterPaid(uint256 _saleContractId) external {
        require(saleContracts[_saleContractId].currentState == State.FullPricePaid, "can only be done after buyer paid in full");
        require(saleContracts[_saleContractId].expiry < block.timestamp, "can only be done after expiry date");
        require(saleContracts[_saleContractId].buyer == msg.sender 
            || saleContracts[_saleContractId].seller == msg.sender, 
            "Withdrawal can only be initiated by the buyer or seller");

        uint256 amountToWithdraw = saleContracts[_saleContractId].downPayment;
        saleContracts[_saleContractId].downPayment = 0; // Set down payment to zero to avoid re-entrancy attacks
        payable(msg.sender).transfer(amountToWithdraw);

        for (uint i = 0; i<homeListings[saleContracts[_saleContractId].listingId].openContracts.length-1; i++){
            if (homeListings[saleContracts[_saleContractId].listingId].openContracts[i] == _saleContractId) {
                delete homeListings[saleContracts[_saleContractId].listingId].openContracts[i];
                break;
            }
        }

        emit Withdrawal();
    }

    function withdraw(uint256 _saleContractId) external beforeFullPricePaid {
        require(saleContracts[_saleContractId].buyer == msg.sender 
            || saleContracts[_saleContractId].seller == msg.sender, 
            "Withdrawal can only be initiated by the buyer or seller");

        uint256 amountToWithdraw = saleContracts[_saleContractId].downPayment;
        saleContracts[_saleContractId].downPayment = 0; // Set down payment to zero to avoid re-entrancy attacks
        payable(msg.sender).transfer(amountToWithdraw);

        for (uint i = 0; i<homeListings[saleContracts[_saleContractId].listingId].openContracts.length-1; i++){
            if (homeListings[saleContracts[_saleContractId].listingId].openContracts[i] == _saleContractId) {
                delete homeListings[saleContracts[_saleContractId].listingId].openContracts[i];
                break;
            }
        }

        emit Withdrawal();
    }
}
