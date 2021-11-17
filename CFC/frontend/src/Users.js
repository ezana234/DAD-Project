import React, {useState} from 'react'
import {Card} from 'react-bootstrap';
import 'bootstrap/dist/css/bootstrap.min.css';
import Header from './Header';
import Search from './Search';

function Users(props) {
    console.log("Users", props)
    const clients = props.location.state;
    console.log("Type: ", typeof(clients))
    const filterClients = (clients, query) => {
        console.log("Inside filter clients method")
        console.log(!query)
        if (!query) {
            return clients;
        }
    
        return clients.filter((client) => {
            const clientName = client.FirstName.toLowerCase();
            console.log(clientName);
            return clientName.includes(query);
        });
    };
    
    let query = "";
    const [searchQuery, setSearchQuery] = useState(query || '');
    const { search } = window.location;
    query = searchQuery;
    console.log(query)
    const filteredClients = filterClients(clients, query);
    console.log(searchQuery);
    return (
        <>
        <Header header="Clients"/>
        <br></br>
        <br></br>
        <br></br>
        <br></br>
        <div style={{textAlign:"center"}}>
            
            <h3>Clients</h3>
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
                    <Card style={{marginLeft:"auto",marginRight:"auto", marginTop:"3%", width: '18rem' }}>
                    <Card.Body>
                        <Card.Title>{client.FirstName + client.LastName}</Card.Title>
                        <Card.Link href="#">Safety Plan</Card.Link>
                        <Card.Link href="#">Profile</Card.Link>
                    </Card.Body>
                </Card>
                ))}
            </div>
        </div>
        </>
    )
}

export default Users
