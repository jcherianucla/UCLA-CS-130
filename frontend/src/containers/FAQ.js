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
          <h1 className="dark-gray">What is GradePortal?</h1>
          <h3 className="gray">GradePortal is a project submission portal where students and professors can easily access all the information they need regarding their classes and assignment. We even have project analytics, automatic grading, and submission feedback!</h3>

          <h1 className="dark-gray">How do you contact us?</h1>
          <h3 className="gray">Email us at gradeportal@gmail.com</h3>

          <h1 className="dark-gray">How do I start?</h1>
          <h3 className="gray">If you're a student, just log in with your Bruin Online Login. If you're a professor, email us at gradeportal@gmail.com and we will add your account for you!</h3>
        </div>
      </div>
    );
  }
}

export default FAQ;
