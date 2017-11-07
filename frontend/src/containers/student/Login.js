import React, { Component } from 'react';
import '../../styles/student/Login.css';

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
        <h1>Student Login</h1>
        <button onClick={() => this.back()}>
          Back
        </button>
        <button onClick={() => this.classes()}>
          Classes
        </button>
      </div>
    );
  }
}

export default StudentLogin;
