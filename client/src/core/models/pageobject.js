export default class PageObject {
     
    constructor(type, ...options) {
        this._ptype = type; // object type as an int
        this._rotation = options['rot'];
        this._width = options['width'];
        this._height = options['height'];
        this._color = options['pcolor'];
        this._stroke = options['pcolor2'];

        // object type
        // https://wiki.scribus.net/canvas/File_Format_Specification_for_Scribus_1.4#Scribus_Template_Files_.28XML.29
        // Seems to be the same for Scribus 1.5

        this._objtypes = {
            2: 'image',
            4: 'text',
            5: 'line',
            6: 'polygon',
            7: 'polyline',
            8: 'textpath'
        }

        this._options = options;
    }

    getX() {
        return this.ypos;
    }

    getY() {
        return this._xpos;
    }

    getType() {
        if(this._objtypes.hasOwnProperty(this._ptype))
            return this._objtypes[this._ptype];
    }

    getRotation() {
        return this._rotation;
    }

    getColor() {
        return this._color;
    }

    getStrokeColor() {
        return this._stroke;
    }

    // To be continued

}