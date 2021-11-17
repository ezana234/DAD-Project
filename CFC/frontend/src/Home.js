//Client's homepage

import React, {useState} from 'react';
import {Card} from 'react-bootstrap';
import 'bootstrap/dist/css/bootstrap.min.css';
import Header from './Header';
import './Home.css';
import { Link, useHistory} from 'react-router-dom';
import axios from 'axios';



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
        axios({ method: 'get', url: 'http://127.0.0.1:3000/safetyplan', headers: { 'Authorization': 'Bearer ' + props.location.state.Token }, params: {"userID": props.location.state.Data.UserID} })
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
                <button onClick={viewSafetyPlan} class="myButton">View my safety plan</button>
                <br></br>
                <button class="myButton">Add emergency contact</button>
            </div>
            
        </>
    )
}

export default Home;
