import React from 'react';
import './App.css';

import Topmenu from './components/navigation/Topmenu';
import Header from './components/Header.js';
import Toolbox from './components/Toolbox.js';
import Canvasa from './components/Canvas.js';
import LoadFileModal from './components/dialogs/LoadFileModal';

function App() {
  return (
      <div className="App">
        <Topmenu boxed="false" />
        <Header />
        <section className="columns">
          <div className="column is-10"><Canvasa name="c" /></div>
          <div className="column"><Toolbox /></div>
        </section>
        <LoadFileModal />
      </div>

      
  );
}

export default App;
