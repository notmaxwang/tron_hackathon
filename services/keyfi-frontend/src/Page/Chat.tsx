import { useState } from 'react';
import './Chat.css'
import ChatBox from '../Component/ChatBox';

interface Message {
  sender: 'AI' | 'You';
  content: string;
}

// interface ChatBoxProps {
//   index?: number;
//   messages?: Message[];
//   onSendMessage?: (message: string) => void;
//   onCloseChat?: () => void;
//   onReceivedMessage?: (message: string) => void;
// }

const Chat = () => {
  const [chats, setChats] = useState<Message[][]>([[]]); 
  const [activeIndex, setActiveIndex] = useState<number>(0);
  
  const handleNewChat = () => {
    const newChatIndex = chats.length;
    console.log("Add new chat idx:", newChatIndex);
    setChats(prevChats => [...prevChats, []]); 
    setActiveIndex(newChatIndex);
  }
  
  const handleDeleteChat = (chatIndex: number) => {
    setChats(prevChats => prevChats.filter((_, index) => index !== chatIndex));
    setActiveIndex(0);
  };

  const handleSendMessage = (message: string, chatIndex: number) => {
    setChats(prevChats =>
      prevChats.map((chat, index) =>
        index === chatIndex ? [...chat, { sender: 'You', content: message }] : chat
      )
    );
  };

  const handleReceivedMessage = (message: string, chatIndex: number) => {
    setChats(prevChats =>
      prevChats.map((chat, index) =>
        index === chatIndex ? [...chat, { sender: 'AI', content: message }] : chat
      )
    );
  };

  const handleSetActiveChat = (chatIndex: number) => {
    console.log("Set active chat:", chatIndex);
    setActiveIndex(chatIndex);
  };

  return (
    <div className='chat-container'>
      <aside className='side-menu'>
        <div className='side-menu-button' onClick={handleNewChat}>
          <span>+</span>
          New Chat
        </div>
        {chats.map((_, index) => (
          <div 
            key={index} 
            className={`chat-list-item ${index === activeIndex ? 'active-chat' : ''}`}
            onClick={() => handleSetActiveChat(index)}
          >
            Chat {index + 1}
            <div className="dropdown">
              <button className="btn btndropdown-toggle" type="button" data-bs-toggle="dropdown" aria-expanded="false">
                &#x2026;
              </button>
              <ul className='dropdown-menu'>
                <li><a className='dropdown-item' onClick={() => handleDeleteChat(index)}>Delete</a></li>
              </ul>
            </div>
          </div>
        ))}
      </aside>
      <div className="chatbox-container">
        {/* instead of mapping, dynamically render whichever one is active */}
        {/* active chat and index, you can pass in index as index => you only show the one where teh active index is equal to the current index.... eerything else oyu hide */}
        {/* if active index = current index, show it */}
        {chats[activeIndex] && ( // Check if the active chat exists
          <ChatBox
            index={activeIndex}
            messages={chats[activeIndex]}
            onSendMessage={(message: string) => handleSendMessage(message, activeIndex)}
            onCloseChat={() => handleDeleteChat(activeIndex)}
            onReceivedMessage={(message: string) => handleReceivedMessage(message, activeIndex)}
          />
        )}
      </div>
    </div>
  );
};

export default Chat;