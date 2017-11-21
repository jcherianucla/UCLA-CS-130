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

  render() {
    return (
      <div>
        <SidePanel />
        <div className="page">
          <Header title="Welcome!" path="Create/Edit Class" />
            {/* TODO: Change class creation form to match designs */}
            <form id="class-form">
              <div className="class-form-group">
                <input className="class-form-input" type="text" required="required" />
                <span className="class-form-bar"></span>
                <label className="class-form-label">Class Name</label>
              </div>
              <div className="class-form-group">
                <input className="class-form-input secret" type="text" required="required"/>
                <span className="class-form-bar"></span>
                <label className="class-form-label">Class Description</label>
              </div>
            </form>
            <div class="upload-btn-wrapper">
              <button class="btn">Upload .csv</button>
              <input type="file" name="myfile" accept=".csv"/>
            </div>
        </div>
      </div>
    );
  }
}

export default ProfessorUpsertClass;
