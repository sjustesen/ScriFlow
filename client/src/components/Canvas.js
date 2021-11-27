import React, { Component } from 'react';
import './Canvas.css';

class Canvas extends React.Component {
   constructor(props) {
      super(props)
   }

   componentDidMount() {
     this.initCanvas();
   }

   initCanvas() {
      var canvasEl = document.getElementById(this.props.name);
      var ctx = canvasEl.getContext('2d');
      ctx.fillStyle = 'red';
      ctx.fillRect(100, 100, 20, 20);
   }

   render() {
		return <div className="canvascontainer">
               <div className="canvas"><canvas id="c"></canvas></div>
               </div>
   }
}

export default Canvas;
