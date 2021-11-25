import React, {useState} from 'react'
import {Card} from 'react-bootstrap';
import 'bootstrap/dist/css/bootstrap.min.css';
import Header from './Header';
import Search from './Search';
import {withRouter, Link, useHistory } from 'react-router-dom';
import axios from 'axios';
import './Users.css'

function Users(props) {
    const history = useHistory();
    console.log("Users", props)
    const clients = props.location.state.Data;
    console.log("Type: ", typeof(clients))
    const filterClients = (clients, query) => {
        console.log("Inside filter clients method")
        console.log(!query)
        if (!query) {
            return clients;
        }
    
        return clients.filter((client) => {
            const clientName = client.FirstName;
            console.log(clientName);
            //return clientName.includes(query);
            return clientName.startsWith(query);

        });

    };
    
    let query = "";
    const [searchQuery, setSearchQuery] = useState(query || '');
    const { search } = window.location;
    query = searchQuery;
    console.log(query)
    const filteredClients = filterClients(clients, query);
    console.log(searchQuery);

    function viewProfile(client){
        console.log("Inside view Profile")
        console.log(client);
        if(client!=null){
            history.push({
                pathname: '/profile',
                state: {"Data":client}
            })
        }
    }

    function viewSafetyPlan(client){
        console.log("Inside view safety plan")
        console.log(client);
        console.log("Index: "+clients.indexOf(client))
        if(client!=null){
            axios({ method: 'get', url: 'http://127.0.0.1:3000/clinician/safetyplan', headers: { 'Authorization': 'Bearer ' + props.location.state.Token }, params: {"userID": client.UserID} })
            .then((response) => {
                        console.log("Safety plan clinician", response.data)
                        history.push({
                            pathname: '/safetyplan',
                            state: response.data
                        })
                        }, (error) => {
                            console.log("Error"+error)
                        }
                    );
        }
    }
    return (
        <>
        <Header header="List of clients"/>
        <br></br>
        <br></br>
        <br></br>
        <br></br>
        <div style={{textAlign:"center"}}>
            
            <h3>Find your clients</h3>
            <br></br>
            <Search
            searchQuery={searchQuery}
            setSearchQuery={setSearchQuery}
            />
            {
                console.log(searchQuery)
            }
            <div className="DivWithScroll">
            {filteredClients.map((client) => (
                    <Card id="my_div" style={{marginLeft:"auto",marginRight:"auto", marginTop:"3%", width: '18rem' }}>
                    <Card.Body>
                        <Card.Title>{client.FirstName + " " + client.LastName}</Card.Title>
                        <Card.Link onClick={() => viewSafetyPlan(client)}>View Safety Plan</Card.Link> 
                        <Card.Link onClick={() => viewProfile(client)}>View Profile</Card.Link>
                    </Card.Body>
                </Card>
                ))}
            </div>
        </div>
        </>
    )
}

export default withRouter(Users)
