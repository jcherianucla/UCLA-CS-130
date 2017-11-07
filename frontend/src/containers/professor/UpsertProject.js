import React, { Component } from 'react';
import Header from '../../shared/Header.js'
import '../../styles/professor/UpsertProject.css';

class ProfessorUpsertProject extends Component {

  back() {
    this.props.history.goBack();
  }

  render() {
    return (
      <div>
        <Header />
        <h1>Professor Create/Edit Project</h1>
        <button onClick={() => this.back()}>
          Back
        </button>
      </div>
    );
  }
}

export default ProfessorUpsertProject;