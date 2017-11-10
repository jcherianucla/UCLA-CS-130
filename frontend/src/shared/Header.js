import React, { Component } from 'react';
import '../styles/shared/Header.css';

class Header extends Component {
  render() {
    return (
    	<div>
	      <div className="header">
	      	<div className="logo" /> 
	        <p>This is an example header.</p>
	      </div>

	      <div className="welcome">
	      	Welcome Joe Bruin
	      </div>

	      <div className="path">
	        Path
	      </div>

      </div>
    );
  }
}

export default Header;
