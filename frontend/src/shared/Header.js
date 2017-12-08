import React, { Component } from 'react';
import '../styles/shared/Header.css';
import routes from '../utils/routes.js';

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
              <a href="/" onClick={() => this.Home()}>
                <div className="logo"/>
              </a>
            </li>
            <li>
              <div className="faq">
                <h1>FAQ</h1>
              </div>
            </li>
          </ul>
	      </div>

	      <div className="welcome bold">
	      	{this.props.title}
	      </div>

	      <div>
          {this.props.path.map(function(item, key){
            return (
            key === 0 ?
              <a key={key} className="path" href={routes[item]}>{item}</a>
            :
              <a key={key} className="path" href={routes[item]}>> {item}</a>
            );
          })}
	      </div>
      </div>
    );
  }
}

export default Header;
