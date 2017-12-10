import React, { Component } from 'react';
import Header from '../../shared/Header.js'
import SidePanel from '../../shared/SidePanel.js'
import '../../styles/professor/Form.css';
import '../../styles/professor/UpsertClass.css';

/**
* Form where professors can add a new class or update an existing class.
*/

class ProfessorUpsertClass extends Component {

  classes() {
    this.props.history.push('/classes');
  }

  getFile() {
    var x = document.getElementById("upload").value;
    if (x === "") {
      document.getElementById("filename").innerHTML = "";
    } else {
      document.getElementById("filename").innerHTML = "*" + x.replace(/^.*\\/, "");
    }
  }

  createClass(name, description, quarter, year) {
    let self = this
    fetch('http://grade-portal-api.herokuapp.com/api/v1.0/classes', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        name: name,
        description: description,
        quarter: "Fall",
        year: "2017"
      })
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
    return (
      <div>
        <SidePanel />
        <div className="page">
          { window.location.href.substr(window.location.href.lastIndexOf('/') + 1) === "create" ?
            <Header title="Welcome!" path={["Classes", "Create Class"]} />
            :
            <Header title="Welcome!" path={["Classes", ["Edit Class", this.props.match.params.class_id]]} />
          }
            <br /><br />
            <p ref="error" className="red"></p>
            <div className="create-form">
              <form onSubmit={() => this.classes()}>
                <label className="upsert-label"><b>Class Name</b></label>
                <input type="text" placeholder="Enter class name"/>
                
                <label className="upsert-label"><b>Class Description</b></label>
                <textarea placeholder="Enter short description of your class" rows="3" cols="40"/>

                <label className="upsert-label"><b>Upload Student Roster</b></label>
                <div className="upload-btn-wrapper">
                  <input id="upload" className="btn" type="file" name="myfile" onChange={() => this.getFile()} accept=".csv"/>
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
