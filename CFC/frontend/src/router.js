import React, { Component } from 'react';
import { BrowserRouter, Switch, Route} from 'react-router-dom';
import Home from './Home';
import App from './App';

function router() {
    return (
        <div>
            <BrowserRouter>
                <Switch>
                    <Route path='/' exact component={App} />
                    <Route path='/home' component={Home} />
                </Switch>
            </BrowserRouter>
        </div>
        
    )
}

export default router
