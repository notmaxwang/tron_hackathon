import React, { useState, useEffect } from 'react';
import './Chat.css'

const Chat: React.FC = () => {
  const [ws, setWs] = useState<WebSocket | null>(null);
  const [messages, setMessages] = useState<string[]>([]);
  const [inputValue, setInputValue] = useState<string>('');

  useEffect(() => {
    // Create a new WebSocket connection when the component mounts
    const newWs = new WebSocket('ws://localhost:50052');

    newWs.onopen = () => {
      console.log('WebSocket connected');
    };

    newWs.onmessage = (event) => {
      // Add received message to the messages state
      setMessages((prevMessages) => [...prevMessages, event.data]);
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
      // Send the message through the WebSocket connection
      ws.send(inputValue);
      setInputValue('');
    }
  };

  return (
    <div>
      <h2>Chat</h2>
      <div>
        {messages.map((message, index) => (
          <div key={index}>{message}</div>
        ))}
      </div>
      <input
        type="text"
        value={inputValue}
        onChange={(e) => setInputValue(e.target.value)}
        placeholder="Type a message..."
      />
      <button onClick={sendMessage}>Send</button>
    </div>
  );
};

export default Chat;