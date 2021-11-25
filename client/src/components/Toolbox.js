import React, { Component, useState } from 'react'

class Toolbox extends React.Component {
   
    constructor(props) {
        super(props);
        this.state = {
            myvar: 0
        };
   }
   
    render() {
        return <div class="toolbox">toolbox</div>
   }
}

export default Toolbox;
