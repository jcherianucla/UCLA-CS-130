import React, { Component } from 'react';
import '../styles/Projects.css';

class Projects extends Component {

  back() {
    this.props.history.goBack();
  }

  studentSubmission() {
    this.props.history.push('/student/submission');
  }

  professorUpsertProject() {
    this.props.history.push('/professor/upsert_project');
  }

  professorAnalytics() {
    this.props.history.push('/professor/analytics');
  }

  render() {
    return (
      <div>
        <h1>Projects</h1>
        <button onClick={() => this.back()}>
          Back
        </button>
        <button onClick={() => this.studentSubmission()}>
          Student Submission
        </button>
        <button onClick={() => this.professorUpsertProject()}>
          Professor Create/Edit Project
        </button>
        <button onClick={() => this.professorAnalytics()}>
          Professor Analytics
        </button>
      </div>
    );
  }
}

export default Projects;
