import { fabric } from '../../lib/fabric.min.js';
import Guides from "@scena/guides";

import PageObject from '../models/pageobject.js';

import SF2DDrawing from './draw'
import SFText from './text';

export default class SFCanvasLib {

    constructor() {
        this.initCanvas();
        this.initGuides();
        this.registerCanvasEvents();
    }

    initCanvas() {
        let grandparent = document.querySelector('.canvascontainer');
        let parent = document.querySelector('.canvas');
  
        let width = grandparent.clientWidth - 60;
  
        parent.setAttribute('style', 'width: ' + width + 'px');
  
        this.canvas = new fabric.Canvas('c');
        this.canvas.setHeight(parent.clientHeight);
        this.canvas.setWidth(parent.clientWidth);
  
        window.addEventListener('resize', (e) => {
           let width = grandparent.clientWidth - 60;
           parent.setAttribute('style', 'width: ' + width + 'px');
           this.canvas.setHeight(parent.clientHeight);
           this.canvas.setWidth(parent.clientWidth);
  
        })
     }
     

    registerCanvasEvents() {
        window.addEventListener('XmlDocumentLoaded', () => {

        })
        window.addEventListener('LayersPanelChanged', () => {
           console.log('Update Layer Panel signaled..')
        })
  
        window.addEventListener('CanvasRedrawn', (event) => {
        });
  
        window.addEventListener('CanvasUpdated', (event) => {
           for (let canvas_object of event.detail) {
              this.drawObject(canvas_object);
           }
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

    drawObject(obj) {
        var pageobject;

        switch (obj.attributes.ptype) {
            case "4":
                pageobject = new fabric.Text('Hest', { left: 0, top: 0 });
                break;
            case 5:
                break;
            case 6:
                break;
            default:
                break;
        }
        console.dir(pageobject)
        // add object to the referenced canvas
       this.canvas.add(pageobject);
    }


}