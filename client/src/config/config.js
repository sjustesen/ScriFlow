export default class Config {
    
    static getProjectsListUrl = () => 'http://localhost:8080/projects/list';
    
    loadProjectById(id) { return `http://localhost:8080/project/${id}`; }
}