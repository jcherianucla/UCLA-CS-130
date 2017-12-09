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
    if (state == null) {
      let role = localStorage.getItem('role');
      if (role == null) {
        this.props.history.push('/');
      }
      this.setState({role: role});
    } else {
      this.setState({role: state.role});
    }
  }

  classes() {
    this.props.history.push('/classes');
  }

  login() {
    localStorage.setItem('role', this.state.role);
    this.classes();
  }

  changeLogin(role) {
    this.setState({role: role});
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
          { this.state.role === "student" ?
            <div>
              <Header title="Welcome to GradePortal!" path={["Login as Student"]} />
              <GoogleLogin
                clientId="770443881218-53j89rnpv5539ad9dn69vd4mj51lmr1n.apps.googleusercontent.com"
                buttonText=""
                className="google"
                onSuccess={(response) => this.responseGoogle(response)}
                onFailure={(response) => this.responseGoogle(response)}
              />
              <button id="change-login" className="blue text-center" onClick={() => this.changeLogin("professor")}>Or Login as a Professor</button>
            </div>
            :
            <div>
              <Header title="Welcome to GradePortal!" path={["Login as Professor"]} />
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
              <button id="change-login" className="blue text-center" onClick={() => this.changeLogin("student")}>Or Login as a Student</button>
            </div>
          }
        </div>
      </div>
    );
  }
}

export default Login;
