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

  constructor(props) {
    super(props);
    this.state = {
      'projects': [],
      'class_name': ''
    }
  }

  componentWillMount() {
    if (localStorage.getItem('role') === "" || localStorage.getItem('token') === "") {
      this.props.history.push('/login');
    }
    this.loadCurrentClass();
    this.loadCards(this.props.match.params.class_id);
  }

  loadCurrentClass() {
    let token = localStorage.getItem('token');
    let self = this
    fetch('http://grade-portal-api.herokuapp.com/api/v1.0/classes/' + self.props.match.params.class_id, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': 'Bearer ' + token,
        'Accept': 'application/json'
      },
    })
    .then((response) => response.json())
    .then(function (responseJSON) {
      console.log(responseJSON);
      if (responseJSON.class !== null && responseJSON.class.name !== null) {
        self.setState({'class_name': responseJSON.class.name});
      }
    });
  }

  professorUpdateProjectLink(class_id, project_id) {
    return '/classes/' + class_id + '/projects/' + project_id + '/edit';
  }

  professorUpdateProject(class_id, project_id) {
    this.props.history.push(this.professorUpdateProjectLink(class_id, project_id));
  }

  projectLink(project_id) {
    return ("/classes/" + this.props.match.params.class_id + "/projects/" + project_id);
  }

  loadCards(class_id) {
    let token = localStorage.getItem('token');
    let self = this
    fetch('http://grade-portal-api.herokuapp.com/api/v1.0/classes/' + class_id + '/assignments', {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': 'Bearer ' + token,
        'Accept': 'application/json'
      },
    })
    .then((response) => response.json())
    .then(function (responseJSON) {
      console.log(responseJSON);
      if (responseJSON.message !== 'Success') {
        self.refs.error.innerHTML = responseJSON.message;
      } else {
        if (responseJSON.assignments !== null) {
          self.setState({'projects': responseJSON.assignments});
        }
      }
    });
  }

  render() {
    let self = this;
    return (
      <div>
        <SidePanel />
        <div className="page">
          <Header title="Welcome!" path={["Classes", ["Projects", this.props.match.params.class_id]]} props={this.state}/>
          <p ref="error" className="red"></p>
            <Grid fluid>
              <Row>
                {
                  this.state.projects.map(function(item, key){
                    return(
                      <Col key={item.id}>
                        <div>
                          <ItemCard
                            title={item.name}
                            deleteLink={'http://grade-portal-api.herokuapp.com/api/v1.0/classes/' + self.props.match.params.class_id + '/assignments/'+ item.id}
                            editLink={self.professorUpdateProjectLink(self.props.match.params.class_id, item.id)}
                            link={'/classes/' + self.props.match.params.class_id + '/projects/' + item.id}
                            history={self.props.history}
                            cardText={item.description}
                          />
                        </div>
                      </Col>
                    );
                  })
                }
                { localStorage.getItem('role') === "professor" ?
                  <Col>
                    <div>
                      <ItemCard
                        image={require("../images/plus.png")}
                        history={this.props.history}
                        link={"/classes/" + this.props.match.params.class_id + "/projects/create"}>
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
