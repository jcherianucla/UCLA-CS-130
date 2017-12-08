import React, { Component } from 'react';
import Header from '../../shared/Header.js'
import SidePanel from '../../shared/SidePanel.js'
import '../../styles/professor/Analytics.css';

/** 
* Page displaying analytics for a project (i.e. grade distribution)
*/
class ProfessorAnalytics extends Component {

  back() {
    this.props.history.goBack();
  }

  render() {
    return (
      <div>
        <Header title="Welcome!" path={["Login", "Classes", "Projects", "Analytics"]}/>
        <SidePanel />
        <h1>Professor Analytics</h1>
        <button onClick={() => this.back()}>
          Back
        </button>
      </div>
    );
  }
}

export default ProfessorAnalytics;
