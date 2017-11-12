import React, { Component } from 'react';
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
        </div>
      </div>
    );
  }
}

export default StudentLogin;
