export default class Config {
    constructor() {
        this.listProjectsUrl = 'http://localhost:8080/projects/list';
        this.projectDetails = 'http://localhost:8080/project/';
    }

    getProjectsListUrl = () => this.listProjectsUrl;
    loadProjectById(id) {
        return this.projectDetails + id;
    }
}