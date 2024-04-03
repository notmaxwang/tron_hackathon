import React, { useState } from 'react';



export default function Chat() {

  return(
    <>
      <form>
          <label>Prompt:</label> 
          <br/>
          <input type="text"/>
          <br/>
          <button>Submit</button>
        </form>
    </>
  );
}

