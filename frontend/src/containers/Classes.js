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
        self.createCards(responseJSON.classes);
      }
    });
  }

  createCards(classes) {
    classes.map(function(item, key){
      return(
        <Col>
          <div>
            <ItemCard
              title={item.name}
              editLink={this.professorUpdateProjectLink(item.id)}
              link={'/classes/' + item.id}
              history={this.props.history}
              cardText={item.description}
            />
          </div>
        </Col>
      );
    })
  }

  render() {
    return (
      <div>
        <SidePanel />
        <div className="page">
          <Header title="Welcome!" path={["Classes"]} />
          <br /> <br />
          <p ref="error" className="red"></p>
          <Grid fluid>
            <Row>
              { this.loadCards() }
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