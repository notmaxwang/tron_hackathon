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
        <img className='listing-image' src={images[0]} alt="listing image" />
        <p className='listing-price'>${listing.price}</p>
        <p className='listing-sf'>
          {listing.area}
          <p className='listing-sqf'>sq/ft</p>
        </p> 
        <p className='listing-addy'>{listing.address}, {listing.city}, {listing.state}, {listing.zip_code}</p>
        <Link className='listing-payment' to={`/payment/${listing.listingId}`}>Payment</Link>
      </div>
    </>
  );
}