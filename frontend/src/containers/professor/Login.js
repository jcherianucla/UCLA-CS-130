import React, { Component } from 'react';
import GoogleLogin from 'react-google-login';
import Header from '../../shared/Header.js'
import SidePanel from '../../shared/SidePanel.js'
import '../../styles/shared/Login.css';
import '../../styles/shared/Page.css';

class ProfessorLogin extends Component {

  back() {
    this.props.history.goBack();
  }

  classes() {
    this.props.history.replace('/classes', this.state);
  }

  responseGoogle(response) {
    this.setState({firstName: response.profileObj.givenName});
    this.setState({lastName: response.profileObj.familyName});
    this.setState({email: response.profileObj.email});
    this.setState({type: 'professor'});
    this.classes();
  }

  render() {
    return (
      <div>
        <SidePanel />
        <div className="page">
          <Header title="Welcome to GradePortal!" path="Professor Login" />
          <GoogleLogin
            clientId="770443881218-53j89rnpv5539ad9dn69vd4mj51lmr1n.apps.googleusercontent.com"
            buttonText=""
            className="google"
            onSuccess={(response) => this.responseGoogle(response)}
            onFailure={(response) => this.responseGoogle(response)}
          />
        </div>
      </div>
    );
  }
}

export default ProfessorLogin;