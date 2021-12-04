import React, { Component } from 'react';
import { BrowserRouter, Switch, Route} from 'react-router-dom';
import Home from './Home';
import App from './App';
import Login from './Login';
import Home2 from './Home2';
import Profile from './Profile';
import Users from './Users';
import SafetyPlan from './SafetyPlan';
import SignUp from './SignUp';
import AddSafetyPlan from './AddSafetyPlan';
import ModifySafetyPlan from './ModifySafetyPlan';

function router() {
    return (
        <div>
            <BrowserRouter>
                <Switch>
                    <Route exact path='/' exact component={Login} />
                    <Route exact path='/signUp' exact component={SignUp} />
                    <Route exact path='/clientHome' component={Home} />
                    <Route exact path='/clinicianHome' component={Home2} />
                    <Route exact path = '/profile' component={Profile} />
                    <Route exact path = '/users' component = {Users} />
                    <Route exact path = '/safetyplan' component = {SafetyPlan} />
                    <Route exact path = '/addsafetyplan' component ={AddSafetyPlan} />
                    <Route exact path = '/modifysafetyplan' component ={ModifySafetyPlan} />

                    
                </Switch>
            </BrowserRouter>
        </div>
        
    )
}

export default router
