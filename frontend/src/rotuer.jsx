import login from './login';
import logout from './logout';
import config from './config';
import serverConfig from "./serverConfig";
import redisConfig from './redisConfig';
import jwtConfig from './jwtConfig';
import ldapConfig from './ldapConfig';
import React from 'react';
import { BrowserRouter as Router, Route } from 'react-router-dom';

export default (
    <Router >
        <Route exact path="/" component={login}></Route>
        <Route path="/logout" component={logout}></Route>
        <Route path="/config" component={config}></Route>
        <Route path="/serverConfig" component={serverConfig}></Route>
        <Route path="/redisConfig" component={redisConfig}></Route>
        <Route path="/jwtConfig" component={jwtConfig}></Route>
        <Route path="/ldapConfig" component={ldapConfig}></Route>
    </Router>
)