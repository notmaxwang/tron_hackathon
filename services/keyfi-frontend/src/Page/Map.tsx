import { GoogleMap, useLoadScript, MarkerF } from '@react-google-maps/api';
import { QueryServiceClient } from '../../protos/query/query.client';
import { KeyValuePair, GetValuesRequest, GetValuesResponse } from '../../protos/query/query';
import { useEffect } from 'react';
import { GrpcWebFetchTransport } from '@protobuf-ts/grpcweb-transport';
import './Map.css';

export default function Map() {
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

  let GOOGLE_MAP_API_KEY:string = '';

  let listings = [{name: 'Ferry Building', position:{lat: 37.7955, lng: -122.3937,}},
    {name: 'Coit Tower', position:{lat: 37.8024, lng: -122.4058}}];

  const listComponents:any = [];
  listings.forEach((listing, idx) => {
    listComponents.push(<MarkerF key={idx} position={listing.position} />);
  });

  const listName:any = [];
  listings.forEach((listing, idx) => {
    listName.push(<li key={idx}>{listing.name}</li>);
  });


  useEffect(() => {
    async function getAPIKey() {
      let transport = new GrpcWebFetchTransport({
        baseUrl: "http://localhost:8080"
      });
      const client = new QueryServiceClient(transport);
      const request = GetValuesRequest.create({
        keys: ['GOOGLE_MAPS_KEY']
      })
      const call = await client.getValues(request);
      let response = await call.response;
      let status = await call.status;
      console.log("status: " + status)
      if(response.keyValuePairs){
        GOOGLE_MAP_API_KEY = response.keyValuePairs[0].value;
      }
    }
    getAPIKey();
  }, [])

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

  function handleClick(locationName:any) {
    console.log(locationName);
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
          <MarkerF position={center.position} onClick={handleClick(center.name)}/>
          {listComponents}
        </GoogleMap>
        <div className='listing'>
          <p className='listingTitle'>Listings</p>
          <ul>{listName}</ul>
        </div>
      </div>
      
    </div>
  );
}