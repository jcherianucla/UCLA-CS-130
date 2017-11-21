import React, { Component } from 'react';
import Header from '../../shared/Header.js'
import SidePanel from '../../shared/SidePanel.js'
import '../../styles/student/Submission.css';

/**
* Form where students can upload and submit a project. 
*/
class StudentSubmission extends Component {

  back() {
    this.props.history.goBack();
  }

  render() {
    return (
      <div>
        <SidePanel />
        <div className="page">
          <Header title="Welcome!" path="Submission" />
        </div>
      </div>
    );
  }
}

export default StudentSubmission;
