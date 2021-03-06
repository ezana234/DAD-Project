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
    const clients = props.location.state.Data;
    const filterClients = (clients, query) => {
        if (!query) {
            return clients;
        }
    
        return clients.filter((client) => {
            const clientName = client.FirstName;
            return clientName.startsWith(query);

        });

    };
    
    let query = "";
    const [searchQuery, setSearchQuery] = useState(query || '');
    const { search } = window.location;
    query = searchQuery;
    const filteredClients = filterClients(clients, query);

    function viewProfile(client){
        if(client!=null){
            history.push({
                pathname: '/profile',
                state: {"Data":{"Data":client}, "Token": props.location.state.Token, "Role":props.location.state.Role, "oldData": props.location.state.oldData, "prev":props.location.state, "w":1}
            })
        }
    }

    function viewSafetyPlan(client){
        if(client!=null){
            axios({ method: 'get', url: 'http://127.0.0.1:3000/clinician/safetyplan', headers: { 'Authorization': 'Bearer ' + props.location.state.Token }, params: {"userID": client.UserID} })
            .then((response) => {
                        console.log("Safety plan clinician", response.data)
                        history.push({
                            pathname: '/safetyplan',
                            state: {"Data": response.data, "Token": props.location.state.Token, "UserID":client.UserID, "Role":props.location.state.Role, "oldData": props.location.state.oldData, "Clients":clients, "prev":props.location.state, "s":1}
                        })
                        }, (error) => {
                            console.log("Error"+error)
                        }
                    );
        }
    }

    const backClick = () =>{
          history.push({
            pathname: '/clinicianHome',
            state: props.location.state.prev
        })
      }

    return (
        <>
        <Header header="List of Clients" role={props.location.state.Role} oldData={props.location.state.oldData}/>
        <div style={{textAlign:"center", marginTop:"5rem"}}>
            
            <h3>Find your Clients</h3>
            <br></br>
            <Search
            searchQuery={searchQuery}
            setSearchQuery={setSearchQuery}
            />
            <div className="DivWithScroll">
            {filteredClients.map((client, index) => (
                    <Card key={index} id="my_div" style={{marginLeft:"auto",marginRight:"auto", marginTop:"3%", width: '18rem', height: '9rem'}}>
                    <Card.Body>
                        <Card.Title>{client.FirstName + " " + client.LastName}</Card.Title>
                        <Card.Link style={{'cursor': 'pointer'}} onClick={() => viewSafetyPlan(client)}>View Safety Plan</Card.Link> 
                        <Card.Link style={{'cursor': 'pointer'}} onClick={() => viewProfile(client)}>View Profile</Card.Link>
                    </Card.Body>
                </Card>
                ))}
            </div>
            <br></br>
                    <Card.Link style={{'cursor': 'pointer'}} onClick={backClick}>Go back</Card.Link> 
        </div>
        </>
    )
}

export default withRouter(Users)
