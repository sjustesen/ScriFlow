import XMLUtils from '../utils/xmlutils';

export default class DocumentService {
   
    async loadFromUrl(url) {
      var res =  await fetch(url)
            .then(response => response.text())
            .then(fetched_data => {
                let xmlutils = new XMLUtils()
                let parsed_xml = xmlutils.parseXML(fetched_data);
                return parsed_xml;
            });
        return res;
    }
}