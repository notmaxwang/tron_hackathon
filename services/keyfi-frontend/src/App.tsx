import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import  Header  from './Global/Header';
import Chat from './Page/Chat';
import Home from './Page/Home';
import Map from './Page/Map';
import Wallet from './Page/Wallet';
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
              <Route path="/wallet" element={<Wallet />} />
            </Routes>
          </div>
        </Router>
    </>
  );
}

export default App
