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
    lat: 37.7937, // default latitude
    lng: -122.431297, // default longitude
  };

  let GOOGLE_MAP_API_KEY:string = '';

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

  function handleClick() {
    console.log('popup window for listing');
    console.log(GOOGLE_MAP_API_KEY);
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
          <p className='listingTitle'>Listings</p>
        </div>
      </div>
      
    </div>
  );
}