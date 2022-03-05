import React from 'react';

class Layer extends React.Component {
    // eslint-disable-next-line
    constructor(props) {
        super(props);
        this.state = {};
      
    }

    componentDidMount() {
    }

    render() {
        return (
        <div className='togglebutton'><div>{this.props.item.key}</div></div> 
        )
    }

}
export default Layer;
