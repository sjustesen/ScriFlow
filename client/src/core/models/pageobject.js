export default class PageObject {
     
    constructor(type, ...options) {

        this._options = options[0];

        this._ptype = type; // object type as an int
        this._rotation = this._options['rot'];
        this._width = this._options['width'];
        this._height = this._options['height'];
        this._color = this._options['pcolor'];
        this._stroke = this._options['pcolor2'];
        this._xpos = this._options['xpos'];
        this._ypos = this._options['ypos'];

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
    }

    getX() {
        return this._xpos;
    }

    getY() {
        return this._ypos;
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