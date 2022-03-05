import React from 'react'
import Color from './Color';

class Toolbox extends React.Component {
    // eslint-disable-next-line
    constructor(props) {
        super(props)
        this.state = {
            activeTab: '',
            layers: [],
            colors: []
        }
    }

    componentDidMount() {
        window.addEventListener('LayersPanelChanged', (event) => {
            this.setState({ layers: event.detail })
        });

        window.addEventListener('ColorPanelChanged', (event) => {
            this.setState({ colors: event.detail })
        });

    }

    toggleTab() {

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
                <div className="LayerStack">
                    {this.state.layers.map((item, index) =>
                        <div key={index}>{item.properties.name}</div>)}
                </div>
            </div>
            <div className='ColorSwatches'>
                {this.state.colors.map((item, index) =>
                    <Color key={index} item={item}></Color>
                )}
            </div>
        </div>
    }

}

export default Toolbox;
