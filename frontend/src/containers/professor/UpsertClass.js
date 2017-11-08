import React, { Component } from 'react';
import Header from '../../shared/Header.js'
import SidePanel from '../../shared/SidePanel.js'
import '../../styles/professor/UpsertClass.css';

class ProfessorUpsertClass extends Component {

  back() {
    this.props.history.goBack();
  }

  render() {
    return (
      <div>
        <Header />
        <SidePanel />
        <h1>Professor Create/Edit Class</h1>
        <button onClick={() => this.back()}>
          Back
        </button>
      </div>
    );
  }
}

export default ProfessorUpsertClass;
