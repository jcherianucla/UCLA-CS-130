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

  Logout() {
    localStorage.setItem('token', "");
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
              <a href="/" onClick={() => this.Logout()}>
                <div className="logout">
                  <h1>Logout</h1>
                </div>
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
          {this.props.path.map(function(item, key, arr){
            var path;
            let name;
            if (item.constructor === Array) {
              name = item[item.length - 1];
              path = routes[item[0]]
              for (var i = 1; i < item.length; i++) {
                path = path.replace(/:[\w]*/, item[i]);
              }
            } else {
              name = item;
              path = routes[item];
            }
            if (key === 0) {
              if (key === arr.length - 1) {
                return(<p className="path">{name}</p>);
              }
              else {
                return(<a key={key} className="path" href={path}>{name}</a>);
              }
            }
            else {
              if (key === arr.length - 1) {
                return(<p className="path">> {name}</p>);
              }
              else {
                return(<a key={key} className="path" href={path}>> {name}</a>);
              }
            }
          })}
	      </div>
      </div>
    );
  }
}

export default Header;
