export default class Layer {

    constructor(name, attributes) {
        this.name = name;
        this.attributes = attributes;
    }

    getName() {
        return this.name;
    }

    getAttributes() {
        return this.attributes;
    }
}