import React, { Component, useState } from 'react'

class Toolbox extends React.Component {
   
    constructor(props) {
        super(props);
        this.state = {
            myvar: 0
        };
   }
   
    render() {
        return <div className="box" style={{ width: '200px' }}>
        <p className="title is-5">Narrow column</p>
        <p className="subtitle">This column is only 200px wide.</p>
      </div>
   }
}

export default Toolbox;
