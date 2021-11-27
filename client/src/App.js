import React, { Component } from 'react';
import './App.css';

import Header from './components/Header.js';
import Toolbox from './components/Toolbox.js';
import Canvasa from './components/Canvas.js';

function App() {
  return (
    <div className="App">
	<Header />
    <section className="columns">
		<div className="column"><Canvasa name="c" /></div>
		<div className="column is-narrow"><Toolbox /></div>
    </section>
    </div>
  );
}

export default App;
