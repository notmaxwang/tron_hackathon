import { useState } from 'react'
import { KeyFiAIServiceClient} from '../protos/keyFiAI.client'
import { SinglePromptRequest } from '../protos/keyFiAI';
import { GrpcWebFetchTransport } from '@protobuf-ts/grpcweb-transport';
// import keyFiAIService from '../protos/keyFiAI_pb'
import './App.css'


function App() {
  const [textValues, setTextValues] = useState({
    prompt: '',
    response: ''
  })

  const handleUpdatePrompt = (e) => {
    setTextValues({ 
      ...textValues,
      prompt: e.target.value
    })
  }

  const handleButtonClick = async (e) => {
    let transport = new GrpcWebFetchTransport({
      baseUrl: "http://localhost:8080"
    });
    const client = new KeyFiAIServiceClient(transport);
    const request = SinglePromptRequest.create({
      prompt: textValues.prompt
    })
    const call = await client.singlePrompt(request);
    let response = await call.response
    let status = await call.status
    console.log("status: " + status)
    console.log(response);
    setTextValues({ 
      ...textValues,
      response: response.response
    })
  }

  const onFormSubmit = (e) => {
    e.preventDefault();
    console.log("refresh prevented");
  };

  return (
    <>
      <div>
        <form onSubmit={onFormSubmit}>
          <label>Prompt:</label> 
          <br/>
          <input type="text" onChange={handleUpdatePrompt} />
          <br/>
          <button onClick={handleButtonClick}>Submit</button>
        </form>
        <p>{textValues.response}</p>
      </div>
    </>
  )
}

export default App
