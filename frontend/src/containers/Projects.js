import React, { Component } from 'react';
import { Grid, Row, Col } from 'react-flexbox-grid';
import Header from '../shared/Header.js'
import SidePanel from '../shared/SidePanel.js'
import ItemCard from '../shared/ItemCard.js'
import '../styles/Projects.css';

/**
* Displays a list of ItemCards representing the projects available for a class. 
* Students can click on an ItemCard to submit or view their submission,
* while professors click on ItemCards to view project analytics.
* Professors can also update or insert new projects from this page. 
*/
class Projects extends Component {

  back() {
    this.props.history.goBack();
  }

  studentSubmission() {
    this.props.history.push('/student/submission');
  }

  professorUpsertProject() {
    this.props.history.push('/professor/upsert_project');
  }

  professorAnalytics() {
    this.props.history.push('/professor/analytics');
  }

  projectLink() {
    if (localStorage.getItem('role') === "professor") {
      return ('/professor/upsert_project');    
    } else {
      return ('/student/submission');
    }
  }

  render() {
    return (
      <div>
        <SidePanel />
        <div className="page">
          <Header title="Welcome!" path="Projects" />

          <Grid fluid>
              <Row>
                <Col>
                  <div>
                    <ItemCard
                      title='Project 1'
                      link={this.projectLink()}
                      history={this.props.history}
                      cardText='Project 1 description'
                    />
                  </div>
                </Col>
                <Col>
                  <div>
                    <ItemCard
                      title='Project 2'
                      link={this.projectLink()}
                      history={this.props.history}
                      cardText='Project 2 description'
                    />
                  </div>
                </Col>
                <Col>
                  <div>
                    <ItemCard
                      title='Project 3'
                      link={this.projectLink()}
                      history={this.props.history}
                      cardText='Project 3 description'
                    />
                  </div>
                </Col>
                <Col>
                  <div>
                    <ItemCard
                      title='Project 4'
                      link={this.projectLink()}
                      history={this.props.history}
                      cardText='Project 4 description'
                    />
                  </div>
                </Col>
                <Col>
                  <div>
                    <ItemCard
                      title='Project 5'
                      link={this.projectLink()}
                      history={this.props.history}
                      cardText='Project 5 description'
                    />
                  </div>
                </Col>
                { localStorage.getItem('role') === "professor" ?
                  <Col>
                    <div>
                      <ItemCard
                        image={require("../images/plus.png")}
                        history={this.props.history}
                        link="/professor/upsert_project">
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

export default Projects;
