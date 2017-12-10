import React, { Component } from 'react';
import { Grid, Row, Col } from 'react-flexbox-grid';
import Header from '../shared/Header.js'
import SidePanel from '../shared/SidePanel.js'
import ItemCard from '../shared/ItemCard.js'
import '../styles/Classes.css';
import '../styles/shared/Page.css';


class FAQ extends Component {

  professorUpdateClass(class_id) {
    this.props.history.push('/classes/' + class_id + '/edit');
  }

  render() {
    return (
      <div>
        <SidePanel />
        <div className="page">
          <Header title="Frequently Asked Questions" path={["Classes", "FAQ"]} />
          <h1>How do you contact us?</h1>
          <h2>Email us at gradeportal@gmail.com</h2>


        </div>
 
      </div>
    );
  }
}

export default FAQ;
