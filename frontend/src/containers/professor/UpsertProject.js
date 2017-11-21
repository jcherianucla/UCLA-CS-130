import React, { Component } from 'react';
import Header from '../../shared/Header.js'
import SidePanel from '../../shared/SidePanel.js'
import '../../styles/professor/UpsertProject.css';

/**
* Form where professors can add a new project or update an existing project.
*/
class ProfessorUpsertProject extends Component {

  back() {
    this.props.history.goBack();
  }

  render() {
    return (
      <div>
        <SidePanel />
        <div className="page">
          <Header title="Welcome!" path="Create/Edit Project" />
            {/* TODO: Change class creation form to match designs */}
            <form id="class-form">
              <div className="class-form-group">
                <input className="class-form-input" type="text" required="required" />
                <span className="class-form-bar"></span>
                <label className="class-form-label">Project Name</label>
              </div>
              <div className="class-form-group">
                <input className="class-form-input secret" type="text" required="required"/>
                <span className="class-form-bar"></span>
                <label className="class-form-label">Project Description</label>
              </div>
            </form>
            <div class="upload-btn-wrapper">
              <button class="btn">Upload .sh</button>
              <input type="file" name="myfile" accept=".csv"/>
            </div>
            {/* TODO: Add date picker */}
        </div>
      </div>
    );
  }
}

export default ProfessorUpsertProject;
