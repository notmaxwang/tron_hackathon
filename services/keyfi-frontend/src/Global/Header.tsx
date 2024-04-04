import { useState } from "react";
import './Header.css';

export default function Header() {
  return(
    <header>
      <p>KeyFi</p>
      <nav>
        <ul>
          <li><a href="/">Home</a></li>
          <li><a href="/chat">Chat</a></li>
          <li><a href="/map">Map</a></li>
        </ul>
      </nav>
    </header>
  );
}
