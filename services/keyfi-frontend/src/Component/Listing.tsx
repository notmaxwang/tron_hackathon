import { Link } from 'react-router-dom';
import './Listing.css'
import { useState, useEffect } from 'react';

export default function Listing(props: any) {
  const [images, setImages] = useState<any>([]);
  const listing = props.listing;

  useEffect(() => {
    fetch("https://pretentiousbruv.github.io/images/get_image_urls/" + listing.imageKey.toLowerCase() + ".json")
    .then((res) => res.json())
    .then((res) => {
      if (res.success) {
        setImages(res.images)
      }
      else {
        setImages(res.images)
      }
    })
  }, [])
  
  return(
    <>
      <div className="listing-container">
        <img src={images[0]} alt="listing image" />
        <p>{listing.price}</p>
        <p>{listing.address}, {listing.city}, {listing.state}, {listing.zip_code}</p>
        <p>Total Area : {listing.sqft_area}</p>
        <p>{listing.beds} beds, {listing.baths} baths</p>
        <Link to={`/payment/${listing.listingId}`}>Payment</Link>
      </div>
    </>
  );
}