import React, { Component } from 'react';
import Header from '../shared/Header.js'
import SidePanel from '../shared/SidePanel.js'
import '../styles/Projects.css';

/**
* Displays a list of ItemCards representing the projects available for a class. 
* Students can click on an ItemCard to submit or view their submission,
* while professors click on ItemCards to view project analytics.
* Professors can also update or insert new projects from this page. 
*/
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
        <Header />
        <SidePanel />
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
