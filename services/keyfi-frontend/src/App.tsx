import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import  Header  from './Global/Header';
import Chat from './Page/Chat';
import Home from './Page/Home';
import Map from './Page/Map';
import './App.css'



function App() {
  

  return (
    <>
      <Header />
      <Router>
          <div>
            <Routes>
              <Route path="/" element={<Home />} />
              <Route path="/chat" element={<Chat />} />
              <Route path="/map" element={<Map />} />
            </Routes>
          </div>
        </Router>
    </>
  );
}

export default App
