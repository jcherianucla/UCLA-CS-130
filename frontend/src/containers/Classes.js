import React, { Component } from 'react';
import Header from '../shared/Header.js'
import Content from '../shared/Content.js'
import SidePanel from '../shared/SidePanel.js'
import ItemCard from '../shared/ItemCard.js'
import '../styles/Classes.css';
import '../styles/shared/Page.css';

class Classes extends Component {

  componentWillMount() {
    console.log(this.props.history);
  }

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
        <SidePanel />
        <div className="page">
          <Header title={`Welcome ${this.props.history.location.state.firstName}`} path="Classes" />
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
 
      </div>
    );
  }
}

export default Classes;
