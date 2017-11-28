import React, { Component } from 'react';
import '../styles/shared/Header.css';

/**
 * The header that shows up on every page once a user has logged in.
 */
class Header extends Component {

  Home() {
    this.props.history.push('/classes');
  }

  render() {
    return (
    	<div>
	      <div className="header">
          <ul>
  	      	<li>
              <a href="/classes" onClick={() => this.Home()}>
                <div className="logo"/>
              </a>
            </li>
            <li>
              <div className="faq">
                <h1>FAQ</h1>
              </div>
            </li>
            <li>
              <a href="/classes" onClick={() => this.Home()}>
                <div className="home">
                  <h1>Home</h1>
                </div>
              </a>
            </li>
          </ul>
	      </div>

	      <div className="welcome bold">
	      	{this.props.title}
	      </div>

	      <div className="path">
	        {this.props.path}
	      </div>
      </div>
    );
  }
}

export default Header;
