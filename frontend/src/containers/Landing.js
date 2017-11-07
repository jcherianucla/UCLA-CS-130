import React, { Component } from 'react';
import '../styles/Landing.css';

class Landing extends Component {

  back() {
    this.props.history.goBack();
  }

  professorLogin() {
    this.props.history.push('/professor/login');
  }

  studentLogin() {
    this.props.history.push('/student/login');
  }

  render() {
    return (
      <div>
        <h1>Landing</h1>
        <button onClick={() => this.back()}>
          Back
        </button>
        <button onClick={() => this.professorLogin()}>
          Login as Professor
        </button>
        <button onClick={() => this.studentLogin()}>
          Login as Student
        </button>
      </div>
    );
  }
}

export default Landing;
