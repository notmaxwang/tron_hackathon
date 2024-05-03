import { useState, useEffect } from 'react';
import './ChatBox.css';
import { QueryServiceClient } from '../../protos/query/query.client';
import { GetValuesRequest } from '../../protos/query/query';
import { GrpcWebFetchTransport } from '@protobuf-ts/grpcweb-transport';
import Sparkle from '../assets/sparkle.png';

interface Message {
  sender: 'AI' | 'You';
  content: string;
}

export default function ChatBox(props) {
  const [ws, setWs] = useState<WebSocket | null>(null);
  const [showInterface, setShowInterface] = useState(true);
  const [messages, setMessages] = useState<Message[]>([]); //should be stored individually in the component
  const [inputValue, setInputValue] = useState<string>(''); // should be stored indiivudally in the component
  const [key, setKey] = useState("");

  console.log(props);
    const makeCallToBackend = async () => {
      let transport = new GrpcWebFetchTransport({
        baseUrl: "http://ec2-34-236-81-43.compute-1.amazonaws.com:8080"
      });
      const client = new QueryServiceClient(transport);
      const request = GetValuesRequest.create({
        keys: ["yaobin", "foo"]
      })
      const call = client.getValues(request);
      let response = await call.response;
      let status = await call.status;
      console.log("status: " + status)
      console.log(response.keyValuePairs);
    }
  
    useEffect(() => {
      
      makeCallToBackend();
  
      // Create a new WebSocket connection when the component mounts
      const newWs = new WebSocket('ws://ec2-34-236-81-43.compute-1.amazonaws.com:8081');
  
      newWs.onopen = () => {
        console.log('WebSocket connected');
      };
  
      newWs.onmessage = (event) => {
        // Add received message to the messages state
        setMessages((prevMessages) => [...prevMessages, {sender: 'AI', content: event.data}]);
      };
  
      newWs.onclose = () => {
        console.log('WebSocket disconnected');
      };
  
      // Update ws state with the new WebSocket connection
      setWs(newWs);
  
      // Close the WebSocket connection when the component unmounts
      return () => {
        newWs.close();
      };
    }, []);

  const sendMessage = () => {
    if (ws && inputValue.trim() !== '') {
      if (showInterface) {
        // Hide the initial interface when the user sends the first message
        setShowInterface(false);
      }
      // Send the message through the WebSocket connection
      ws.send(inputValue);
      setMessages((prevMessages) => [...prevMessages, {sender: 'You', content: inputValue}]);
      setInputValue('');
    }
  };

  const handleKeyPress = (event: React.KeyboardEvent<HTMLInputElement>) => {
    if (event.key == 'Enter') {
      if (ws && inputValue.trim() !== '') {
        if (showInterface) {
          // Hide the initial interface when the user sends the first message
          setShowInterface(false);
        }
        // Send the message through the WebSocket connection
        ws.send(inputValue);
        setMessages((prevMessages) => [...prevMessages, {sender: 'You', content: inputValue}]);
        setInputValue('');
      }
    }
  }

  return (
    <>
      <section className='chatbox'>
        <h2 className='chat-header'>
          <h2> 
            <h2></h2>
          </h2>
        </h2>
        {showInterface && (
          <div className="initial-interface">
            <div className='steve-ai'>
              Hi! I'm <span className='gradient-ai-text'>Steve.ai <img src={Sparkle} alt="" className="sparkle-interface" /></span>
            </div>
            <p className='steve-description'>Your AI-powered real estate agent. Ask me anything real estate and I’ll do my best to answer. I can help you with...</p>
            <ul className="flex-container">
              <li className="flex1-item">
                <span className='f1-item'>Finding your best home</span>
                <li className='small-item'>Find a place that fits you</li>
                <li className='small-item'>Neighborhood rating</li>
                <li className='small-item'>Home rating</li>
              </li>
              <li className="flex2-item">
                <span className='f2-item'>Contracts</span>
                <li className='small-item'>Generating optimal contracts</li>
                <li className='small-item'>Negotiating purchase price</li>
                <li className='small-item'>Verify MLS</li>
              </li>
              <li className="flex3-item">
                <span className='f3-item'>Mortgage and Loans</span>
                <li className='small-item'>Pre-verify loans</li>
                <li className='small-item'>Best mortgage plan for you</li>
                <li className='small-item'>Loan calculator</li>
              </li>
            </ul>
          </div>
        )}
        <div className='message-container'>
          {messages.map((message, index) => (
            <div key={index} className={message.sender === 'AI' ? 'ai-message-container' : 'user-message-container'}>
              {message.sender === 'AI' && (
                <div className='ai-message'>
                  <p className='sender'><img src={Sparkle} alt="" className="sparkle" /> Steve.ai</p> {message.content}
                </div>
              )}
              {message.sender === 'You' && (
                <div className='user-message'>
                  <p className='sender'>You</p> {message.content}
                </div>
              )}
            </div>
          ))}
        </div>
        <div className='input-container'>
          <input
            className='input-field'
            type="text"
            value={inputValue}
            onChange={(e) => setInputValue(e.target.value)}
            placeholder="Ask me anything..."
            onKeyDown={handleKeyPress}
          />
          <button className='input-button' onClick={sendMessage}>Send</button>
        </div>
      </section>
  </>
  );
};