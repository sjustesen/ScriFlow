/* The purpose of this class is to act like a proxy between "native" ScriFlow calls and JSFabric */
/* Scribus works in points */

export default class SF2DDrawing {

    constructor() {

    }

    drawLine() {

    }

    drawPath() {
        let path = ''

    }

    drawPolyLine(array_of_points) {
        var poly = new fabric.Polyline([
            { x: 10, y: 10 },
            { x: 50, y: 30 },
            { x: 40, y: 70 },
            { x: 60, y: 50 },
            { x: 100, y: 150 },
            { x: 40, y: 100 }
          ], {
          stroke: 'red',
          left: 100,
          top: 100
        });
        return poly
    }

    drawBox(x, y, w, h) {

    }
}