import React from 'react';
import './Header.css';

class Header extends React.Component {

	render() {
		return ( 
		<header id="top_tools">
				<div className="menuitem"><img src="images/menuitems/new_document.svg" /></div>	
				<div className="menuitem openfile"><img src="images/menuitems/open_file.svg" /></div>
				<div className="menuitem"><img src="images/menuitems/new_document.svg" /></div>
				<div className="menuitem"><img src="images/menuitems/new_document.svg" /></div>	
				<div className="menuitem"></div>
				<div className="menuitem"></div>
				<div className="menuitem"></div>	
				<div className="menuitem"></div>
				<div className="menuitem"></div>

		</header>
		);
	}
}

export default Header;
