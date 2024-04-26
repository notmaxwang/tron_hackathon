pragma solidity ^0.8.0;

contract RealEstateMarket {
    enum State { Pending, DownPayment, FullPricePaid, RecorderConfirmation, Complete }

    struct HomeListing {
        address owner;
        string detailsLink;
        string streetAddress;
        uint listingPrice;
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
    }

    uint256 public listingId;
    uint256 public saleContractId;
    mapping(uint256 => HomeListing) public homeListings;
    mapping(uint256 => HomeSaleContract) public saleContracts;
    // mapping(address => uint256[]) public ownerListings;
    // mapping(address => uint256[]) public outboundContracts;
    // mapping(address => uint256[]) public inboundContracts;

    modifier onlyBuyer() {
        require(msg.sender == saleContracts[saleContractId].buyer, "Only buyer can call this function");
        _;
    }

    modifier onlySeller() {
        require(msg.sender == saleContracts[saleContractId].seller, "Only seller can call this function");
        _;
    }

    modifier onlyRecorder() {
        require(msg.sender == saleContracts[saleContractId].officialRecorder, "Only seller can call this function");
        _;
    }

    modifier allParties() {
        require(msg.sender == saleContracts[saleContractId].officialRecorder || msg.sender == saleContracts[saleContractId].buyer || msg.sender == saleContracts[saleContractId].seller, "Only relevant parties can call this function");
        _;
    }

    modifier inState(State _state) {
        require(saleContracts[saleContractId].currentState == _state, "Invalid state transition");
        _;
    }

    modifier beforeFullPricePaid() {
        require(saleContracts[saleContractId].currentState == State.Pending || saleContracts[saleContractId].currentState == State.DownPayment, "Before the full price is paid");
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
    event HomeListingAdded(address owner, string detailsLink, string streetAddress, uint listingPrice);
    event HomeListingRemoved(uint256 listingId);

    function addHomeListing(string memory _detailsLink, string memory _streetAddress, uint _listingPrice) external returns (uint256) {
        HomeListing memory newListing = HomeListing(msg.sender, _detailsLink, _streetAddress, _listingPrice);
        homeListings[listingId] = newListing;

        // uint length = ownerListings[msg.sender].length;
        // ownerListings[msg.sender].push();
        // ownerListings[msg.sender][length] = listingId;
        emit HomeListingAdded(msg.sender, _detailsLink, _streetAddress, _listingPrice);

        return listingId;
    }

    function removeHomeListing(uint256 _listingId) external {
        require(msg.sender == homeListings[_listingId].owner, "only the owner of listing can remove it");

        delete homeListings[_listingId];
        emit HomeListingRemoved(_listingId);
    }

    function startSaleContract(uint256 _listingId, address _officialRecorder, uint _price) external returns (uint256) {
        require(saleContracts[saleContractId].price == 0, "Sale contract already exists");
        
        HomeSaleContract memory saleContract = HomeSaleContract(_listingId, msg.sender, homeListings[_listingId].owner, _officialRecorder, 0, _price, false, false, State.Pending);
        saleContracts[saleContractId] = saleContract;

        // uint outboundContractsLength = outboundContracts[msg.sender].length;
        // outboundContracts[msg.sender].push();
        // outboundContracts[msg.sender][outboundContractsLength] = _listingId;

        // uint inboundContractsLength = inboundContracts[homeListings[_listingId].owner].length;
        // inboundContracts[homeListings[_listingId].owner].push();
        // inboundContracts[homeListings[_listingId].owner][inboundContractsLength] = _listingId;
        emit SaleContractInitiated();

        return saleContractId;
    }

    function makeDownPayment() external payable onlyBuyer inState(State.Pending) {
        emit DownPaymentMade(msg.value);
        saleContracts[saleContractId].currentState = State.DownPayment;
    }

    function payFullPrice() external payable onlyBuyer inState(State.DownPayment) {
        require(saleContracts[saleContractId].buyerApproval == true && saleContracts[saleContractId].sellerApproval == true, "Requires buyer and seller approval before moving forward");
        require(msg.value + saleContracts[saleContractId].downPayment == saleContracts[saleContractId].price, "Incorrect amount for full price payment");
        emit FullPricePaid(msg.value);
        saleContracts[saleContractId].currentState = State.FullPricePaid;
    }

    function confirmOwnershipTransfer() external onlyRecorder inState(State.FullPricePaid) {
        emit RecorderConfirmation();
        saleContracts[saleContractId].currentState = State.RecorderConfirmation;

        _completeTransaction();
    }

    function _completeTransaction() internal inState(State.RecorderConfirmation) {
        payable(saleContracts[saleContractId].seller).transfer(address(this).balance);
        saleContracts[saleContractId].currentState = State.Complete;
        emit Complete();

        delete saleContracts[saleContractId];
        delete homeListings[listingId];
    }

    function approveBuyer() external onlyBuyer beforeFullPricePaid {
        saleContracts[saleContractId].buyerApproval = true;
        emit BuyerApproved();
    }

    function revokeBuyerApproval() external onlyBuyer beforeFullPricePaid {
        saleContracts[saleContractId].buyerApproval = false;
        emit BuyerRevoked();
    }

    function approveSeller() external onlySeller beforeFullPricePaid {
        saleContracts[saleContractId].sellerApproval = true;
        emit SellerApproved();
    }

    function revokeSellerApproval() external onlySeller beforeFullPricePaid {
        saleContracts[saleContractId].sellerApproval = false;
        emit SellerRevoked();
    }
}
