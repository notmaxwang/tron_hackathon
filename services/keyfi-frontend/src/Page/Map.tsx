import { MarkerF } from '@react-google-maps/api';
import Listing from '../Component/Listing';
import { GoogleMap, useLoadScript } from '@react-google-maps/api';
import './Map.css';
import React from 'react';
import { Unstable_Popup as BasePopup } from '@mui/base/Unstable_Popup';
import Sparkle from '../assets/sparkle.png'



export default function Map() {
  // const [currListing, setCurrListing] = useState(null);
  const [anchor, setAnchor] = React.useState<null | HTMLElement>(null);
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

  function handleClick(event:any) {
    setAnchor(anchor ? null : event.domEvent.currentTarget);
  }

  function reset() {
    setAnchor(null);
  }

  let open = Boolean(anchor);
  let id = open ? 'simple-popper' : undefined;

  function setNewListing(event:any) {
    console.log(event);
  }

  listings.forEach((listing, idx) => {
    listComponents.push(
        <MarkerF 
          key={idx} 
          position={listing.position} 
          clickable={true}
          label={listing.name}
          onClick={handleClick}
          onMouseOver={setNewListing}>
        </MarkerF>);
    listName.push(<Listing key={idx} name={listing.name} />);
  });
 
  return (
    <div className='container'>
      <div className='search-bar-container'>
        <div className='top-nav'>
          <div className='search-container'>
            <input className='home-search' type='search' placeholder="search thing" />
            <button className='ai-button'><img src={Sparkle} alt="" className="sparkle" />Ask AI</button>
          </div>
          <div className='filter-dropdown'>
            <select className='price-filter' id='price-filter'>
                <option className='select-item' value='select'>Price</option>
                <option className='select-item' value='price2'>Price 2</option>
                <option className='select-item' value='price3'>Price 3</option>
            </select>
            <select className='bedbath-filter' id='bedbath-filter'>
              <option className='select-item' value='select'>Beds & Baths</option>
              <option className='select-item' value='bedbath2'>BedBath 2</option>
              <option className='select-item' value='bedbath3'>BedBath 3</option>
            </select>
            <select className="style-filter">
              <option className='select-item' value="all">Housing Style</option>
              <option className='select-item' value="type1">Type 1</option>
              <option className='select-item' value="type2">Type 2</option>
            </select>
            <select className="filter">
              <option className='select-item' value="filter">Filter</option>
              <option className='select-item' value="filter1">Filter 1</option>
              <option className='select-item' value="filter2">Filter 2</option>
            </select>
          </div>
        </div>
      </div>
      <div className='bottom-container'>
        <div className='bottom-section'>
          <div className='listing'>
            <p className='listingTitle'>Listings</p>
            <ul className='listings-container'>{listName}</ul>
          </div>
          <GoogleMap
            mapContainerClassName='map'
            mapContainerStyle={mapContainerStyle}
            zoom={13}
            center={center.position}
            >
              {listComponents}
          </GoogleMap>
          <BasePopup id={id} open={open} anchor={anchor} className='popup-window'>
            <button onClick={reset}>X</button>
            <p>test</p>
          </BasePopup>
        </div>
      </div>
    </div>
  );
}