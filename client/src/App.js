import React from 'react';
import './App.css';

import Topmenu from './components/navigation/Topmenu';
import Header from './components/Header.js';
import Toolbox from './components/Toolbox.js';
import Canvasa from './components/Canvas.js';
import LoadFileModal from './components/dialogs/LoadFileModal';

class App extends React.Component {
  constructor(props) {
    super(props);

    this.ModalElement = React.createRef();
  }

  HandleOpenModal = () => {
    this.ModalElement.current.HandleOpenModal();
  }

  render() {
    return (
      <div className="App">
        <Topmenu boxed="false" />
        <Header />
        <section className="columns">
          <div className="column is-10"><Canvasa name="c" /></div>
          <div className="column">
            <Toolbox />
          <button onClick={this.HandleOpenModal}>Demo Load file</button>
          </div>
        </section>
        <LoadFileModal ref={this.ModalElement} />
      </div>
    )
  }
}

export default App;
