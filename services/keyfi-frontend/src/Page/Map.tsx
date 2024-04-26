import './Map.css';
import React, {useState, useEffect, useRef, useMemo} from 'react';
import Sparkle from '../assets/sparkle.png'
import mapboxgl from 'mapbox-gl';
import ListingCard from '../Component/ListingCard';
import Map, { Popup, Marker } from 'react-map-gl'
import 'mapbox-gl/dist/mapbox-gl.css'; 

export default function MapComponent() {
  const MAPBOX_MAP_API_KEY:any = process.env.MAPBOX_MAP_KEY;
  mapboxgl.accessToken = MAPBOX_MAP_API_KEY;
  const mapContainer = useRef(null);
  const map:any = useRef(null);
  const [lng, setLng] = useState(-122.4);
  const [lat, setLat] = useState(37.76);
  const [zoom, setZoom] = useState(11);
  const [showPopup, setShowPopup] = useState<boolean>(true);
  let listings = [{name: 'Ferry Building', position:{lat: 37.7955, lng: -122.3937,}},
  {name: 'Coit Tower', position:{lat: 37.8024, lng: -122.4058}}];
  let listName:any = [];

  listings.forEach((listing, idx) => {
    listName.push(<ListingCard key={idx} listing={listing} />)
  });

  const popup = useMemo(() => {
    return new mapboxgl.Popup().setText('Hello world!');
  }, [])

  useEffect(() => {
    if (!map.current && mapContainer.current) { // Check if map is not initialized and mapContainer exists
      map.current = new mapboxgl.Map({
        container: mapContainer.current,
        style: 'mapbox://styles/mapbox/streets-v12',
        center: [lng, lat],
        zoom: zoom
      });
      map.current.on('move', () => {
        setLng(map.current.getCenter().lng.toFixed(4));
        setLat(map.current.getCenter().lat.toFixed(4));
        setZoom(map.current.getZoom().toFixed(2));
      });  
    }
  });

  return (
    <div className='map-page-container'>
      <div className='search-bar-container'>
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
      <div className='bottom-section'>
        <div className='listing'>
          <p className='listingTitle'>Listings</p>
          <ul className='listings-container'>{listName}</ul>
        </div>
        <div className='map-container'>
        <Map
          mapboxAccessToken={MAPBOX_MAP_API_KEY}
          initialViewState={{
            longitude: -122.4,
            latitude: 37.76,
            zoom: 11,
          }}
          mapStyle="mapbox://styles/mapbox/streets-v9"
        >
          {showPopup && (
            <Marker longitude={-122.39} latitude={37.79}
              anchor="bottom"
              popup={popup}>
              TEST TEST
            </Marker>)}
        </Map>
        </div>
      </div>
    </div>
  );
}