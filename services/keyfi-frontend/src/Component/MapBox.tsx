import { GoogleMap, useLoadScript } from '@react-google-maps/api';

export default function Mapbox(props : any){
  const GOOGLE_MAP_API_KEY:any = process.env.REACT_APP_MAP_KEY;
  const listComponents = props.listComponents;
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

  return isLoaded ? (<GoogleMap
    mapContainerClassName='map'
    mapContainerStyle={mapContainerStyle}
    zoom={13}
    center={center.position}
    >
      {listComponents}
    </GoogleMap>) : <></>;
}