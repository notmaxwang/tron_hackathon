import { GoogleMap, useLoadScript, MarkerF } from '@react-google-maps/api';
import './Map.css';

export default function Map() {
  const libraries:any = ['places'];
  const mapContainerStyle = {
    width: '70vw',
    height: '70vh',
  };
  const center = {
    lat: 37.7937, // default latitude
    lng: -122.431297, // default longitude
  };



  const { isLoaded, loadError } = useLoadScript({
    googleMapsApiKey: GOOGLE_MAP_API_KEY,
    libraries,
  });

  if (loadError) {
    return <div>Error loading maps</div>;
  }

  if (!isLoaded) {
    return <div>Loading maps</div>;
  }

  function handleClick() {
    console.log('popup window for listing');
  }

  return (
    <div>
      <div className='topSection'>
        <GoogleMap
          mapContainerClassName='map'
          mapContainerStyle={mapContainerStyle}
          zoom={13}
          center={center}
        >
          <MarkerF position={center} onClick={handleClick}/>
        </GoogleMap>
        <div className='listing'>
          <h3 className='listingTitle'>Listings</h3>
        </div>
      </div>
      
    </div>
  );
}