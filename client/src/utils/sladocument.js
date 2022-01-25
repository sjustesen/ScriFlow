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

    // stub
    getProperty(propname, attribute) {
        let attributes = new Map();
        var documentNode = this.xmldata.querySelector(propname);

    }

    // stub
    setProperty(propname, value) {

    }

    /* List the overall properties of the document */

    getDocumentProperties() {
        let attributes = new Map();
        var documentNode = this.xmldata.querySelector('DOCUMENT');

        // loop though DOCUMENT xmlnode and put attributes in a map
        // I'm lowercasing every nodeName as the Scribus SLA document uses inconsistent 
        // word casing

        for (let i = 0; i < documentNode.attributes.length; i++) {
            let attr = documentNode.attributes.item(i);
            attributes.set(attr.nodeName.toLowerCase(), attr.nodeValue);
        };
        return attributes;
    }

    getColors() {
        let colors = new Map();
        var colorNodes = this.xmldata.querySelectorAll('COLOR');

        for (var color of colorNodes) {
            let colorAttribs = [];
            for (let i = 0; i < color.attributes.length; i++) {
                let attr = color.attributes.item(i);
                colorAttribs.push({ attribute: attr.nodeName.toLowerCase(), value: attr.nodeValue});
            };
            colors.set(color.attributes[0].nodeValue.toLowerCase(), colorAttribs)
        }
        return colors;
    }

    getColorInfo(colorname) {
        //this.getColors()
    }

    /* get the state of the layers */

    getLayers() {
        const layerdata = this.xmldata.querySelectorAll('LAYERS');
       
        let layers = new Map();
        let node_count = 0;

        layerdata.forEach(layer_node => {
            let nodeAttributes = [];
            for (let i = 0; i < layer_node.attributes.length; i++) {
                let attr = layer_node.attributes.item(i);
                nodeAttributes.push({ attribute: attr.nodeName.toLowerCase(), value: attr.nodeValue});
            };
            layers.set('layer_' + node_count, nodeAttributes)
            node_count++;
        });
        return layers;
    }

    getLayerByIndex(id) {
        let layer = this.xmldata.querySelectorAll('LAYERS')[id];
        return layer;
    }
}