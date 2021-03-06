import React, { Component } from 'react';
import Header from '../../shared/Header.js'
import SidePanel from '../../shared/SidePanel.js'
import '../../styles/professor/Form.css';
import '../../styles/professor/UpsertClass.css';

/**
* Form where professors can add a new class or update an existing class.
*/

class ProfessorUpsertClass extends Component {

  constructor(props) {
    super(props);
    this.state = {
      'projects': [],
      'class_name': ''
    }
    this.loadCurrentClass();
  }

  componentWillMount() {
    if (localStorage.getItem('role') === "" || localStorage.getItem('token') === "") {
      this.props.history.push('/login');
    }
  }

  loadCurrentClass() {
    if (this.props.match.params.class_id) {
      let token = localStorage.getItem('token');
      let self = this
      fetch('http://grade-portal-api.herokuapp.com/api/v1.0/classes/' + self.props.match.params.class_id, {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': 'Bearer ' + token,
          'Accept': 'application/json'
        },
      })
      .then((response) => response.json())
      .then(function (responseJSON) {
        console.log(responseJSON);
        if (responseJSON.class !== null && responseJSON.class.name !== null) {
          self.setState({'class_name': responseJSON.class.name});
          self.refs.name.value = responseJSON.class.name;
          self.refs.description.value = responseJSON.class.description;
        }
      });
    }
  }

  getFile() {
    var x = document.getElementById("upload").value;
    if (x === "") {
      document.getElementById("filename").innerHTML = "";
    } else {
      document.getElementById("filename").innerHTML = "*" + x.replace(/^.*\\/, "");
    }
  }

  upsertClass(e, isCreate) {
    let token = localStorage.getItem('token');
    let self = this
    e.preventDefault();
    var data = new FormData();
    data.append('name', self.refs.name.value);
    data.append('description', self.refs.description.value);
    data.append('myfile', self.refs.myfile.files[0]);

    let method = isCreate ? 'POST' : 'PUT';
    let id_for_path = isCreate ? '' : '/' + + self.props.match.params.class_id;
    console.log(data);
    fetch('http://grade-portal-api.herokuapp.com/api/v1.0/classes' + id_for_path, {
      method: method,
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
        self.props.history.push('/classes');
      }
    });
  }

  render() {
    let isCreate = window.location.href.substr(window.location.href.lastIndexOf('/') + 1) === "create";
    return (
      <div>
        <SidePanel />
        <div className="page">
          { window.location.href.substr(window.location.href.lastIndexOf('/') + 1) === "create" ?
            <Header title="Welcome!" path={["Classes", "Create Class"]} props={this.state}/>
            :
            <Header title="Welcome!" path={["Classes", ["Edit Class", this.props.match.params.class_id]]} props={this.state}/>
          }
            <p ref="error" className="red"></p>
            <div className="create-form">
              <form onSubmit={(e) => this.upsertClass(e, isCreate)} encType="multipart/form-data" method={isCreate ? "post" : "put"}>
                <label className="upsert-label"><b>Class Name</b></label>
                <input ref="name" name="name" type="text" placeholder="Enter class name"/>
                
                <label className="upsert-label"><b>Class Description</b></label>
                <textarea ref="description" name="description" placeholder="Enter short description of your class" rows="3" cols="40"/>

                <label className="upsert-label"><b>Upload Student Roster</b></label>
                <div className="upload-btn-wrapper">
                  <input ref="myfile" name="myfile" id="upload" className="btn" type="file" onChange={() => this.getFile()} accept=".csv"/>
                  <button className="btn">Upload .csv</button>
                  <label className="filename" id="filename"></label>
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

export default ProfessorUpsertClass;
