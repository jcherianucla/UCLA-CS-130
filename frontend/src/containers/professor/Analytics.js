import React, { Component } from 'react';
import Header from '../../shared/Header.js'
import '../../styles/professor/Analytics.css';

class ProfessorAnalytics extends Component {

  back() {
    this.props.history.goBack();
  }

  render() {
    return (
      <div>
        <Header />
        <h1>Professor Analytics</h1>
        <button onClick={() => this.back()}>
          Back
        </button>
      </div>
    );
  }
}

export default ProfessorAnalytics;
