import React from 'react';
// import logo from './logo.svg';
import './App.css';
import axios from 'axios';

function App() {
  const response = axios.get('/hello');

  return (
    <div className="App">
      <header className="App-header">
        {response.data}
      </header>
    </div>
  );
}

export default App;
