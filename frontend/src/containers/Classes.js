import React, { Component } from 'react';
import Header from '../shared/Header.js'
import SidePanel from '../shared/SidePanel.js'
import ItemCard from '../shared/ItemCard.js'
import '../styles/Classes.css';

class Classes extends Component {

  back() {
    this.props.history.goBack();
  }

  projects() {
    this.props.history.push('/projects');
  }

  professorUpsertClass() {
    this.props.history.push('/professor/upsert_class');
  }

  render() {
    return (
      <div>
        <Header />
        <SidePanel />
        <h1>Classes</h1>
        <button onClick={() => this.back()}>
          Back
        </button>
        <button onClick={() => this.projects()}>
          Projects
        </button>
        <button onClick={() => this.professorUpsertClass()}>
          Professor Create/Edit Class
        </button>
        <ItemCard />
      </div>
    );
  }
}

export default Classes;
