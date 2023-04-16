export default class Config {
    constructor() {
        this.listProjectsUrl = 'http://localhost:8080/projects/list';
        this.projectDetails = 'http://localhost:8080/project/';
    }

    getProjectsListUrl() = () => this.listProjectsUrl;
    getProjectLoadUrl(id) {
        return this.projectDetails + id;
    }
}