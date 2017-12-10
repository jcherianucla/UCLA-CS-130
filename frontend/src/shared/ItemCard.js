import React, { Component } from 'react';
import {Card, CardActions, CardMedia, CardHeader, CardText, FlatButton} from 'material-ui';
import MuiThemeProvider from 'material-ui/styles/MuiThemeProvider';
import '../styles/shared/ItemCard.css';

/**
* Represents a card object with a title, subtitle, and description.
* Cards are meant to be clickable to display more information about the object they represent.
*/
class ItemCard extends Component {

  go(e) {
    e.preventDefault();
    if (this.props.link) {
      this.props.history.push(this.props.link);
    }
  }

  editProjectOrClass(editLink) {
    this.props.history.push(editLink);
  }

  reloadClasses() {
    window.location.reload();
    console.log('Reloading classes page');
  }

  deleteProjectOrClass(deleteLink) {
    let token = localStorage.getItem('token');
    let self = this
    fetch(deleteLink, {
      method: 'DELETE',
      headers: {
        'Authorization': 'Bearer ' + token,
      },
    })
    .then(response => response.json())
    .then(function (responseJSON) {
      console.log(responseJSON);
      if (responseJSON.message === 'Success') {
        self.reloadClasses();
      }
    })
  }

  render() {
    return (
      <MuiThemeProvider>
          <Card className="item-card">
            { !this.props.image ?
              <div>
                <a href={this.props.link} onClick={(e) => this.go(e)} >
                  <CardHeader
                    title={this.props.title}
                    titleStyle={{'fontSize':'35px', 'color': '#43A5FD', 'fontFamily': 'Circular-Book'}}
                  />
                </a>
                <CardText className="card-text"
                  style={{'paddingTop': '0px', 'fontSize':'15px', 'color': '#868687', 'fontFamily': 'Circular-Book'}}
                >
                  {this.props.cardText}
                </CardText>
                { localStorage.getItem('role') === "professor" ?
                    <CardActions>
                      <FlatButton label="Delete" onClick={() => this.deleteProjectOrClass(this.props.deleteLink)} style={{'float': 'right', 'color': '#43A5FD', 'fontFamily': 'Circular-Book'}}/>
                      <FlatButton label="Edit" onClick={() => this.editProjectOrClass(this.props.editLink)} style={{'float': 'right', 'color': '#43A5FD', 'fontFamily': 'Circular-Book'}}/>
                    </CardActions> 
                    :
                    <CardActions /> 
                }
              </div>
              :
              <CardMedia className="center-item-card-vertically"> 
                <a href={this.props.link} onClick={(e) => this.go(e)} >
                  <div className="center-item-card-horizontally">
                      <img className="item-card-image" src={this.props.image} alt="" />
                  </div>
                </a>
              </CardMedia>
            }
          </Card>
      </MuiThemeProvider>
    );
  }
}

export default ItemCard;
