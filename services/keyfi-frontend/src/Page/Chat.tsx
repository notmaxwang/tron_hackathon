import React, { useState } from 'react';


export default function Chat() {
  const LOCAL_HOST_URL = 'http://localhost:8080/';
  const [promptValue, setPromptValue] = useState('');
  const [responseValue, setResponseValue] = useState('');

  function onSubmit(e: any) {
    e.preventDefault();
    let value = {prompt: promptValue};
    let data = new FormData();
    data.append('prompt', promptValue);
    postJSON(value);
  }

  async function postJSON(data:any) {
    try {
      let headers = new Headers();
      headers.append('Content-Type', 'application/json');
      console.log(data);
      const response = await fetch(LOCAL_HOST_URL + 'simplePrompt', {
        method: "POST", 
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(data),
      });
  
      const result = await response.json();
      console.log("Success:", result);
    } catch (error) {
      console.error("Error:", error);
    }
  }

  function onInputChange(e: any) {
    setPromptValue(e.target.value);
  }

  return(
    <>
      <form onSubmit={onSubmit}>
        <label>Prompt:</label> 
        <br/>
        <input type="text" onChange={onInputChange}/>
        <br/>
        <button>Submit</button>
      </form>
      <p>{promptValue}</p>
    </>
  );
}

