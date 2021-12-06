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
    var firstname = props.location.state.Data.FirstName;
    const viewProfile = event => {
        event.preventDefault();
        history.push({
            pathname: '/profile',
            state: {"Data":props.location.state, "Token": props.location.state.Token, "Role":props.location.state.Role, "oldData":props.location.state, "prev":props.location.state, "w":0}
        })
    }
    const viewSafetyPlan = event =>{
        event.preventDefault();
        axios({ method: 'get', url: 'http://127.0.0.1:3000/client/safetyplan', headers: { 'Authorization': 'Bearer ' + props.location.state.Token }})
        .then((response) => {
                    console.log("Safety plan", response.data)
                    history.push({
                        pathname: '/safetyplan',
                        state: {"Data": response.data, "Token": props.location.state.Token, "Role":props.location.state.Role, "oldData":props.location.state, "prev":props.location.state, "s":0}
                    })
                    }, (error) => {
                        console.log("Error"+error)
                    }
                );
    }

    const viewAppointments = event =>{
        event.preventDefault();
        axios({ method: 'get', url: 'http://127.0.0.1:3000/client/appointments', headers: { 'Authorization': 'Bearer ' + props.location.state.Token }})
        .then((response) => {
                    console.log("Appointments", response.data)
                    history.push({
                        pathname: '/appointments',
                        state: {"Data":response.data, "Token": props.location.state.Token, "Role":props.location.state.Role, "oldData":props.location.state, "prev":props.location.state}
                    })
                    }, (error) => {
                        console.log("Error"+error)
                    }
                );
    }
    return (
        <>
            <Header header="Client's Homepage" role={props.location.state.Role} oldData={props.location.state}/>

            <div className="container-center-horizontal">
                <div className="client-home-screenscreen">
                    <h1 className="place">
                        {"Welcome "+ firstname}
                    </h1>

                    <div className="text-1">
                        What would you like to do today
                    </div>

                    <TodoButton onClick={viewProfile}>View my profile</TodoButton>

                    <TodoButton className={"todo-button-1"} onClick={viewSafetyPlan}>
                        View my safety plan
                    </TodoButton>

                    <TodoButton className={"todo-button-2"} onClick={viewAppointments}>
                        View my appointments
                    </TodoButton>
                </div>
            </div>
        </>
    )
}

export default Home;
