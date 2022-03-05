import React from 'react';

class Color extends React.Component {
    // eslint-disable-next-line
    constructor(props) {
        super(props);
        this.props = props;
        this.state = {};
      
    }

    componentDidMount() {
       
    }

    render() {
        return (
        <div>{this.props.item.name}</div> 
        )
    }

}
export default Color;
