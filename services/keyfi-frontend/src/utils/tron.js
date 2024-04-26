let account = null
let RealEstateMarketAddress = 'TTpHKYaZMkxTMFDocbVwWTf1UJKuGX3Jvq' // Paste Contract address here
let realEstateMarket = null


export const accountAddress = () => {
  return account
}

export function getTronWeb(){
  // Obtain the tronweb object injected by tronLink 
  var obj = setInterval(async ()=>{
    if (window.tronWeb && window.tronWeb.defaultAddress.base58) {
        clearInterval(obj)
        console.log("tronWeb successfully detected!")
    }
  }, 10)
}
 

export async function setRealEstateMarketContract() {
  // TODO: abtain contract Object

  realEstateMarket = await window.tronWeb.contract().at(RealEstateMarketAddress);

}

export async function addHomeListing(detailsLink, streetAddress, listingPrice) {
  const result = await realEstateMarket.addHomeListing(detailsLink, streetAddress, listingPrice).send({
    feeLimit:100_000_000,
    callValue:0,
    shouldPollResponse:true
  })

  alert('Listing Posted Successfully')
}

export async function fetchAllListings() {
  const listings = [];

  const listingId  = await realEstateMarket.listingId().call();
  //iterate from 0 till bookId
  for (let i = 0; i < listingId; i++){
    const listing = await realEstateMarket.homeListings(i).call()
    if(listing.detailsLink!="") // filter the deleted listings
    {
      listings.push(
        {id: i, owner: listing.owner, link: listing.detailsLink, streetAddress: listing.streetAddress, price: window.tronWeb.fromSun(listing.listingPrice)}
      )
    }
    
  }
  return listings

}