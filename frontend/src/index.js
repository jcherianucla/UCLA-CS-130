import React from 'react';
import ReactDOM from 'react-dom';
import { AppContainer } from 'react-hot-loader';
import './styles/index.css';
import './styles/common.css';
import { BrowserRouter, Switch, Route } from 'react-router-dom';

import Classes from './containers/Classes';
import Landing from './containers/Landing';
import Login from './containers/Login';
import Project from './containers/Project';
import Projects from './containers/Projects';
import ProfessorUpsertClass from './containers/professor/UpsertClass';
import ProfessorUpsertProject from './containers/professor/UpsertProject';

const base = document.querySelector('base')
const baseHref = base ? base.getAttribute('href') : '/'

ReactDOM.render(
  <AppContainer>
    <BrowserRouter basename={baseHref.replace(/\/$/, '')}>
      <Switch>
        <Route path="/" exact={true} component={Landing} />
        <Route path="/login" exact={true} component={Login} />
        <Route path="/classes" exact={true} component={Classes} />
        <Route path="/classes/create" exact={true} component={ProfessorUpsertClass} />
        <Route path="/classes/:class_id" exact={true} component={Projects} />
        <Route path="/classes/:class_id/edit" exact={true} component={ProfessorUpsertClass} />
        <Route path="/classes/:class_id/projects/create" exact={true} component={ProfessorUpsertProject} />
        <Route path="/classes/:class_id/projects/:project_id" exact={true} component={Project} />
        <Route path="/classes/:class_id/projects/:project_id/edit" exact={true} component={ProfessorUpsertProject} />
      </Switch>
    </BrowserRouter>
  </AppContainer>,
  document.getElementById('root')
  );
