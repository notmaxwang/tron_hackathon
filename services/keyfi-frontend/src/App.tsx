import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import  Navbar  from './Global/Navbar';
import Chat from './Page/Chat';
import Home from './Page/Home';
import MapComponent from './Page/Map';
import Wallet from './Page/Wallet';
import PaymentPage from './Page/Payment';
import ListingPage from './Page/ListingPage';
import 'bootstrap/dist/css/bootstrap.min.css';
import 'bootstrap/dist/js/bootstrap.bundle.min.js'
import 'bootstrap-icons/font/bootstrap-icons.css'
import './App.css'

import {useEffect, useState} from 'react';



function App() {
  
  const [walletAddress, setWalletAddress] = useState(' ');

  useEffect(() => {
    if (walletAddress) {
      console.log(walletAddress);
    }
  })
  return (
    <>
      <Navbar />
      <Router>
          <div>
            <Routes>
              <Route path="/" element={<Home />} />
              <Route path="/wallet" element={<Wallet setWalletAddress={setWalletAddress}/>} />
              <Route path="/chat" element={<Chat />} />
              <Route path="/map" element={<MapComponent />} />
              <Route path="/payment/:id" element={<PaymentPage walletAddress={walletAddress}/>} />
              <Route path="/listing/:id" element={<ListingPage />} />
            </Routes>
          </div>
        </Router>
    </>
  );
}

export default App
