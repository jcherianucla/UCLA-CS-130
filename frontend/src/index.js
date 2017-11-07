import React from 'react';
import ReactDOM from 'react-dom';
import { AppContainer } from 'react-hot-loader'
import './styles/index.css';
import Landing from './containers/Landing';
// import registerServiceWorker from './utils/registerServiceWorker';

ReactDOM.render(
  <AppContainer>
    <Landing />
  </AppContainer>,
  document.getElementById('root')
  );
// registerServiceWorker();
