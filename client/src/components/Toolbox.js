import React from 'react'
import { Button } from 'react-bulma-components';

class Toolbox extends React.Component {
    constructor( prop ) {
        super(prop)
        this.state = {
            show : false
        }
    }
    
    HandleOpenModal = () => {
        this.setState({show : true});
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
            <div id="layers"><Button onClick={this.HandleOpenModal}>Modal</Button></div>
        </div>
    }
}

export default Toolbox;
