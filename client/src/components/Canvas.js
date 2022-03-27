import React from 'react';
import './Canvas.css';
import { fabric } from '../lib/fabric.min.js';
import Guides from "@scena/guides";

class Canvasa extends React.Component {

  /* constructor(props) {
      super(props)
   } */

   componentDidMount() {
      this.initCanvas();
      this.initGuides();
      this.registerCanvasEvents();
   }

   registerCanvasEvents() {
      window.addEventListener('XmlDocumentLoaded', this.XMLDocumentLoaded)
      window.addEventListener('LayersPanelChanged', ()=> {
         console.log('Update Layer Panel signaled..')
      })
      window.addEventListener('CanvasUpdated', (page_objects) => {
         console.dir(page_objects)
      });

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

   initGuides() {
      const guide1 = new Guides(document.querySelector('.ruler'), {
         type: "vertical",
      }).on("changeGuides", e => {
         console.log(e.guides);
      });

      const guide2 = new Guides(document.querySelector('.rulerHorizontal'), {
         type: "horizontal",
      }).on("changeGuides", e => {
         console.log(e.guides);
      });


      let scrollX = 0;
      let scrollY = 0;
      window.addEventListener("resize", () => {
         guide1.resize();
         guide2.resize();
      });

      window.addEventListener("wheel", e => {
         scrollX += e.deltaX;
         scrollY += e.deltaY;

         guide1.scrollGuides(scrollY);
         guide1.scroll(scrollX);
      });
   }

   initCanvas() {
      let grandparent = document.querySelector('.canvascontainer');
      let parent = document.querySelector('.canvas');

      let width = grandparent.clientWidth - 60;

      parent.setAttribute('style', 'width: ' + width + 'px');

      const canvas = new fabric.Canvas('c');
      canvas.setHeight(parent.clientHeight);
      canvas.setWidth(parent.clientWidth);

      window.addEventListener('resize', (e) => {
         let width = grandparent.clientWidth - 60;
         parent.setAttribute('style', 'width: ' + width + 'px');
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

   // event
   XMLDocumentLoaded = (data) => {
      console.log('xmldoc loaded');
   }

   // event
   XMLDocumentUnloaded() {
      console.log('xmldoc unloaded');
   }

}

export default Canvasa;
