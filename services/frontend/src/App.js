import logo from './logo.svg';
import './App.css';

function App() {
  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <input name="myInput" />
        <Button
          onPress={submitPrompt}
          title="Submit Prompt"
          color="#841584"
          accessibilityLabel="Submits the prompt to Gemini"
        />
      </header>
    </div>
  );

  function submitPrompt() {
    
  }

}

export default App;
