import React, { Component } from 'react';
import { BrowserRouter, Switch, Route} from 'react-router-dom';
import Home from './Home';
import App from './App';
import Login from './Login';

function router() {
    return (
        <div>
            <BrowserRouter>
                <Switch>
                    <Route path='/' exact component={Login} />
                    <Route path='/home' component={Home} />
                </Switch>
            </BrowserRouter>
        </div>
        
    )
}

export default router
