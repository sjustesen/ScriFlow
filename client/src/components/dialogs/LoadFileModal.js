import React from 'react';
import 'bulma/css/bulma.min.css';
import EventService from '../../services/eventservice';
import SLADocument from '../../utils/sladocument';
import DocumentService from '../../services/documentservice';

class LoadFileModal extends React.Component {

    constructor(props) {
        super(props)

        this.eventService = new EventService();
        this.state = {
            show: false,
        };
    }

    render() {
        let showHideClassName = (this.state.show === true) ? "modal is-active" : "modal";
        return (<div className={showHideClassName}>
            <div className="modal-background"></div>
            <div className="modal-content">
                <p>this will be a filedialog of some sort, right now it just loads a demo (run the server and watch the console log)<br />
                    <button onClick={this.LoadFile}>Load File</button>
                </p>
            </div>
            <button className="modal-close is-large" onClick={this.HandleCloseModal} aria-label="close"></button>
        </div>)

    }

    componentDidMount() {
        this.eventService.subscribe('LayersPanelChanged', (event) => {
            console.dir('LayersPanelChanged')
        })

        this.eventService.subscribe('ColorPanelChanged', (event) => {
            console.dir('LayersPanelChanged')
        })

        this.eventService.subscribe('CanvasUpdated', (event) => {
            console.dir('CanvasUpdateEvent init')
        })
    }

    LoadFile = () => {
        // TODO: Put network actions in a service worker
        // TODO: Make a service class for file operations

        const docService = new DocumentService();
        let res = docService.loadFromUrl('http://localhost:8080/projects/load/1')
        res.then((file) => {

            if (file != null) {
                const doc = new SLADocument(file);

                // grab the initial state of the layers
                let page_objects = doc.getPageObjects();
                let colors = doc.getColors();
                let layers = doc.getLayers();
                
                // signal subscribers to update the layers panel
                this.eventService.publish('LayersPanelChanged', layers)
                this.eventService.publish('ColorPanelChanged', colors)
                this.eventService.publish('XMLDocumentLoaded')
                this.eventService.publish('CanvasUpdated', page_objects)

            }
        });

        this.HandleCloseModal()
    }

    HandleOpenModal = () => {
        this.setState({ show: true });
    }

    HandleCloseModal = () => {
        this.setState({ show: false })
    }
}

export default LoadFileModal;