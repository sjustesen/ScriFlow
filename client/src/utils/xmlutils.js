/* 
    (c) 2022 Simon Justesen
    This class contains functions to manipulate generic XML documents BEFORE it is turned 
    into SLADocuments (see utils/sladocuments.js).
    
    MIT Licensed
    */

export default class XMLUtils {

    parseXML(val) {
        let xmlDoc = null;

        if (window.DOMParser) {
            try {
                let parser = new DOMParser();
                xmlDoc = parser.parseFromString(val, "text/xml");
            } catch {
                console.dir('XMLParser couldn\'t parse xml');
            }
        }
        return xmlDoc;
    }
}