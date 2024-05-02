import { useState, useEffect } from 'react';
import { setRealEstateMarketContract,  } from '../utils/tron.ts';
import Listing from '../Component/Listing.tsx';
import { useParams } from 'react-router-dom';
import { ListingServiceClient } from '../../protos/listing/listing.client';
import { GetListingDetailRequest }  from '../../protos/listing/listing'
import { GrpcWebFetchTransport } from '@protobuf-ts/grpcweb-transport';

import './Payment.css'; // Import your CSS file

const PaymentPage = () => {
  const [step, setStep] = useState(1); // Default step is 1
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

  const handleNextStep = () => {
    setStep(step + 1);
  };

  const handlePrevStep = () => {
    setStep(step - 1);
  };

  useEffect(() => {
    setRealEstateMarketContract();
    let backendCall = async() => {
      await makeCallToBackend();
    }
    backendCall();
  }, [])

  return (
    <div className="payment-page">
      {step === 1 && (
        <div className="step-container">
          <h2>Step 1: Review Listing</h2>
          {listing&&<Listing notIsListing={true} listing={listing}/>}
          <button onClick={handleNextStep}>Next</button>
        </div>
      )}
      {step === 2 && (
        <div className="step-container">
          <h2>Step 2: Payment</h2>
          <button onClick={handlePrevStep}>Previous</button>
          <button onClick={() => console.log('test')}>sign</button>
          <button onClick={handleNextStep}>Next</button>
        </div>
      )}
      {step === 3 && (
        <div className="step-container">
          <h2>Step 3: Down Payment</h2>
          {/* Content for confirmation */}
          <button onClick={handlePrevStep}>Previous</button>
        </div>
      )}
    </div>
  );
};

export default PaymentPage;

