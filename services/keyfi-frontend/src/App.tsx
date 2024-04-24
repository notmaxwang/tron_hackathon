import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import  Header  from './Global/Header';
import Chat from './Page/Chat';
import Home from './Page/Home';
import MapComponent from './Page/Map';
import Wallet from './Page/Wallet';
import 'bootstrap/dist/css/bootstrap.min.css';
import 'bootstrap/dist/js/bootstrap.bundle.min.js'
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
              <Route path="/map" element={<MapComponent />} />
              <Route path="/wallet" element={<Wallet />} />
            </Routes>
          </div>
        </Router>
    </>
  );
}

export default App
