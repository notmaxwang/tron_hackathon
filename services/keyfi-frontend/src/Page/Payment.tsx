import { useState, useEffect } from 'react';
import { setRealEstateMarketContract, 
         fetchAllListings,
         startSaleContract, 
         makeDownPayment, 
         makePayment,
         approveBuyer,
         approveSeller} from '../utils/tron.ts';
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
  const [transaction, setTransaction] = useState('');
  const [listingId, setListigId] = useState(0);
  let address = window.tronWeb ? window.tronWeb!.defaultAddress!.base58 : '';

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

  const fetchListings = async () => {
    console.log("starting fetch")
    await setRealEstateMarketContract();
    console.log("finished init")
    let storeListings = await fetchAllListings();
    for(let i = 0; i < storeListings.length; i++) {
      if(storeListings[i].streetAddress === listing.address){
        setListigId(i);
      };
    }
    console.log("finished fetch")
  }

  useEffect(() => {
    fetchListings();
    let backendCall = async() => {
      await makeCallToBackend();
    }
    backendCall();

    fetch(`https://nileapi.tronscan.org/api/transaction?count=true&limit=10&address=${address}&sort=-timestamp`)
    .then((res) => res.json())
    .then((res) => {
      if (res) {
        console.log(res.data);
        setTransaction(res.data[0].hash);
      }
    })
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
          <button onClick={() => startSaleContract(listingId, listing.price)}>Make Offer</button>
          <div className='buttons'>
            <button onClick={handlePrevStep}>Previous</button>
            <button onClick={handleNextStep}>Next</button>
          </div>
        </div>
      )}
      {step === 3 && (
        <div className="step-container">
          <h2>Step 3: Down Payment</h2>
          <p>This is your Transaction Hash: {transaction}</p>
          <button onClick={() => makeDownPayment(listingId)}>Make DownPayment</button>
          <div className='buttons'>
            <button onClick={handlePrevStep}>Previous</button>
            <button onClick={handleNextStep}>Next</button>
          </div>
        </div>
      )}
      {step === 4 && (
        <div className="step-container">
          <h2>Step 4: Approvals</h2>
          <button onClick={() => approveBuyer(listingId)}>Buyer Approval</button>
          <button onClick={() => approveSeller(listingId)}>Seller Approval</button>
          <div className='buttons'>
            <button onClick={handlePrevStep}>Previous</button>
            <button onClick={handleNextStep}>Next</button>
          </div>
        </div>
      )}
      {step === 5 && (
        <div className="step-container">
          <h2>Step 4: Payments</h2>
          <button onClick={() => makePayment(0x01)}>Make Payment</button>
          <div className='buttons'>
            <button onClick={handlePrevStep}>Previous</button>
            <button onClick={handleNextStep}>Next</button>
          </div>
        </div>
      )}
    </div>
  );
};

export default PaymentPage;

