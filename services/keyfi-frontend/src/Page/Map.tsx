import {  MarkerF } from '@react-google-maps/api';
import Listing from '../Component/Listing';
import { GoogleMap, useLoadScript } from '@react-google-maps/api';
import './Map.css';
import { useState } from 'react';


export default function Map() {
  const [popupValue, setPopupValue] = useState<any []>([]);
  const [showPopup, setShowPopup] = useState(false);
  const GOOGLE_MAP_API_KEY:any = process.env.REACT_APP_MAP_KEY;
  const libraries:any = ['places'];
  const mapContainerStyle = {
    width: '70vw',
    height: '70vh',
  };
  const center = {
    name:'center',
    position: {
    lat: 37.7937, // default latitude
    lng: -122.431297, // default longitude
  }};
  const { isLoaded, loadError } = 
  useLoadScript({
    googleMapsApiKey: GOOGLE_MAP_API_KEY,
    libraries,
  });

  if (loadError) {
    return <div>Error loading maps</div>;
  }

  if (!isLoaded) {
    return <div>Loading maps</div>;
  }

  let listings = [{name: 'Ferry Building', position:{lat: 37.7955, lng: -122.3937,}},
    {name: 'Coit Tower', position:{lat: 37.8024, lng: -122.4058}}];
  let listComponents:any = [];
  let listName:any = [];


  listings.forEach((listing, idx) => {
    listComponents.push(
      <MarkerF 
        key={idx} 
        position={listing.position} 
        onClick={(evt) => handleClickMarker(evt, listing)}>
          {popupValue[2] === listing.name ? <Listing x={popupValue[0]} y={popupValue[1]} name={popupValue[2]} isPopup={true} className='popupListing' /> : <></>}
      </MarkerF>);
    listName.push(<Listing key={idx} name={listing.name} />);
  });


  function handleClickMarker(event:any, listing:any) {
    console.log(listing.name);
    // const { screenX, screenY } = event;
    console.log(event.domEvent, event.domEvent.clientX, event.domEvent.clientY);
    setPopupValue([event.domEvent.clientX, event.domEvent.clientY, listing.name ]);;
    setShowPopup(true);
  }

  
  return (
    <div>
      <div className='topSection'>
        <GoogleMap
          mapContainerClassName='map'
          mapContainerStyle={mapContainerStyle}
          zoom={13}
          center={center.position}
          >
            {listComponents}
        </GoogleMap>
        <div className='listing'>
          <p className='listingTitle'>Listings</p>
          <ul className='listings-container'>{listName}</ul>
        </div>
      </div>
    </div>
  );
}