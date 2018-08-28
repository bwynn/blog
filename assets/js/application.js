require("expose-loader?$!expose-loader?jQuery!jquery");
require("bootstrap/dist/js/bootstrap.bundle.js");

import React, { Component } from 'react';
import { render } from 'react-dom';

class App extends Component {
  render () {
    return (
      <div>Hello from within!</div>
    )
  }
}

render(
  <App />,
  document.getElementById('root')
)
