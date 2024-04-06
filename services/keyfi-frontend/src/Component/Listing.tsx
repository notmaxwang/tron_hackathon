import './Listing.css';
import React from 'react';

export default function Listing(props: any) {
  const listingName = props.name;
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
      <div className="listing-container" style={props.isPopup ? popupStyle : {}}>
        <p>{listingName}</p>
      </div>
    </>
  );
}