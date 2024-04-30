import './ListingCard.css';
import React from 'react';

export default function ListingCard(props: any) {
  const listing = props.listing;
  const x = props.x;
  const y = props.y;

  const popupStyle: React.CSSProperties = {
    position: 'relative',
    left: x - 55 + 'px',
    top: y - 40 + 'px',
    border: '1px solid black',
  };

  return(
    <>
      <div className="listing-card-container" style={props.isPopup ? popupStyle : {}}>
        <img src="" alt="listing image" />
        <p>{listing.price}</p>
        <p>{listing.address}, {listing.city}, {listing.state}, {listing.zip_code}</p>
        <p>Total Area : {listing.sqft_area}</p>
        <p>{listing.beds} beds, {listing.baths} baths</p>
      </div>
    </>
  );
}