import React from 'react';

class LayerElement extends React.Component {
    // eslint-disable-next-line
    constructor(props) {
        super(props);
        this.state = {};
      
    }

    componentDidMount() {
    }

    render() {
        return (
        <div className='togglebutton'><div data-item={this.props.item.key}>{this.props.item.name}</div></div> 
        )
    }

}
export default LayerElement;
