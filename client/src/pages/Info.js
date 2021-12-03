import React from 'react';
import '../App.css';

import Topmenu from '../components/navigation/Topmenu';
import Header from '../components/Header.js';
import Toolbox from '../components/Toolbox.js';

function Info() {
  return (
      <div className="App">
        <Topmenu boxed="false" />
        <Header />
        <section className="columns">
          <div className="column is-10">test</div>
          <div className="column"><Toolbox /></div>
        </section>
      </div>
  );
}

export default Info;
