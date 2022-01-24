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
        let attributes = Set();
        var documentNode = this.xmldata.querySelector('DOCUMENT');

            // loop though DOCUMENT node and put everything in a set
            for (let attribute of documentNode.attributes) {
                attributes.add(attribute);
            };
        return documentNode;
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
        let layers = this.xmldata
        return Layers;
    }
}