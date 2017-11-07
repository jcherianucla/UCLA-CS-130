import React, { Component } from 'react';
import '../styles/shared/Header.css';

class Header extends Component {
  render() {
    return (
      <div className="header">
        <p>This is an example header. Code that is repeated on multiple pages should be implemented as shared components</p>
      </div>
    );
  }
}

export default Header;
