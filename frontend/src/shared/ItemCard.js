import React, { Component } from 'react';
import {Card, CardMedia, CardHeader, CardText} from 'material-ui';
import MuiThemeProvider from 'material-ui/styles/MuiThemeProvider';
import '../styles/shared/ItemCard.css';

/**
* Represents a card object with a title, subtitle, and description.
* Cards are meant to be clickable to display more information about the object they represent.
*/
class ItemCard extends Component {
  render() {
    return (
      <MuiThemeProvider>
        <Card className="item-card">
          <CardHeader 
            title={this.props.title}
            titleStyle={{'font-size':'35px', 'color': '#43A5FD', 'font-family': 'Circular-Book'}}
          />
          <CardText
            style={{'padding-top': '0px', 'font-size':'15px', 'color': '#868687', 'font-family': 'Circular-Book'}}
          >
            {this.props.cardText}
          </CardText>
          <CardMedia> 
            <img src={this.props.plus}/>
          </CardMedia>
        </Card>
      </MuiThemeProvider>
    );
  }
}

export default ItemCard;
