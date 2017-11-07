import React, { Component } from 'react';
import '../../styles/professor/UpsertClass.css';

class ProfessorUpsertClass extends Component {

  back() {
    this.props.history.goBack();
  }

  render() {
    return (
      <div>
        <h1>Professor Create/Edit Class</h1>
        <button onClick={() => this.back()}>
          Back
        </button>
      </div>
    );
  }
}

export default ProfessorUpsertClass;
