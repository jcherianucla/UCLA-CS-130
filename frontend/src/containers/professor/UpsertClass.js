import React, { Component } from 'react';
import Header from '../../shared/Header.js'
import SidePanel from '../../shared/SidePanel.js'
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

  render() {
    return (
      <div>
        <SidePanel />
        <div className="page">
          <Header title="Welcome!" path="Create/Edit Class" />
            <div class="class-create-form">
              <form onSubmit={() => this.classes()}>
                <label class="upsert-label"><b>Class Name</b></label>
                <input type="text" placeholder="Enter class name"/>
                
                <label class="upsert-label"><b>Class Description</b></label>
                <textarea placeholder="Enter short description of your class" rows="3" cols="40"/>

                <label class="upsert-label"><b>Upload Student Roster</b></label>
                <div class="upload-btn-wrapper">
                  <input type="file" name="myfile" accept=".csv"/>
                  <button class="btn">Upload .csv</button>
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
