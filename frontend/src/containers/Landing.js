import React, { Component } from 'react';
import '../styles/Landing.css';

class Landing extends Component {

  professorLogin() {
    this.props.history.push('/professor/login');
  }

  studentLogin() {
    this.props.history.push('/student/login');
  }

  render() {
    return (
      <div id="landing">
        <div id="landing-section-top">
          <div id="landing-background-top"></div>
          <div id="landing-content-top">
            <div>
              <div className="landing-logo" />
            </div>
            <h1 className="landing-title bold">GradePortal</h1>
            <h2 className="landing-subtitle">The real-time project submission and feedback platform for UCLA</h2>
            <button id="landing-login-professor" className="landing-login-button" onClick={() => this.professorLogin()}>
              Login as Professor
            </button>
            <button id="landing-login-student" className="landing-login-button" onClick={() => this.studentLogin()}>
              Login as Student
            </button>
            <div id="landing-analytics-card"></div>
            <div id="landing-classes-card"></div>
            <div id="landing-feedback-card"></div>
          </div>
        </div>
        <div id="landing-section-about">
          <h1 className="landing-header">About</h1>
          <p id="landing-about-paragraph">Project submissions can be scary tasks, especially with the fear of the unknown when it comes to test cases. GradePortal aims to make this process easier on you. By working with professors we want to bring you pre-project grades, and post deadline project feedback, otherwise only known to TA’s and professors. We want to make the grading process transparent and beneficial to the students.</p>
        </div>
        <div id="landing-section-team">
        </div>
        <div id="landing-section-contact">
        </div>
        <div id="landing-section-footer">
        </div>
      </div>
    );
  }
}

export default Landing;
