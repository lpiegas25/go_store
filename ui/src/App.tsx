import React from 'react';
import logo from './logo.svg';
import './App.css';
import axios from 'axios';

function App() {

  const fetchClient = async () => {
    const endpoint = 'http://localhost:8000/api/v1/clients/1';
    const result = await axios({
      method: 'GET',
      url: endpoint,
      responseType: 'json'
    });
    console.log(result.data);
  }

  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>
          Edit <code>src/App.tsx</code> and save to reload.
        </p>
        <a
          className="App-link"
          href="https://reactjs.org"
          target="_blank"
          rel="noopener noreferrer"
        >
          Learn React
        </a>
        <button onClick={fetchClient}>Fetch</button>
      </header>
    </div>
  );
}

export default App;
