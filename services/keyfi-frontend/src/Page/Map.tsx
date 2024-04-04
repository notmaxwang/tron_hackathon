import { GoogleMap, useLoadScript, MarkerF } from '@react-google-maps/api';
import { QueryServiceClient } from '../../protos/query/query.client';
import { KeyValuePair, GetValuesRequest, GetValuesResponse } from '../../protos/query/query';
import { useState, useEffect } from 'react';
import { GrpcWebFetchTransport } from '@protobuf-ts/grpcweb-transport';
import Mapbox from '../Component/MapBox';
import './Map.css';

export default function Map() {

  let GOOGLE_MAP_API_KEY:string = '';

  let listings = [{name: 'Ferry Building', position:{lat: 37.7955, lng: -122.3937,}},
    {name: 'Coit Tower', position:{lat: 37.8024, lng: -122.4058}}];

  const listComponents:any = [];
  listings.forEach((listing, idx) => {
    listComponents.push(<MarkerF key={idx} position={listing.position} onClick={() => handleClick(listing.name)}/>);
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

  function handleClick(locationName:any) {
    console.log(locationName);
    console.log(GOOGLE_MAP_API_KEY);
  }

  
  return (
    <div>
      <div className='topSection'>
        <Mapbox GOOGLE_MAP_API_KEY={GOOGLE_MAP_API_KEY} listComponents={listComponents}/>
        <div className='listing'>
          <p className='listingTitle'>Listings</p>
          <ul>{listName}</ul>
        </div>
      </div>
      
    </div>
  );
}