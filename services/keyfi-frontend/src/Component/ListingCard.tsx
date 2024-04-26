import './Listing.css';
import React, { useEffect } from 'react';
import { setRealEstateMarketContract, addHomeListing, fetchAllListings } from '../utils/tron.ts';

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

  useEffect(() => {
    setRealEstateMarketContract();
  }, [])

  return(
    <>
      <div className="listing-container" style={props.isPopup ? popupStyle : {}}>
        <p>{listing.name}</p>
        <button onClick={()=> addHomeListing('test2.url', '222 Market St', 500)}>test</button>
        <button onClick={()=> fetchAllListings()}>fetch</button>
      </div>
    </>
  );
}