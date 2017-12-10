import React, { Component } from 'react';
import { Grid, Row, Col } from 'react-flexbox-grid';
import Header from '../shared/Header.js'
import SidePanel from '../shared/SidePanel.js'
import ItemCard from '../shared/ItemCard.js'
import '../styles/Classes.css';
import '../styles/shared/Page.css';

/** 
* Displays a list of ItemCards representing all of the classes 
* a student or professor is taking or teaching, respectively.
* Clicking on a class will take you to the projects page for that class. 
*/
class Classes extends Component {

  professorUpdateClass(class_id) {
    this.props.history.push('/classes/' + class_id + '/edit');
  }

  render() {
    return (
      <div>
        <SidePanel />
        <div className="page">
          <Header title="Welcome!" path={["Classes"]} />
          <Grid fluid>
              <Row>
                <Col>
                  <div>
                    <ItemCard
                      title='CS 31'
                      link='/classes/1'
                      history={this.props.history}
                      cardText='Introductory computer science class at UCLA, aimed at teaching the fundamentals of C++'
                    />
                  </div>
                </Col>
                <Col>
                  <div>
                    <ItemCard
                      title='CS 31'
                      link='/classes/2'
                      history={this.props.history}
                      cardText='Introductory computer science class at UCLA, aimed at teaching the fundamentals of C++'
                    />
                  </div>
                </Col>
                <Col>
                  <div>
                    <ItemCard
                      title='CS 31'
                      link='/classes/3'
                      history={this.props.history}
                      cardText='Introductory computer science class at UCLA, aimed at teaching the fundamentals of C++'
                    />
                  </div>
                </Col>
                <Col>
                  <div>
                    <ItemCard
                      title='CS 31'
                      link='/classes/4'
                      history={this.props.history}
                      cardText='Introductory computer science class at UCLA, aimed at teaching the fundamentals of C++'
                    />
                  </div>
                </Col>
                <Col>
                  <div>
                    <ItemCard
                      title='CS 31'
                      link='/classes/5'
                      history={this.props.history}
                      cardText='Introductory computer science class at UCLA, aimed at teaching the fundamentals of C++'
                    />
                  </div>
                </Col>
                { localStorage.getItem('role') === "professor" ?
                  <Col>
                    <div>
                      <ItemCard
                        image={require("../images/plus.png")}
                        history={this.props.history}
                        link="/classes/create">
                      </ItemCard>
                    </div>
                  </Col>
                  :
                  <div />
                }
              </Row>
            </Grid>
          
        </div>
 
      </div>
    );
  }
}

export default Classes;
