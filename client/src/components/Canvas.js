import React from 'react';
import './Canvas.css';
import { fabric } from '../lib/fabric.min.js';
import Guides from "@scena/guides";

class Canvasa extends React.Component {
  

   componentDidMount() {
      this.initCanvas();
      this.initGuides();
   }

   render() {
      return <div className="canvascontainer">
         <div className="canvas"><div className="ruler" style={{width: 60 + 'px' }}></div>
         <canvas id="c"></canvas></div>
      </div>
   }

   initGuides() {
      const guides = new Guides(document.querySelector('.ruler'), {
         type: "vertical",
     }).on("changeGuides", e => {
         console.log(e.guides);
     });
     
     
     let scrollX = 0;
     let scrollY = 0;
     window.addEventListener("resize", () => {
         guides.resize();
     });
     
     window.addEventListener("wheel", e => {
         scrollX += e.deltaX;
         scrollY += e.deltaY;
     
         guides.scrollGuides(scrollY);
         guides.scroll(scrollX);
     }); 
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
