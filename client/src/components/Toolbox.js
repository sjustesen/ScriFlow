import React from 'react'

class Toolbox extends React.Component {

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
        </div>
    }
}

export default Toolbox;
