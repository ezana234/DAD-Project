//Clinician's homepage

import React from 'react';
import {Card} from 'react-bootstrap';
import 'bootstrap/dist/css/bootstrap.min.css';
import Header from './Header';
import './Home.css';
import {useHistory} from 'react-router-dom';
import axios from 'axios';
import TodoButton from './TodoButton';
import ViewAppointments from './ViewAppointments';

function Home2(props) {
    console.log(props);
    console.log(props.location.state);
    const token = props.location.state.Token;
    const history = useHistory();
    var firstname = props.location.state.Data.FirstName;
    const viewProfile = event => {
        event.preventDefault();
        history.push({
            pathname: '/profile',
            state: {"Data":props.location.state, "Token": props.location.state.Token, "Role":props.location.state.Role, "oldData":props.location.state, "prev":props.location.state, "w":0}

        })
    }
    const viewClients = event => {
        event.preventDefault();
        axios({ method: 'get', url: 'http://127.0.0.1:3000/clinician/clients', headers: { 'Authorization': 'Bearer ' + token } })
                    .then((response) => {
                    console.log("FINAL", response)
                    if(response.status  == 200){
                        history.push({
                            pathname: '/users',
                            state: {"Data": response.data, "Token": token, "Role":props.location.state.Role, "oldData":props.location.state, "prev":props.location.state}
                        })
                    }
                }, (error) => {
                    console.log(error)
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
            <Header header="Clinician's Homepage" role={props.location.state.Role} oldData={props.location.state}/>

            <div className="container-center-horizontal">
                <div className="client-home-screenscreen">
                    <h1 className="place">
                        {"Welcome "+firstname}
                    </h1>

                    <div className="text-1">
                        What would you like to do today
                    </div>

                    <TodoButton onClick={viewProfile}>View my profile</TodoButton>

                    <TodoButton className={"todo-button-1"} onClick={viewClients}>
                    View my clients
                    </TodoButton>
                    <TodoButton className={"todo-button-1"} onClick={viewAppointments}>
                    View Appointments
                    </TodoButton>
                </div>
            </div>
        </>
    )
}

export default Home2;
