import './Listing.css'

export default function Listing(props: any) {
  const listing = props.listing;
  
  return(
    <>
      <div className="listing-container">
        <img src="" alt="listing image" />
        <p>{listing.price}</p>
        <p>{listing.address}, {listing.city}, {listing.state}, {listing.zip_code}</p>
        <p>Total Area : {listing.sqft_area}</p>
        <p>{listing.beds} beds, {listing.baths} baths</p>
      </div>
    </>
  );
}