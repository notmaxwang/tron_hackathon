import { useState } from 'react';
import './Chat.css'
import ChatBox from '../Component/ChatBox';

interface Message {
  sender: 'AI' | 'You';
  content: string;
}

const Chat: React.FC = () => {
  const [chats, setChats] = useState<{ [key: number]: Message[] }>({});
  const [nextId, setNextId] = useState<number>(1);
  const [activeChat, setActiveChat] = useState<number | null>(null);

  const handleNewChat = () => {
    const newId = nextId;
    setNextId(prevId => prevId + 1);
    setChats(prevChats => ({ ...prevChats, [newId]: [] }));
    setActiveChat(newId);
    console.log("Active new chat:", newId);
  };

  const handleDeleteChat = (chatId: number) => {
    setChats(prevChats => {
      const updatedChats = { ...prevChats };
      delete updatedChats[chatId];
      return updatedChats;
    });
    setNextId(prevId => Math.max(1, prevId - 1));
  };

  const handleSendMessage = (message: string) => {
    if (activeChat !== null) {
      setChats(prevChats => ({
        ...prevChats, // Updates chat state
        [activeChat]: [...(prevChats[activeChat] || []), { sender: 'You', content: message }]
      }));
    }
  };

  const handleSetActiveChat = (chatId: number) => {
    setActiveChat(chatId === activeChat ? null : chatId);
  };

  return (
    <div className='chat-container'>
      <aside className='side-menu'>
        <div className='side-menu-button' onClick={handleNewChat}>
          <span>+</span>
          New Chat
        </div>
        {Object.keys(chats).map(chatId => (
          <div
            key={chatId}
            className={`chat-list-item ${parseInt(chatId) === activeChat ? 'active-chat' : ''}`}
            onClick={() => handleSetActiveChat(parseInt(chatId))}
          >
            Chat {chatId}
            <div className="dropdown">
              <button className="btn btndropdown-toggle" type="button" data-bs-toggle="dropdown" aria-expanded="false">
                &#x2026;
              </button>
              <ul className='dropdown-menu'>
                <li><a className='dropdown-item' onClick={() => handleDeleteChat(parseInt(chatId))}>Delete</a></li>
              </ul>
            </div>
          </div>
        ))}
      </aside>
      <div className="chatbox-container">
        {activeChat !== null && (
          <ChatBox
            index={activeChat}
            messages={chats[activeChat]}
            onSendMessage={handleSendMessage}
            onCloseChat={() => handleDeleteChat(activeChat)}
          />
        )}
      </div>
    </div>
  );
};


export default Chat;