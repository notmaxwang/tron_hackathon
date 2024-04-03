import { useState, useEffect } from 'react'
import { KeyFiAIServiceClient} from '../../protos/keyFiAI.client'
import { SinglePromptRequest } from '../../protos/keyFiAI';
import { GrpcWebFetchTransport } from '@protobuf-ts/grpcweb-transport';
// import keyFiAIService from '../protos/keyFiAI_pb'
import './Chat.css'

export default function Chat() {

  const [messages, setMessages] = useState<string[]>([]);

  const [currentValue, setCurrentValue] = useState('');


  const handleUpdatePrompt = (e: any) => {
    setCurrentValue(e.target.value);
  }

  const handleButtonClick = () => {
    setMessages([...messages, currentValue]);
  }

  const onFormSubmit = async (e: any) => {
    e.preventDefault();
    let transport = new GrpcWebFetchTransport({
      baseUrl: "http://localhost:8080"
    });
    const client = new KeyFiAIServiceClient(transport);
    const request = SinglePromptRequest.create({
      prompt: currentValue
    })
    const call = await client.singlePrompt(request);
    let response = await call.response;
    let status = await call.status;
    console.log("status: " + status)
    console.log(response);
    setMessages([ 
      ...messages,
      response.response
    ])
    e.target.reset();
    console.log("refresh prevented");
  };

  return(
    <div className="chat-container">
      <div className="chat-messages">
        {messages.map((message, index) => (
          <div key={index} className="message">
            {message}
          </div>
        ))}
      </div>
      <form onSubmit={onFormSubmit}>
        <label>Type your Message:</label> 
        <br/>
        <input type="text" onChange={handleUpdatePrompt} />
        <br/>
        <button onClick={handleButtonClick}>Send</button>
      </form>
    </div>
  );
}