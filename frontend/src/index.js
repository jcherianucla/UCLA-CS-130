import React from 'react';
import ReactDOM from 'react-dom';
import { AppContainer } from 'react-hot-loader'
import './styles/index.css';
import { BrowserRouter, Switch, Route } from 'react-router-dom'
import Landing from './containers/Landing';
import ProfessorLogin from './containers/professor/Login';
import StudentLogin from './containers/student/Login';

const base = document.querySelector('base')
const baseHref = base ? base.getAttribute('href') : '/'

ReactDOM.render(
  <AppContainer>
    <BrowserRouter basename={baseHref.replace(/\/$/, '')}>
      <Switch>
        <Route path="/" exact={true} component={Landing} />
        <Route path="/professor/login" exact={true} component={ProfessorLogin} />
        <Route path="/student/login" exact={true} component={StudentLogin} />
      </Switch>
    </BrowserRouter>
  </AppContainer>,
  document.getElementById('root')
  );
