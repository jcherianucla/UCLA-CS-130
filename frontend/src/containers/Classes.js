import React, { Component } from 'react';
import { Grid, Row, Col } from 'react-flexbox-grid';
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

  displayCreateCard() {
    if (this.props.history.location.state.type === "professor") {
      return (<ItemCard plus="http://www.freepngimg.com/download/dog/1-2-dog-png-10.png"></ItemCard>);
    }
      
  }

  render() {
    return (
      <div>
        <SidePanel />
        <div className="page">
          <Header title={`Welcome ${this.props.history.location.state.firstName}`} path="Classes" />

          <Grid fluid>
              <Row>
                <Col>
                  <div>
                    <ItemCard title='CS 31' cardText='Introductory computer science class at UCLA, aimed at teaching the fundamentals of C++'/>
                  </div>
                </Col>
                <Col>
                  <div>
                    <ItemCard title='CS 31' cardText='Introductory computer science class at UCLA, aimed at teaching the fundamentals of C++'/>
                  </div>
                </Col>
                <Col>
                  <div>
                    <ItemCard title='CS 31' cardText='Introductory computer science class at UCLA, aimed at teaching the fundamentals of C++'/>
                  </div>
                </Col>
                <Col>
                  <div>
                    <ItemCard title='CS 31' cardText='Introductory computer science class at UCLA, aimed at teaching the fundamentals of C++'/>
                  </div>
                </Col>
                <Col>
                  <div>
                    <ItemCard title='CS 31' cardText='Introductory computer science class at UCLA, aimed at teaching the fundamentals of C++'/>
                  </div>
                </Col>
                <Col>
                  <div>
                    {this.displayCreateCard()}
                  </div>
                </Col>
              </Row>
            </Grid>
          
        </div>
 
      </div>
    );
  }
}

export default Classes;
