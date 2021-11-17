//Clinician's homepage

import React from 'react';
import {Card} from 'react-bootstrap';
import 'bootstrap/dist/css/bootstrap.min.css';
import Header from './Header';
import './Home2.css';
import {useHistory} from 'react-router-dom';
import axios from 'axios';

function Home2(props) {
    console.log(props);
    console.log(props.location.state);
    const token = props.location.state.Token;
    const history = useHistory();
    var firstname = props.location.state.FirstName;
    const viewProfile = event => {
        event.preventDefault();
        // history.push({
        //     pathname: '/profile',
        //     state: props.location.state
        // })
    }
    const viewClients = event => {
        event.preventDefault();
        axios({ method: 'get', url: 'http://127.0.0.1:3000/clinician/clients', headers: { 'Authorization': 'Bearer ' + token } })
                    .then((response) => {
                    console.log("FINAL", response)
                    if(response.status  == 200){
                        history.push({
                            pathname: '/users',
                            state: {"Data": response.data, "Token": token, "Role":props.location.state.Role}
                        })
                    }
                }, (error) => {
                    console.log(error)
                }
            );

    }
    return (
        <>
        <Header header="Clinician's Homepage"/>
            <br></br>
            <br></br>
            <br></br>
            <br></br>
            <div style={{textAlign:"center", marginLeft:"auto", marginRight:"auto"}}>
                <h4>Welcome {firstname}</h4>
                <br></br>
                <h5>What would you like to do today?</h5>
                <button onClick={viewProfile} class="myButton">View my profile</button>
                <br></br>
                <button onClick={viewClients} class="myButton">View my clients</button>

            </div>
        </>
    )
}

export default Home2;
