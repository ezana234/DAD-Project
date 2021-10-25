import React, { Component } from 'react';
import { BrowserRouter, Switch, Route} from 'react-router-dom';
import Home from './Home';
import App from './App';
import Login from './Login';
import Home2 from './Home2';

function router() {
    return (
        <div>
            <BrowserRouter>
                <Switch>
                    <Route path='/' exact component={Login} />
                    <Route path='/home' component={Home} />
                    <Route path='/clinicianHome' component={Home2} />
                </Switch>
            </BrowserRouter>
        </div>
        
    )
}

export default router
