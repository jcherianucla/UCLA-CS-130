import React, { Component } from 'react';
import logo from '../images/logo.svg';
import '../styles/Login.css';

class Login extends Component {
  render() {
    return (
      <div className="Login">
        <header className="Login-header">
          <img src={logo} className="Login-logo" alt="logo" />
          <h1 className="Login-title">Welcome to React</h1>
        </header>
        <p className="Login-intro">
          To get started, edit <code>src/containers/Login.js</code> and save to reload.
        </p>
      </div>
    );
  }
}

export default Login;
