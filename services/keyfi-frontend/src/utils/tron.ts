let account: string | null = null;
let RealEstateMarketAddress = 'THSc1hFTiLL8kstgYDdtD268vWxP9eF7Jh'; // Paste Contract address here
let realEstateMarket: any = null;

export const accountAddress = (): string | null => {
  return account;
};

export function getTronWeb(): void {
  // Obtain the tronweb object injected by tronLink
  var obj = setInterval(async () => {
    if (window.tronWeb && window.tronWeb.defaultAddress && window.tronWeb.defaultAddress.base58) {
      clearInterval(obj);
      console.log("tronWeb successfully detected!");
    }
  }, 10);
}

export async function setRealEstateMarketContract(): Promise<void> {
  // TODO: obtain contract Object
  realEstateMarket = await (window as any).tronWeb.contract().at(RealEstateMarketAddress);
}

export async function addHomeListing(detailsLink: string, streetAddress: string, listingPrice: number): Promise<void> {
  const result = await realEstateMarket.addHomeListing(detailsLink, streetAddress, 'test', listingPrice).send({
    feeLimit: 100_000_000,
    callValue: 0,
    shouldPollResponse: true,
  });
  console.log(result);
  alert('Listing Posted Successfully');
}

export async function fetchAllListings(): Promise<any[]> {
  const listings: any[] = [];

  const listingId = await realEstateMarket.listingId().call();
  console.log(listingId);
  // iterate from 0 till bookId
  for (let i = 0; i < listingId; i++) {
    const listing = await realEstateMarket.homeListings(i).call();
    if (listing.detailsLink !== "") {
      listings.push(
        { id: i, owner: listing.owner, link: listing.detailsLink, streetAddress: listing.streetAddress, price: (window as any).tronWeb.fromSun(listing.listingPrice) }
      );
    }
  }
  return listings;
}

export async function startSaleContract(listingId: number, listingPrice: number): Promise<any> {
  let OfficialReecorder: any = 'TH9JuMJTTjoQnupHWyAKYVHG5cxebDMbuW';
  let saleContractId:any = await realEstateMarket.startSaleContract(listingId, OfficialReecorder, Math.floor(listingPrice/100000)).send({
    feeLimit: 100_000_000,
    callValue: 0,
    shouldPollResponse: true,
  });
  console.log(saleContractId);
  return saleContractId;
}

export async function makeDownPayment(listingId:number): Promise<any> {
  await realEstateMarket.makeDownPayment(listingId);
}

export async function makePayment(listingId:number): Promise<any> {
  await realEstateMarket.makePayment(listingId);
}

export async function confirmOwnershipTransfer(listingId:number): Promise<any> {
  await realEstateMarket.confirmOwnershipTransfer(listingId);
}

export async function getCurrListingId(): Promise<any> {
  const listingId = await realEstateMarket.listingId().call();
  console.log(listingId);
}