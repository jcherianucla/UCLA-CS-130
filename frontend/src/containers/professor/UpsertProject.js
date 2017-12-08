import React, { Component } from 'react';
import Header from '../../shared/Header.js'
import SidePanel from '../../shared/SidePanel.js'
import '../../styles/professor/Form.css';
import '../../styles/professor/UpsertProject.css';

/**
* Form where professors can add a new project or update an existing project.
*/
class ProfessorUpsertProject extends Component {

  back() {
    this.props.history.goBack();
  }

  projects() {
    this.props.history.push('/projects');
  }

  render() {
    return (
      <div>
        <SidePanel />
        <div className="page">
          <Header title="Welcome!" path={["Login", "Classes", "Projects", "Create/Edit Project"]} />
            <div class="create-form">
              <form onSubmit={() => this.projects()}>
                <label class="upsert-label"><b>Project Name</b></label>
                <input type="text" placeholder="Enter project name"/>
                
                <label class="upsert-label"><b>Project Description</b></label>
                <textarea placeholder="Enter short description of your project" rows="3" cols="40"/>

                <label class="upsert-label"><b>Upload Grading Script</b></label>
                <div class="upload-btn-wrapper">
                  <input class="btn" type="file" name="myfile" accept=".sh"/>
                  <button class="btn">Upload .sh</button>
                </div>

                <div class="deadline-wrapper">
                  <label class="upsert-label"><b>Project Deadline</b></label>
                  <input type="text" placeholder="MM" maxlength="2"/> &nbsp; / &nbsp;
                  <input type="text" placeholder="DD" maxlength="2"/> &nbsp; / &nbsp;
                  <input type="text" placeholder="YY" maxlength="2"/> &nbsp; &nbsp; &nbsp;
                  <input type="text" placeholder="00" maxlength="2"/> &nbsp; : &nbsp;
                  <input type="text" placeholder="00" maxlength="2"/>
                </div>
                
                <div>
                  <input className="submit-btn" type="submit" />
                </div>
              </form>
            </div>
        </div>
      </div>
    );
  }
}

export default ProfessorUpsertProject;
