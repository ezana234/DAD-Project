//Client's homepage

import React, {useState} from 'react';
import {Card} from 'react-bootstrap';
import 'bootstrap/dist/css/bootstrap.min.css';
import Header from './Header';
import './Home.css';
import { Link, useHistory} from 'react-router-dom';
import axios from 'axios';
import TodoButton from './TodoButton';

function Home(props) {
    const history = useHistory();
    console.log("props: ", props)
    var firstname = props.location.state.FirstName;
    const viewProfile = event => {
        event.preventDefault();
        history.push({
            pathname: '/profile',
            state: props.location.state
        })
    }
    const viewSafetyPlan = event =>{
        event.preventDefault();
        axios({ method: 'get', url: 'http://127.0.0.1:3000/client/safetyplan', headers: { 'Authorization': 'Bearer ' + props.location.state.Token }})
        .then((response) => {
                    console.log("Safety plan", response.data)
                    history.push({
                        pathname: '/safetyplan',
                        state: response.data
                    })
                    }, (error) => {
                        console.log("Error"+error)
                    }
                );
    }
    return (
        <>
            <Header header="Client's Homepage"/>

            <div className="container-center-horizontal">
                <div className="client-home-screenscreen">
                    <h1 className="place">
                        Welcome! {firstname}
                    </h1>

                    <div className="text-1">
                        What would you like to do today
                    </div>

                    <TodoButton onClick={viewProfile}>View my profile</TodoButton>

                    <TodoButton className={"todo-button-1"} onClick={viewSafetyPlan}>
                        View my safety plan
                    </TodoButton>

                    <TodoButton className={"todo-button-2"}>
                        View my appointments
                    </TodoButton>
                </div>
            </div>
        </>
    )
}

export default Home;
