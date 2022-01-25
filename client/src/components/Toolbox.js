import React from 'react'
import EventService from '../services/eventservice';

class Toolbox extends React.Component {
    // eslint-disable-next-line
    constructor(props) {
        super(props)
        this.state = {
            layers: []
        }
    }

    componentDidMount() {
        window.addEventListener('LayersPanelChanged', (event) => {
            this.setState({ layers: event.detail })
        });

       
    }

    HandleOpenModal = () => {
        this.ModalElement.current.HandleOpenModal();
    }

    render() {
        
        return <div className="box" style={{ width: '300px' }}>
            <div className="tabs is-centered">
                <ul>
                    <li className="is-active">
                        <a href="#layers">
                            <span className="icon is-small"><i className="fas fa-image" aria-hidden="true"></i></span>
                            <span>Layers</span>
                            <div id="LayerStack">
                                <ul>
                                    {this.state.layers.forEach(item => <li>{item.name}</li>)}
                                </ul>
                            </div>
                        </a>
                    </li>
                    <li>
                        <a href="#colors">
                            <span className="icon is-small"><i className="fas fa-music" aria-hidden="true"></i></span>
                            <span>Colors</span>
                        </a>
                    </li>
                </ul>
            </div>
            <div id="layers">

            </div>
        </div>
    }

}

export default Toolbox;
