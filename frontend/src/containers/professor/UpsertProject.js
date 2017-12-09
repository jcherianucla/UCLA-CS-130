import React, { Component } from 'react';
import Header from '../../shared/Header.js'
import SidePanel from '../../shared/SidePanel.js'
import '../../styles/professor/Form.css';
import '../../styles/professor/UpsertProject.css';

/**
* Form where professors can add a new project or update an existing project.
*/
class ProfessorUpsertProject extends Component {

  projects(class_id) {
    this.props.history.push('/classes/' + this.props.match.params.class_id);
  }

  getFile() {
    var x = document.getElementById("upload").value;
    if (x === "") {
      document.getElementById("filename").innerHTML = "";
    } else {
      document.getElementById("filename").innerHTML = "*" + x.replace(/^.*\\/, "");
    }
  }

  render() {
    return (
      <div>
        <SidePanel />
        <div className="page">
          { window.location.href.substr(window.location.href.lastIndexOf('/') + 1) === "create" ?
            <Header title="Welcome!" path={["Login", "Classes", "Projects", "Create Project"]} />
            :
            <Header title="Welcome!" path={["Login", "Classes", "Projects", "Edit Project"]} />
          }
            <div className="create-form">
              <form onSubmit={() => this.projects()}>
                <label className="upsert-label"><b>Project Name</b></label>
                <input type="text" placeholder="Enter project name"/>
                
                <label className="upsert-label"><b>Project Description</b></label>
                <textarea placeholder="Enter short description of your project" rows="3" cols="40"/>

                <label className="upsert-label"><b>Upload Grading Script</b></label>
                <div className="upload-btn-wrapper">
                  <input id="upload" className="btn" type="file" name="myfile" onChange={() => this.getFile()} accept=".sh"/>
                  <button className="btn">Upload .sh</button>
                  <label className="filename" id="filename"></label>
                </div>

                <div className="deadline-wrapper">
                  <label className="upsert-label"><b>Project Deadline</b></label>
                  <input type="text" placeholder="MM" maxLength="2"/> &nbsp; / &nbsp;
                  <input type="text" placeholder="DD" maxLength="2"/> &nbsp; / &nbsp;
                  <input type="text" placeholder="YY" maxLength="2"/> &nbsp; &nbsp; &nbsp;
                  <input type="text" placeholder="00" maxLength="2"/> &nbsp; : &nbsp;
                  <input type="text" placeholder="00" maxLength="2"/>
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
