/* 
    (c) 2022 Simon Justesen
    This class contains functions to work with SLADocuments AFTER they've been parsed into the JS DOM hierachy.
    Right now I'm using built-in queryselectors to inspect nodes, perhaps there are better and/or faster ways of doing this,
    like on-the-fly parsing. 

    Note: All XML Element attributes are converted to lowercase to make up for inconsistencies in the SLA Document format.

    TODO: Import children of elements
    
    MIT Licensed
    */

import Layer from "../components/Layer";
import SLAUtils from "./slautils";

export default class SLADocument {

    constructor(xmldata) {
        this.xmldata = xmldata;
    }

    // get the elements of the SLADoc by property name

    getElement(propname) {
        let result = [];
        let xmlattributes = new Map();
        let xmlelements = this.xmldata.querySelectorAll(propname);

        if (xmlelements.length === 0) {
            throw Error("Sorry, there are no tag(s) named " +propname);
        }
        
        let elementId = 0;
        for (var element of xmlelements) {
            let elementName = 'element_' + elementId;
        
            for (let i = 0; i < element.attributes.length; i++) {
                let attr = element.attributes.item(i);
                let nodename = attr.nodeName.toLowerCase();
                xmlattributes.set(nodename, attr.nodeValue);

                if (nodename == 'name') {
                    elementName = attr.nodeValue;
                }
            };
           
            result.push({ 
                id: elementName.trim().toLowerCase().replace(' ', '_'),
                name: elementName, 
                type: propname.toLowerCase(), 
                attributes: xmlattributes
            })

            elementId++;
        }
        return result;
    }

    // stub
    setProperty(propname, value) {

    }

    /* List the overall properties of the document */

    getDocumentProperties() {
        let styles = this.getElement('STYLE');
        return styles;
    }

    getPageObjects() {
        let pageobjects = this.getElement('PAGEOBJECT');
        // note: This has subnodes
        return pageobjects;
    }

    getMasterpage(uid) {
        let masterpage = this.getElement('MASTERPAGE');
        return masterpage;
    }

    getMasterpages() {
        let masterpages = this.getElement('MASTERPAGE');
        return masterpages;
    }

    getStyles() {
        let styles = this.getElement('STYLE');
        return styles;
    }

    getStyle() {
        let styles = this.getElement('STYLE');
        return styles;
    }

    getColors() {
        let colors = this.getElement("COLOR");
        return colors;
    }

    getColorInfo(colorname) {
        //this.getColors()
    }

    getLPI() {
        return this.getElement("LPI");
    }

    /* get the state of the layers */

    getLayers() {
        const layerdata = this.xmldata.querySelectorAll('LAYERS');
       
        let layers = [];
        let node_count = 0;

        layerdata.forEach(layer_node => {
            let nodeAttributes = [];
            for (let i = 0; i < layer_node.attributes.length; i++) {
                let attr = layer_node.attributes.item(i);
                nodeAttributes.push({ attribute: attr.nodeName.toLowerCase(), value: attr.nodeValue});
            };
            let layer = new Layer('layer_' + node_count, nodeAttributes);
            layers.push({ key: 'layer_'+node_count, properties: layer })
            node_count++;
        });
        return layers;
    }

    getLayerByIndex(id) {
        let layer = this.xmldata.querySelectorAll('LAYERS')[id];
        return layer;
    }
}