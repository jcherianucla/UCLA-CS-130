import React from 'react';
import ReactDOM from 'react-dom';
import { AppContainer } from 'react-hot-loader';
import './styles/index.css';
import './styles/common.css';
import { BrowserRouter, Switch, Route } from 'react-router-dom';
import Landing from './containers/Landing';
import Classes from './containers/Classes';
import Projects from './containers/Projects';
import ProfessorAnalytics from './containers/professor/Analytics';
import ProfessorLogin from './containers/professor/Login';
import ProfessorUpsertClass from './containers/professor/UpsertClass';
import ProfessorUpsertProject from './containers/professor/UpsertProject';
import StudentLogin from './containers/student/Login';
import StudentSubmission from './containers/student/Submission';

const base = document.querySelector('base')
const baseHref = base ? base.getAttribute('href') : '/'

ReactDOM.render(
  <AppContainer>
    <BrowserRouter basename={baseHref.replace(/\/$/, '')}>
      <Switch>
        <Route path="/" exact={true} component={Landing} />
        <Route path="/classes" exact={true} component={Classes} />
        <Route path="/projects" exact={true} component={Projects} />
        <Route path="/professor/analytics" exact={true} component={ProfessorAnalytics} />
        <Route path="/professor/login" exact={true} component={ProfessorLogin} />
        <Route path="/professor/upsert_class" exact={true} component={ProfessorUpsertClass} />
        <Route path="/professor/upsert_project" exact={true} component={ProfessorUpsertProject} />
        <Route path="/student/login" exact={true} component={StudentLogin} />
        <Route path="/student/submission" exact={true} component={StudentSubmission} />
      </Switch>
    </BrowserRouter>
  </AppContainer>,
  document.getElementById('root')
  );
