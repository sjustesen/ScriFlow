import React from 'react'
import Color from '../Color';
import LayerElement from './LayerElement';

class LayersPanel extends React.Component {
    // eslint-disable-next-line
    constructor(props) {
        super(props)
        this.state = {
            activeTab: 'layersTab',
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

    toggleTab(tabId) {
        // first hide the other tabs
        ['layers','swatches'].forEach(element => {
            let tabToHide = document.getElementById(element);
            tabToHide.style.display='none';
            document.querySelector('#'+element+'Tab').classList.remove('is-active');
        });
        
        document.querySelector('#'+tabId).style.display='block';
        document.querySelector('#'+tabId+'Tab').classList.add('is-active');
    }

    HandleOpenModal = () => {
        this.ModalElement.current.HandleOpenModal();
    }

    render() {

        return <div className="box" style={{ width: '300px' }}>
            <div className="tabs is-centered">
                <ul>
                    <li className='listitem' id="layersTab">
                        <a href="#" onClick={()=> this.toggleTab('layers')}>
                            <span className="icon is-small"><i className="fas fa-image" aria-hidden="true"></i></span>
                            <span>Layers</span>
                        </a>
                    </li>
                    <li className='listitem' id="swatchesTab">
                        <a href="#" onClick={()=> this.toggleTab('swatches')}>
                            <span className="icon is-small"><i className="fas fa-music" aria-hidden="true"></i></span>
                            <span>Colors</span>
                        </a>
                    </li>
                </ul>
            </div>
            <div id="layers">
                <div className="LayerStack">
                    {this.state.layers.map((item, index) =>
                        <LayerElement key={index} item={item}></LayerElement>
                    )}
                </div>
            </div>
            <div id="swatches" style={{display: 'none'}}>
            <div className='ColorSwatches'>
                {this.state.colors.map((item, index) =>
                    <Color key={index} item={item}></Color>
                )}
            </div>
            </div>
        </div>
    }

}

export default LayersPanel;
