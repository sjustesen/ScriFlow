import React from 'react';
import './Canvas.css';

import { fabric } from '../lib/fabric.min.js';

import SFCanvasLib from '../core/drawing/canvaslib'

class SFCanvas extends React.Component {

   componentDidMount() {
      this.sfcanvas = new SFCanvasLib();

   }

   render() {
      return <div className="canvascontainer">

         <div className="ruler" style={{ width: 60 + 'px' }}></div>
         <div className="innerframe">
            <div className="rulerHorizontal" style={{ height: 50 + 'px', minWidth: 100 + '%' }}></div>
            <div className="canvas">
               <canvas id="c"></canvas>
            </div>
         </div>
      </div>
   }

}

export default SFCanvas;
