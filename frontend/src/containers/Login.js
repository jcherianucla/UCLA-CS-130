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
    if (localStorage.getItem('role') !== "" && localStorage.getItem('token') !== "") {
      this.props.history.push('/classes');
    }
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

  loginAsStudent(firstName, lastName, email) {
    let self = this
    fetch('http://grade-portal-api.herokuapp.com/api/v1.0/bol', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        first_name: firstName,
        last_name: lastName,
        email: email,
      })
    })
    .then((response) => response.json())
    .then(function (responseJSON) {
      console.log(responseJSON);
      if (responseJSON.message !== 'Success') {
        alert(responseJSON.message);
      } else {
        localStorage.setItem('token', responseJSON.token);
        localStorage.setItem('role', 'student');
        self.props.history.push('/classes');
      }
    });
  }

  loginAsProfessor() {
    localStorage.setItem('role', 'professor');
    this.props.history.push('/classes')
  }

  changeLogin(role) {
    this.setState({role: role});
  }

  responseGoogle(response) {
    this.loginAsStudent(response.profileObj.givenName, response.profileObj.familyName, response.profileObj.email);
  }

  render() {
    return (
      <div>
        <SidePanel />
        <div className="page">
          <Header title="Welcome to GradePortal!" path={[]} />
          <br /><br />
          { this.state.role === "student" ?
            <div>
              <h1 className="blue text-center">Login as a Student</h1>
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
              <h1 className="blue text-center">Login as a Professor</h1>
              <form id="login-form" onSubmit={() => this.loginAsProfessor()}>
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
