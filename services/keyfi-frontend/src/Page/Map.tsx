import { GoogleMap, useLoadScript, MarkerF } from '@react-google-maps/api';
import { QueryServiceClient } from '../../protos/query/query.client';
import { KeyValuePair, GetValuesRequest, GetValuesResponse } from '../../protos/query/query';
import React, { useState, useEffect } from 'react';
import { GrpcWebFetchTransport } from '@protobuf-ts/grpcweb-transport';
import Mapbox from '../Component/MapBox';
import './Map.css';

export default class Map extends React.Component {

  GOOGLE_MAP_API_KEY:any = '';

  listings = [{name: 'Ferry Building', position:{lat: 37.7955, lng: -122.3937,}},
    {name: 'Coit Tower', position:{lat: 37.8024, lng: -122.4058}}];

  listComponents:any = [];
  listName:any = [];

  addToList() {
    this.listings.forEach((listing, idx) => {
      this.listComponents.push(<MarkerF key={idx} position={listing.position} onClick={() => handleClick(listing.name)}/>);
      this.listName.push(<li key={idx}>{listing.name}</li>);
    });
  }
  
  


  componentWillMount() {
    this.addToList();
    this.getAPIKey();
  }

  async getAPIKey() {
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
      this.GOOGLE_MAP_API_KEY = response.keyValuePairs[0].value;
    }
  }

  handleClick(locationName:any) {
    console.log(locationName);
    console.log(this.GOOGLE_MAP_API_KEY);
  }

  
  render(){
    return (
      <div>
        <div className='topSection'>
          <Mapbox GOOGLE_MAP_API_KEY={this.GOOGLE_MAP_API_KEY} listComponents={this.listComponents}/>
          <div className='listing'>
            <p className='listingTitle'>Listings</p>
            <ul>{this.listName}</ul>
          </div>
        </div>
        
      </div>
    );
  }
  
}