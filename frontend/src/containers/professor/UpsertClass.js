import React, { Component } from 'react';
import Header from '../../shared/Header.js'
import SidePanel from '../../shared/SidePanel.js'
import '../../styles/professor/Form.css';
import '../../styles/professor/UpsertClass.css';

/**
* Form where professors can add a new class or update an existing class.
*/

class ProfessorUpsertClass extends Component {

  back() {
    this.props.history.goBack();
  }

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

  render() {
    return (
      <div>
        <SidePanel />
        <div className="page">
          <Header title="Welcome!" path={["Login", "Classes", "Create/Edit Class"]} />
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
