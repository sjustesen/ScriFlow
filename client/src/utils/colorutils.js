/* 
    (c) 2022 Simon Justesen
    This class contains functions to complement SLADocuments
    
    MIT Licensed
    */

export default class ColorUtil {

    static convertCMYKtoRGB(c,m,y,k) {
        let normalized = false;

        c = (c / 100);
        m = (m / 100);
        y = (y / 100);
        k = (k / 100);
        
        c = c * (1 - k) + k;
        m = m * (1 - k) + k;
        y = y * (1 - k) + k;
        
        var r = 1 - c;
        var g = 1 - m;
        var b = 1 - y;
        
        if(!normalized){
            r = Math.round(255 * r);
            g = Math.round(255 * g);
            b = Math.round(255 * b);
        }
        
        return {
            r: r,
            g: g,
            b: b
        }
    }

    static convertRGBtoCMYK(r, g, b) {
        // http://javascripter.net/faq/rgb2cmyk.htm

        var computedC = 0;
        var computedM = 0;
        var computedY = 0;
        var computedK = 0;

        // BLACK
        if (r == 0 && g == 0 && b == 0) {
            computedK = 1;
            return [0, 0, 0, 1];
        }

        computedC = 1 - (r / 255);
        computedM = 1 - (g / 255);
        computedY = 1 - (b / 255);

        var minCMY = Math.min(computedC,
            Math.min(computedM, computedY));
        computedC = (computedC - minCMY) / (1 - minCMY);
        computedM = (computedM - minCMY) / (1 - minCMY);
        computedY = (computedY - minCMY) / (1 - minCMY);
        computedK = minCMY;

        return {c: computedC, m: computedM, y: computedY, k: computedK };
    }

}