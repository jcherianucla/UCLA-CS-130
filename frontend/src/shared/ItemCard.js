import React, { Component } from 'react';
import {Card, CardHeader, CardText} from 'material-ui';
import MuiThemeProvider from 'material-ui/styles/MuiThemeProvider';
import '../styles/shared/ItemCard.css';

class ItemCard extends Component {
  render() {
    return (
      <MuiThemeProvider>
        <Card>
          <CardHeader
            title="Without Avatar"
            subtitle="Subtitle"
          />
          <CardText>
            Lorem ipsum dolor sit amet, consectetur adipiscing elit.
            Donec mattis pretium massa. Aliquam erat volutpat. Nulla facilisi.
            Donec vulputate interdum sollicitudin. Nunc lacinia auctor quam sed pellentesque.
            Aliquam dui mauris, mattis quis lacus id, pellentesque lobortis odio.
          </CardText>
        </Card>
      </MuiThemeProvider>
    );
  }
}

export default ItemCard;
