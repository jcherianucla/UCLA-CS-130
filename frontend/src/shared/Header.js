import React, { Component } from 'react';
import '../styles/shared/Header.css';

class Header extends Component {
  render() {
    return (
    	<div>
	      <div className="header">
          <ul>
  	      	<li>
              <div className="logo" />
            </li>
            <li>
              <div className="faq">
                <h1>FAQ</h1>
              </div>
            </li>
            <li>
              <div className="home">
                <h1>Home</h1>
              </div>
            </li>
          </ul>
	      </div>

	      <div className="welcome bold">
	      	Welcome!
	      </div>

	      <div className="path">
	        Path
	      </div>
      </div>
    );
  }
}

export default Header;
