import { fabric } from '../../lib/fabric.min.js';
import PageObject from '../models/pageobject.js';

import SF2DDrawing from './draw'
import SFText from './text';

export default class SFCanvasLib {

    constructor(renderer) {
        this.renderer = renderer;
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
       this.renderer.add(pageobject);
    }


}