import React, { Component } from 'react';
import '../../styles/student/Submission.css';

class StudentSubmission extends Component {

  back() {
    this.props.history.goBack();
  }

  render() {
    return (
      <div>
        <h1>Student Submission</h1>
        <button onClick={() => this.back()}>
          Back
        </button>
      </div>
    );
  }
}

export default StudentSubmission;
