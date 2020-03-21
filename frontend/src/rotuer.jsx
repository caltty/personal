import login from './login'
import logout from './logout'
import React from 'react';
import { BrowserRouter as Router, Route } from 'react-router-dom';

export default (
    <Router >
        <Route exact path="/" component={login}></Route>
        <Route path="/logout" component={logout}></Route>
    </Router>
)