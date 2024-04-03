import React from 'react';
import { BrowserRouter as Router, Route, Routes, Link } from 'react-router-dom';
import  Header  from './Global/Header';
import Chat from './Page/Chat';
import Home from './Page/Home';

import './App.css';


function App() {
  

  return (
    <>
      <div>
        <Header />
        <Router>
          <div>
            <Header />
            <Routes>
              <Route path="/" element={<Home />} />
              <Route path="/chat" element={<Chat />} />
            </Routes>
          </div>
        </Router>
      </div>
    </>
  );
};

export default App;
