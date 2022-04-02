export default class PageObject {
     
    constructor(type, ...options) {
        this._ptype = type; // object type as an int
        this._width = width;
        this._height = height;

        // object type
        // https://wiki.scribus.net/canvas/File_Format_Specification_for_Scribus_1.4#Scribus_Template_Files_.28XML.29
        // Seems to be the same for Scribus 1.5

        this._ptypes = {
            2: 'image',
            4: 'text',
            5: 'line',
            6: 'polygon',
            7: 'polyline',
            8: 'textpath'
        }

        this._options = options;
        this._color = options['color'] != null ? options['color'] : null;
    }

    getType() {
        if(this._ptypes.contains(this._ptype))
            return this._ptypes[this._ptype];
    }

    getRotation() {

    }

    getColor() {
        return this._color;
    }

    // To be continued

}