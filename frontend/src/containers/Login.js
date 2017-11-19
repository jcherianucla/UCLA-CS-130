import React, { Component } from 'react';
import GoogleLogin from 'react-google-login';
import Header from '../shared/Header.js'
import SidePanel from '../shared/SidePanel.js'
import '../styles/shared/Login.css';
import '../styles/shared/Page.css';

/**
* Login page.
*/

class Login extends Component {

  componentWillMount() {
    let state = this.props.history.location.state;
    this.setState({
      role: state.role
    });
  }

  classes() {
    this.props.history.push('/classes');
  }

  login() {
    localStorage.setItem('role', this.state.role);
    this.classes();
  }

  responseGoogle(response) {
    // let profileObj = response.profileObj;
    this.login();
  }

  render() {
    return (
      <div>
        <SidePanel />
        <div className="page">
          <Header title="Welcome to GradePortal!" path="Login" />
          { this.state.role === "student" ?
            <GoogleLogin
              clientId="770443881218-53j89rnpv5539ad9dn69vd4mj51lmr1n.apps.googleusercontent.com"
              buttonText=""
              className="google"
              onSuccess={(response) => this.responseGoogle(response)}
              onFailure={(response) => this.responseGoogle(response)}
            />
            :
            <form id="login-form" onSubmit={() => this.login()}>
              <div className="login-form-group">
                <input className="login-form-input" type="text" required="required" />
                <span className="login-form-bar"></span>
                <label className="login-form-label">Email</label>
              </div>
              <div className="login-form-group">
                <input className="login-form-input secret" type="text" required="required"/>
                <span className="login-form-bar"></span>
                <label className="login-form-label">Password</label>
              </div>
              <input className="login-form-btn" type="submit" />
            </form>
          }
        </div>
      </div>
    );
  }
}

export default Login;
