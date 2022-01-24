/* 
    (c) 2022 Simon Justesen
    This class contains functions to work with SLADocuments AFTER they've been parsed into the JS DOM hierachy.
    Right now I'm using built-in queryselectors to inspect nodes, perhaps there are better and/or faster ways of doing this,
    like on-the-fly parsing. 
    
    MIT Licensed
    */

import Layers from "../components/Layers";

export default class SLADocument {

    constructor(xmldata) {
        this.xmldata = xmldata;
    }

    /* List the overall properties of the document */

    getDocumentProperties() {
        let attributes = new Map();
        var documentNode = this.xmldata.querySelector('DOCUMENT');
            
        // loop though DOCUMENT xmlnode and put attributes in a map
            // I'm lowercasing every nodeName as the Scribus SLA document uses inconsistent 
            // word casing

            for (let i = 0; i < documentNode.attributes.length; i++) {
                let attr = documentNode.attributes.item(i)
                attributes.set(attr.nodeName.toLowerCase(), attr.nodeValue);
            };
        return attributes;
    }

    /* get the state of the layers */

    getLayers() {
        const layerdata = this.xmldata.querySelectorAll('LAYERS');
        let parser = new DOMParser();
        layerdata.forEach(layer_node => {
            layer_node.attributes.forEach(attribute => {

            })
        });
        return layerdata;
    }

    getLayer(id) {
        let layer= this.xmldata.querySelectorAll('LAYERS')[id];
        return layer;
    }
}