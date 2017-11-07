import React, { Component } from 'react';
import Header from '../../shared/Header.js'
import '../../styles/professor/Login.css';

class ProfessorLogin extends Component {

  back() {
    this.props.history.goBack();
  }

  classes() {
    this.props.history.push('/classes');
  }

  render() {
    return (
      <div>
        <Header />
        <h1>Professor Login</h1>
        <button onClick={() => this.back()}>
          Back
        </button>
        <button onClick={() => this.classes()}>
          Classes
        </button>
      </div>
    );
  }
}

export default ProfessorLogin;
