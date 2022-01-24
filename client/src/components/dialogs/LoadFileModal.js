import React from 'react';
import 'bulma/css/bulma.min.css';
import EventService from '../../services/eventservice';
import SLADocument from '../../utils/sladocument';
import DocumentService from '../../services/documentservice';

class LoadFileModal extends React.Component {

    constructor(props) {
        super(props)

        this.state = {
            show: false,
        };
    }

    componentDidMount() {
        this.eventService = new EventService();
        this.eventService.subscribe('UpdateLayerPanel', this.UpdateLayersPanelEvent)
    }

    LoadDemoFile = () => {
        const docService = new DocumentService();
        let response = docService.loadFromUrl('http://localhost:8080/projects/load/1')
        if (response.ok) {
            console.dir(response)
            const doc = new SLADocument(response.data);

            // the initial state of the layers
            var layers = doc.getLayers();
    
            // signal subscribers to update the layers panel
            this.eventService.publish('UpdateLayerPanel', layers)
            
        // TODO: move to service worker
        }

        this.HandleCloseModal()
    }

    HandleOpenModal = () => {
        this.setState({ show: true });
    }

    HandleCloseModal = () => {
        this.setState({ show: false })
    }

    UpdateLayersPanelEvent(data) {
        console.log('updating')
        console.log(data)
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