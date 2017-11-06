import React from 'react';
import ReactDOM from 'react-dom';
import './styles/index.css';
import Login from './containers/Login';
import registerServiceWorker from './utils/registerServiceWorker';

ReactDOM.render(<Login />, document.getElementById('root'));
registerServiceWorker();
