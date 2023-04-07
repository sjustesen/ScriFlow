import React from 'react';

class MasterPages extends React.Component {
    // eslint-disable-next-line
    constructor(props) {
        super(props);
        this.state = {};
      
    }

    componentDidMount() {
    }

    render() {
        return (
            <div className="box" style={{ width: '300px' }}>
            <div className="tabs is-centered">
                <ul>
                    <li className='listitem' id="layersTab">
                        <a href="#" onClick={()=> this.toggleTab('layers')}>
                            <span className="icon is-small"><i className="fas fa-image" aria-hidden="true"></i></span>
                            <span>Masterpages</span>
                        </a>
                    </li>
                </ul>
            </div>
            </div>
        )
    }

}
export default MasterPages;