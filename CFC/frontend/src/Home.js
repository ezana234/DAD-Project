//Client's homepage

import React, {useState} from 'react';
import {Card} from 'react-bootstrap';
import 'bootstrap/dist/css/bootstrap.min.css';
import Header from './Header';
import './Home.css';
import { Link, useHistory} from 'react-router-dom';



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
                <button class="myButton">View my safety plan</button>
                <br></br>
                <button class="myButton">Add emergency contact</button>
            </div>
            
        </>
    )
}

export default Home;
