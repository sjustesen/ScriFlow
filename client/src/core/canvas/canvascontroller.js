import { fabric } from '../../lib/fabric.min.js';
import Guides from "@scena/guides";

import PageObject from '../models/pageobject.js';

import SF2DDrawing from '../drawing/draw'
import SFText from '../drawing/text';

export default class SFCanvasController {

   constructor() {
      this.initCanvas();
      this.initGuides();
      this.registerCanvasEvents();
   }

   initCanvas() {
      let grandparent = document.querySelector('.canvascontainer');
      let parent = document.querySelector('.canvas');

      let width = grandparent.clientWidth - 100;

      parent.setAttribute('style', 'width: ' + width + 'px');

      this.canvas = new fabric.Canvas('c');
      this.canvas.setHeight(parent.clientHeight);
      this.canvas.setWidth(parent.clientWidth);

      window.addEventListener('resize', (e) => {
         let width = grandparent.clientWidth - 100;
         parent.setAttribute('style', 'width: ' + width + 'px');
         this.canvas.setHeight(parent.clientHeight);
         this.canvas.setWidth(parent.clientWidth);

      })
   }


   registerCanvasEvents() {
      window.addEventListener('XmlDocumentLoaded', () => {

      })
      window.addEventListener('LayersPanelChanged', () => {
         console.log('Layers Panel change signaled..')
      })

      window.addEventListener('CanvasRedrawn', (event) => {
      });

      window.addEventListener('CanvasUpdated', (event) => {
         this.drawObject(event.detail);

      });

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

   drawObject(objs) {
      objs.forEach(element => {
         let el = new PageObject(parseInt(element.attributes.ptype), element.attributes);
         console.dir(el)

         switch (el.getType()) {
            case "text":
               this.canvas.add(new fabric.Text('Hest', { left: el.getX(), top: el.getY() }));
               break;
            case 5:
               break;
            case "polygon":
               let path = new fabric.Path(element.attributes.path);
               path.set({ fill: el.getColor(), stroke: el.getStrokeColor(), opacity: 0.5 });
               this.canvas.add(path);
               break;
            default:
               break;
         }
      });
      // add object to the referenced canvas
      //this.canvas.add(pageobject);
   }


}