import React from 'react';
import 'bulma/css/bulma.min.css';


class LoadFileModal extends React.Component {
    constructor(props) {
        super(props)
        this.state = { show: false };
    }

    componentDidMount() {
    }

    LoadDemoFile = () => {
        fetch('http://localhost:8080/projects/load/1')
        .then( response => response.text())
        .then(data => console.log(data))
        this.HandleCloseModal()
    }

    HandleOpenModal = () => {
        this.setState({ show: true });
    }

    HandleCloseModal = () => {
        this.setState({ show: false })
    }

    render() {
        let showHideClassName = (this.state.show === true) ? "modal is-active" : "modal";
        return (<div className={showHideClassName}>
            <div className="modal-background"></div>
            <div className="modal-content">
              <p>this will be a filedialog of some sort, right now it just loads a demo (run the server and watch the console log)<br />
              <button onClick={this.LoadDemoFile}>Load File</button>
              </p>
            </div>
            <button className="modal-close is-large" onClick={this.HandleCloseModal} aria-label="close"></button>
        </div>)

    }
}

export default LoadFileModal;