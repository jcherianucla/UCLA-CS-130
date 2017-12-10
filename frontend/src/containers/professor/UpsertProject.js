import React, { Component } from 'react';
import Header from '../../shared/Header.js'
import SidePanel from '../../shared/SidePanel.js'
import '../../styles/professor/Form.css';
import '../../styles/professor/UpsertProject.css';

/**
* Form where professors can add a new project or update an existing project.
*/
class ProfessorUpsertProject extends Component {

  componentWillMount() {
    if (localStorage.getItem('role') === "" || localStorage.getItem('token') === "") {
      this.props.history.push('/login');
    }
  }

  projects(class_id) {
    this.props.history.push('/classes/' + this.props.match.params.class_id);
  }

  getFile(upload, filename) {
    var x = document.getElementById(upload).value;
    if (x === "") {
      document.getElementById(filename).innerHTML = "";
    } else {
      document.getElementById(filename).innerHTML = "*" + x.replace(/^.*\\/, "");
    }
  }

  createProject(e) {
    let token = localStorage.getItem('token');
    let self = this
    e.preventDefault();

    let deadline = self.refs.month.value + '-' + self.refs.day.value + '-' + self.refs.year.value + ' ' + self.refs.hour.value + ':' + self.refs.minute.value;
    console.log(deadline);

    var data = new FormData();
    data.append('name', self.refs.name.value);
    data.append('description', self.refs.description.value);
    data.append('grade_script', self.refs.grading.files[0]);
    data.append('sanity_script', self.refs.sanity.files[0]);
    data.append('language', 'C++');
    data.append('deadline', deadline);

    console.log(data);
    fetch('http://grade-portal-api.herokuapp.com/api/v1.0/classes/' + self.props.match.params.class_id + '/assignments', {
      method: 'POST',
      headers: {
        'Authorization': 'Bearer ' + token
      },
      body: data
      })
    .then((response) => response.json())
    .then(function (responseJSON) {
      console.log(responseJSON);
      if (responseJSON.message !== 'Success') {
        self.refs.error.innerHTML = responseJSON.message;
      } else {
        self.props.history.push('/classes/' + self.props.match.params.class_id);
      }
    });
  }

  constrainLength(id) {
    var x = document.getElementById(id);
    if (x.value.length > x.maxLength)
      x.value = x.value.slice(0, x.maxLength);
  }

  render() {
    return (
      <div>
        <SidePanel />
        <div className="page">
          { window.location.href.substr(window.location.href.lastIndexOf('/') + 1) === "create" ?
            <Header title="Welcome!" path={["Classes", ["Projects", this.props.match.params.class_id], "Create Project"]} />
            :
            <Header title="Welcome!" path={["Classes", ["Projects", this.props.match.params.class_id], ["Edit Project", this.props.match.params.class_id, this.props.match.params.project_id]]} />
          }
          <p ref="error" className="red"></p>
            <div className="create-form">
              <form onSubmit={(e) => this.createProject(e)} encType="multipart/form-data" method="post">
                <label className="upsert-label"><b>Project Name</b></label>
                <input ref="name" type="text" placeholder="Enter project name"/>
                
                <label className="upsert-label"><b>Project Description</b></label>
                <textarea ref="description" placeholder="Enter short description of your project" rows="3" cols="40"/>

                <label className="upsert-label"><b>Upload Grading Script</b></label>
                <div className="upload-btn-wrapper">
                  <input ref="grading" id="upload" className="btn" type="file" name="myfile" onChange={() => this.getFile('upload', 'filename')} accept=".sh"/>
                  <button className="btn">Upload .sh</button>
                  <label className="filename" id="filename"></label>
                </div>

                <label className="upsert-label"><b>Upload Sanity Testing Script</b></label>
                <div className="upload-btn-wrapper">
                  <input ref="sanity" id="upload2" className="btn" type="file" name="myfile" onChange={() => this.getFile('upload2', 'filename2')} accept=".sh"/>
                  <button className="btn">Upload .sh</button>
                  <label className="filename" id="filename2"></label>
                </div>

                <div className="deadline-wrapper">
                  <label className="upsert-label"><b>Project Deadline</b></label>
                  <input type="number" id="month" onInput={() => this.constrainLength("month")} ref="month" placeholder="MM" maxLength="2"/> &nbsp; / &nbsp;
                  <input type="number" id="day" onInput={() => this.constrainLength("day")} ref="day" placeholder="DD" maxLength="2"/> &nbsp; / &nbsp;
                  <input type="number" id="year" onInput={() => this.constrainLength("year")} ref="year" placeholder="YY" maxLength="2"/> &nbsp; &nbsp; &nbsp;
                  <input type="number" id="hour" onInput={() => this.constrainLength("hour")} ref="hour" placeholder="00" maxLength="2"/> &nbsp; : &nbsp;
                  <input type="number" id="minute" onInput={() => this.constrainLength("minute")} ref="minute" placeholder="00" maxLength="2"/>
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
