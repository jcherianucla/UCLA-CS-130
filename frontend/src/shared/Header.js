import React, { Component } from 'react';
import '../styles/shared/Header.css';
import routes from '../utils/routes.js';

/**
 * The header that shows up on every page once a user has logged in.
 */
class Header extends Component {

  constructor(props) {
    super(props);
    this.state = {
      'name': ''
    }
    this.getUsername();
  }

  Home() {
    this.props.history.push('/classes');
  }

  Logout() {
    localStorage.setItem('token', "");
  }

  FAQ() {
    this.props.history.push('/faq');
  }

  getUsername() {
    let token = localStorage.getItem('token');
    let self = this;
    fetch('http://grade-portal-api.herokuapp.com/api/v1.0/user', {
      method: 'GET',
      headers: {
        'Authorization': 'Bearer ' + token,
      },
    })
    .then((response) => response.json())
    .then(function (responseJSON) {
      console.log(responseJSON);
      if (responseJSON.user !== null) {
        if (responseJSON.user.Is_professor === true) {
          self.setState({'name': "Professor " + responseJSON.user.Last_name});
        } else {
          self.setState({'name': responseJSON.user.First_name});
        }
      }
    });
  }

  render() {
    let self = this;
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
              <a href="/faq" onClick={() => this.FAQ()}>
                <div className="faq">
                  <h1>FAQ</h1>
                </div>
              </a>
            </li>
          </ul>
	      </div>

	      <div className="welcome bold">
	      	Welcome {this.state.name}
	      </div>

	      <div>
          {this.props.path.map(function(item, key, arr){
            var path;
            let name;
            if (item.constructor === Array) {
              if (key === 1) {
                name = self.props.props.class_name;
              } else if (key === 2) {
                name = self.props.props.project_name;
              } else {
                name = item[item.length - 1];
              }
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
                return(<p key={key} className="path">{name}</p>);
              }
              else {
                return(<a key={key} className="path" href={path}>{name}</a>);
              }
            }
            else {
              if (key === arr.length - 1) {
                return(<p key={key} className="path">> {name}</p>);
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
