import './Map.css';
import {useState, useEffect, useRef, useMemo} from 'react';
import Sparkle from '../assets/sparkle.png'
import mapboxgl from 'mapbox-gl';
import ListingCard from '../Component/ListingCard';
import Map, { Marker } from 'react-map-gl'
import 'mapbox-gl/dist/mapbox-gl.css'; 
// import { setRealEstateMarketContract, fetchAllListings } from '../utils/tron';
import { GetListingsRequest }  from '../../protos/listing/listing'
import { ListingServiceClient } from '../../protos/listing/listing.client'
import { GrpcWebFetchTransport } from '@protobuf-ts/grpcweb-transport';
import MapMarker from '../assets/marker.webp';

export default function MapComponent() {
  const MAPBOX_MAP_API_KEY:any = process.env.MAPBOX_MAP_KEY;
  mapboxgl.accessToken = MAPBOX_MAP_API_KEY;
  const mapContainer = useRef(null);
  const map:any = useRef(null);
  const [lng, setLng] = useState(-122.4);
  const [lat, setLat] = useState(37.76);
  const [zoom, setZoom] = useState(11);
  const [showPopup, setShowPopup] = useState<boolean>(true);
  const [listings, setListings] = useState<any>([]);
  const [markerList, setMarkerList] = useState<any>([]);

  const makeCallToBackend = async () => {
    let transport = new GrpcWebFetchTransport({
      baseUrl: "http://ec2-34-236-81-43.compute-1.amazonaws.com:8080"
    });
    const client = new ListingServiceClient(transport);
    const request = GetListingsRequest.create({
      cities: ["San Francisco"]
    })
    const call = client.getListings(request);
    let response = await call.response;
    console.log('second',response);
    setListings(response.listings);
  }

  const popup = useMemo(() => {
    return new mapboxgl.Popup().setText('Hello world!');
  }, [])

  useEffect(() => {
    console.log('first');
    let backendCall = async() => {
      await makeCallToBackend();
      console.log('third', listings);
    }
    backendCall();
    setShowPopup(true);
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
          <ul className='listings-container'>{listings.map((listing, idx) => (
            <ListingCard key={idx} 
                         listing={listing} 
                         x={listing.coordLat} 
                         y={listing.coordLong} />
          ))}</ul>
        </div>
        <div className='map-container'>
        <Map
          mapboxAccessToken={MAPBOX_MAP_API_KEY}
          initialViewState={{
            longitude: -122.44,
            latitude: 37.76,
            zoom: 11,
          }}
          mapStyle="mapbox://styles/mapbox/streets-v9"
        >
          {showPopup && (
            <>
            {listings.map((listing, idx) => (
            <Marker longitude={listing.coordLong} latitude={listing.coordLat}
              anchor="bottom"
              popup={popup}
              onClick={() => console.log('test')}>
                <img className='marker' src={MapMarker} alt='marker'/>
            </Marker>
          ))}
          </>)}
        </Map>
        </div>
      </div>
    </div>
  );
}