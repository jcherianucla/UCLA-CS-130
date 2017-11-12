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
          <div id="landing-background-team"></div>
          <div id="landing-content-team" className="text-center">
            <h1 className="landing-header text-center">Team</h1>
            <div id="landing-katie" className="landing-pic">
              <div className="landing-pic-caption text-center">
                <p>Katie Aspinwall</p>
                <p>Frontend Developer</p>
              </div>
            </div>
            <div id="landing-shalini" className="landing-pic">
              <div className="landing-pic-caption text-center">
                <p>Shalini Dangi</p>
                <p>Frontend Developer</p>
              </div>
            </div>
            <div id="landing-prit" className="landing-pic">
              <div className="landing-pic-caption text-center">
                <p>Prit Joshi</p>
                <p>Backend Developer</p>
              </div>
            </div>
            <br />
            <div id="landing-connor" className="landing-pic">
              <div className="landing-pic-caption text-center">
                <p>Connor Kenny</p>
                <p>Frontend Developer</p>
              </div>
            </div>
            <div id="landing-jahan" className="landing-pic">
              <div className="landing-pic-caption text-center">
                <p>Jahan Kuruvilla Cherian</p>
                <p>Backend Developer</p>
              </div>
            </div>
            <div id="landing-omar" className="landing-pic">
              <div className="landing-pic-caption text-center">
                <p>Omar Ozgur</p>
                <p>Backend Developer</p>
              </div>
            </div>
          </div>
        </div>
        <div id="landing-section-contact">
          <h1 className="landing-header text-center">Contact</h1>
          <p id="text-center">Have feedback? Want to learn more? Email us at gradeportal@gmail.com</p>
        </div>
        <div id="landing-section-footer">
        </div>
      </div>
    );
  }
}

export default Landing;
