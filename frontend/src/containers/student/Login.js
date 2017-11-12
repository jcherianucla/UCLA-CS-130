import React, { Component } from 'react';
import GoogleLogin from 'react-google-login';
import Header from '../../shared/Header.js'
import SidePanel from '../../shared/SidePanel.js'
import '../../styles/student/Login.css';
import '../../styles/shared/Page.css';

class StudentLogin extends Component {

  back() {
    this.props.history.goBack();
  }

  classes() {
    this.props.history.push('/classes');
  }

  responseGoogle(response) {
    console.log(response.profileObj);
  }

  render() {
    return (
      <div>
        <SidePanel />
        <div className="page">
          <Header />
          <h1>Student Login</h1>
          <button onClick={() => this.back()}>
            Back
          </button>
          <button onClick={() => this.classes()}>
            Classes
          </button>
          <GoogleLogin
            clientId="770443881218-53j89rnpv5539ad9dn69vd4mj51lmr1n.apps.googleusercontent.com"
            buttonText="Login With Google"
            onSuccess={(response) => this.responseGoogle(response)}
            onFailure={(response) => this.responseGoogle(response)}
          />
        </div>
      </div>
    );
  }
}

export default StudentLogin;
