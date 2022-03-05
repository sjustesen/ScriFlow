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
            color = ColorUtil.convertCMYKtoRGB(attribs.get('c'), attribs.get('m'), attribs.get('y'), attribs.get('k'))
        } else { // Color is RGB or unknown
            color.r = attribs.get('r');
            color.g = attribs.get('g'); 
            color.b = attribs.get('b');
        }
        return [
            color.r,
            color.g,
            color.b
        ]
        
    }

    render() {
        let attribs = this.props.item.attributes;
        let colorspace = (attribs.get('space') == 'CMYK') ? 'CMYK' : 'RGB';
        var swatch_color = this.showColor(colorspace, attribs);
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
