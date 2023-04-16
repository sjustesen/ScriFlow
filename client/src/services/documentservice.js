import XMLUtils from '../utils/xmlutils';
import Config from '../config/config';

export default class DocumentService {

    async loadProjectList() {
        var res =  await fetch(new Config().getProjectsListUrl())
              .then(response => response.json())
              .then(fetched_data => {
                  return fetched_data;
              });
          return res;
      }


    async loadFromUrl(id) {
      var res =  await fetch(new Config().loadProjectByUrl(id))
            .then(response => response.text())
            .then(fetched_data => {
                let xmlutils = new XMLUtils()
                let parsed_xml = xmlutils.parseXML(fetched_data);
                return parsed_xml;
            });
        return res;
    }
}