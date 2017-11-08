import React, { Component } from 'react';
import Header from '../../shared/Header.js'
import SidePanel from '../../shared/SidePanel.js'
import '../../styles/student/Submission.css';

class StudentSubmission extends Component {

  back() {
    this.props.history.goBack();
  }

  render() {
    return (
      <div>
        <Header />
        <SidePanel />
        <h1>Student Submission</h1>
        <button onClick={() => this.back()}>
          Back
        </button>
      </div>
    );
  }
}

export default StudentSubmission;
