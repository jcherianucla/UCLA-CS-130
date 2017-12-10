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

  render() {
    return (
      <MuiThemeProvider>
        <a href={this.props.link} onClick={(e) => this.go(e)} >
          <Card className="item-card">
            { !this.props.image ?
              <div>
                <CardHeader 
                  title={this.props.title}
                  titleStyle={{'fontSize':'35px', 'color': '#43A5FD', 'fontFamily': 'Circular-Book'}}
                />
                <CardText
                  style={{'paddingTop': '0px', 'fontSize':'15px', 'color': '#868687', 'fontFamily': 'Circular-Book'}}
                >
                  {this.props.cardText}
                </CardText>
                <CardActions>
                  <FlatButton label="Delete" style={{'float': 'right', 'color': '#43A5FD', 'fontFamily': 'Circular-Book'}}/>
                  <FlatButton label="Edit" style={{'float': 'right', 'color': '#43A5FD', 'fontFamily': 'Circular-Book'}}/>
                </CardActions>
              </div>
              :
              <CardMedia className="center-item-card-vertically"> 
                <div className="center-item-card-horizontally">
                  <img className="item-card-image" src={this.props.image} alt="" />
                </div>
              </CardMedia>
            }
          </Card>
        </a>
      </MuiThemeProvider>
    );
  }
}

export default ItemCard;
