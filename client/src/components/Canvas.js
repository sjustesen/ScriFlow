import React from 'react';
import './Canvas.css';
import { fabric } from '../lib/fabric.min.js';

class Canvasa extends React.Component {
  

   componentDidMount() {
      this.initCanvas();
   }

   render() {
      return <div className="canvascontainer">
         <div className="canvas"><canvas id="c"></canvas></div>
      </div>
   }

   initCanvas() {
      let parent = document.querySelector('.canvas');
      
      const canvas = new fabric.Canvas('c');
      canvas.setHeight(parent.clientHeight);
      canvas.setWidth(parent.clientWidth);

      window.addEventListener('resize', (e) => {
         canvas.setHeight(parent.clientHeight);
         canvas.setWidth(parent.clientWidth);
   
      })

      var rect = new fabric.Rect({
         left: 100,
         top: 100,
         fill: 'red',
         width: 20,
         height: 20
       });

       canvas.add(rect);
   }


}

export default Canvasa;
