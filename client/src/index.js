import React from 'react';
import ReactDOM from 'react-dom';
import { HashRouter, Routes, Route } from "react-router-dom";

import './index.css';
import App from './App';
import Info from './pages/Info';

import reportWebVitals from './reportWebVitals';

// note to self: routes looks like this: localhost:3000/#/info
// consider using browser routing later on

ReactDOM.render(
  <React.StrictMode>
    <HashRouter>
    <Routes>
      <Route path="/" element={<App />} />
      <Route path="/info" element={<Info/>} />    
    </Routes>
    </HashRouter>
  </React.StrictMode>,
  document.getElementById('root')
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
