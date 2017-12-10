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

  constructor(props) {
    super(props);
    this.state = {
      'classes': []
    }
  }

  componentWillMount() {
    this.loadCards();
  }

  professorUpdateProjectLink(class_id) {
    return '/classes/' + class_id + '/edit';
  }

  professorUpdateClass(class_id) {
    this.props.history.push(this.professorUpdateProjectLink(class_id));
  }

  loadCards() {
    let token = localStorage.getItem('token');
    let self = this
    fetch('http://grade-portal-api.herokuapp.com/api/v1.0/classes', {
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
        if (responseJSON.classes !== null) {
          self.setState({'classes': responseJSON.classes});
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
          <Header title="Welcome!" path={["Classes"]} />
          <br /> <br />
          <p ref="error" className="red"></p>
          <Grid fluid>
            <Row>
              {
                this.state.classes.map(function(item, key){
                  return(
                    <Col>
                      <div>
                        <ItemCard
                          title={item.name}
                          editLink={'/classes/' + item.id + '/edit'}
                          deleteLink={'http://grade-portal-api.herokuapp.com/api/v1.0/classes/' + item.id}
                          link={'/classes/' + item.id}
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