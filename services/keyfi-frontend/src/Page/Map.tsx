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
    setListings(response.listings);
  }

  const popup = useMemo(() => {
    return new mapboxgl.Popup().setText('Hello world!');
  }, [])

  useEffect(() => {
    let backendCall = async() => {
      await makeCallToBackend();
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
          <input className='home-search' type='search' placeholder="Search for a house" />
          <button className='ai-button'><img src={Sparkle} alt="" className="sparkle" />Ask AI</button>
        </div>
        <div className='filter-dropdown'>
          <button className="btn btn-secondary dropdown-toggle" type="button" data-bs-toggle="dropdown" aria-expanded="false">
            Price
          </button>
          <ul className="dropdown-menu">
            <li><a className="dropdown-item" href="#">Price 1</a></li>
            <li><a className="dropdown-item" href="#">Price 2</a></li>
            <li><a className="dropdown-item" href="#">Price 3</a></li>
          </ul>
          <button className="btn btn-secondary dropdown-toggle" type="button" data-bs-toggle="dropdown" aria-expanded="false">
            Beds & Baths
          </button>
          <ul className="dropdown-menu">
            <li><a className="dropdown-item" href="#">Beds & Baths 1</a></li>
            <li><a className="dropdown-item" href="#">Beds & Baths 2</a></li>
            <li><a className="dropdown-item" href="#">Beds & Baths 3</a></li>
          </ul>
          <button className="btn btn-secondary dropdown-toggle" type="button" data-bs-toggle="dropdown" aria-expanded="false">
            Housing Style
          </button>
          <ul className="dropdown-menu">
            <li><a className="dropdown-item" href="#">Style 1</a></li>
            <li><a className="dropdown-item" href="#">Style 2</a></li>
            <li><a className="dropdown-item" href="#">Style 3</a></li>
          </ul>
          <button className="btn btn-secondary dropdown-toggle" type="button" data-bs-toggle="dropdown" aria-expanded="false">
            Filter
          </button>
          <ul className="dropdown-menu">
            <li><a className="dropdown-item" href="#">Filter 1</a></li>
            <li><a className="dropdown-item" href="#">Filter 2</a></li>
            <li><a className="dropdown-item" href="#">Filter 3</a></li>
          </ul>
        </div>
      </div>
      <div className='bottom-section'>
        <div className='listing'>
          <p className='listing-title'>Real Estate in San Francisco, California for Sale</p>
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
            {listings.map((listing) => (
            <Marker 
              longitude={listing.coordLong} 
              latitude={listing.coordLat}
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