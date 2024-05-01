import { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
import  Listing  from '../Component/Listing';
import { ListingServiceClient } from '../../protos/listing/listing.client';
import { GetListingDetailRequest }  from '../../protos/listing/listing'
import { GrpcWebFetchTransport } from '@protobuf-ts/grpcweb-transport';

export default function ListingPage () {

  const { id } = useParams();
  const [listing, setListing] = useState<any>(null);

  const makeCallToBackend = async () => {
    let transport = new GrpcWebFetchTransport({
      baseUrl: "http://ec2-34-236-81-43.compute-1.amazonaws.com:8080"
    });
    const client = new ListingServiceClient(transport);
    const request = GetListingDetailRequest.create({
      listingId: id,
    })
    const call = client.getListingDetail(request);
    let response = await call.response;
    console.log(response);
    setListing(response.listing);
  }

  useEffect(() => {
    let backendCall = async() => {
      await makeCallToBackend();
    }
    backendCall();
  }, [])

  return(<>
      {listing&&<Listing listing={listing}/>}
  </>);
}