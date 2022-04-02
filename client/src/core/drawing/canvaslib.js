import { fabric } from '../../lib/fabric.min.js';
import PageObject from '../models/pageobject.js';

import SF2DDrawing from './draw'
import SFText from './text';

export default class SFCanvasLib {
    
    constructor(renderer) {
        this.renderer = renderer;
    }

    drawObject(objtype) {
        var pageobject = new PageObject();
        
        console.log('drawing object on the canvas');
        console.dir(objtype.attributes.ptype);

        var rect = new fabric.Rect({
            left: 100,
            top: 100,
            fill: 'red',
            width: 20,
            height: 20
         });
         
         // add object to the referenced canvas
         this.renderer.add(rect);
    }


}