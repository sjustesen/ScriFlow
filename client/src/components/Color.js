import React from 'react';
import ColorUtil from '../utils/colorutils';

class Color extends React.Component {
    // eslint-disable-next-line
    constructor(props) {
        super(props);
        this.state = {};
      
    }

    showColor(colorspace, attribs) {
        let color = {
            r: 0,
            g: 0,
            b: 0
        }
        
        if (colorspace.toUpperCase() == 'CMYK') {
            color = ColorUtil.convertCMYKtoRGB(attribs.c, attribs.m, attribs.y, attribs.k)
        } else {  
            color.r = attribs.r;
            color.g = attribs.g; 
            color.b = attribs.b;       
        }

        return `${color.r},${color.g},${color.b}`;
        
    }

    render() {
        let item = this.props.item;
        let colorspace = (item.attributes["space"] == 'RGB') ? 'RGB' : 'CMYK';
        var swatch_color = this.showColor(colorspace, item.attributes);
        let name = this.props.item.name;

        return (
        <div className='togglebutton'>
            <div className='colorbox' style={{backgroundColor: `rgb(${swatch_color.toString()})`}}></div>
            <div>{name}</div>
        </div> 
        )
    }

}
export default Color;
