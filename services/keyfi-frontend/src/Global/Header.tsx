import { useState } from "react";
import './Header.css';

export default function Header() {
  return(
    <header>
      <h1>Keyfi</h1>
      <nav>
        <ul>
          <li><a href="/">Home</a></li>
          <li><a href="/chat">Chat</a></li>
        </ul>
      </nav>
    </header>
  );
}

