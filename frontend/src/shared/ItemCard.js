import React, { Component } from 'react';
import {Card, CardHeader, CardText} from 'material-ui';
import MuiThemeProvider from 'material-ui/styles/MuiThemeProvider';
import '../styles/shared/ItemCard.css';

class ItemCard extends Component {
  render() {
    return (
      <MuiThemeProvider>
        <Card className="item-card">
          <CardHeader
            title="CS 31"
            subtitle="Subtitle"
          />
          <CardText>
            Introductory computer science class at UCLA, aimed at teaching the fundamentals of C++
          </CardText>
        </Card>
      </MuiThemeProvider>
    );
  }
}

export default ItemCard;
