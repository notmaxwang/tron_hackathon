import './ListingCard.css';
import React, { useState, useEffect} from 'react';
import { Link } from 'react-router-dom';

export default function ListingCard(props: any) {
  const listing = props.listing;
  const x = props.x;
  const y = props.y;
  const [images, setImages] = useState<any>([]);
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

  const popupStyle: React.CSSProperties = {
    position: 'relative',
    left: x - 55 + 'px',
    top: y - 40 + 'px',
    border: '1px solid black',
  };

  return(
    <>
      <div className="listing-card-container" style={props.isPopup ? popupStyle : {}}>
        <img className='listing-thumbnail' src={images[0]} alt="listing image" />
        <div className='listing-pd'>
          <p className='listing-price'>${listing.price} <Link className='listing-detail'to={`/listing/${listing.listingId}`}>View Details</Link></p>
        </div>
        <p className='listing-area'>
          {listing.area}
          <p className='listing-sqft'>sq/ft</p>
        </p> 
        <p className='listing-address'>{listing.address}, {listing.city}, {listing.state}, {listing.zipcode}</p>
      </div>
    </>
  );
}