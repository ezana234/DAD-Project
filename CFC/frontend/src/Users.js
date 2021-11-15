import React, {useState} from 'react'
import {Card} from 'react-bootstrap';
import 'bootstrap/dist/css/bootstrap.min.css';
import Header from './Header';
import Search from './Search';

function Users(props) {
    console.log("Users", props)
    const clients = props.location.state;
    const filterClients = (clients, query) => {
        if (!query) {
            return clients;
        }
    
        return clients.filter((client) => {
            const clientName = client.FirstName.toLowerCase();
            console.log(clientName);
            return clientName.includes(query);
        });
    };
    const { search } = window.location;
    const query = new URLSearchParams(search).get('s');
    const filteredClients = filterClients(clients, query);
    const [searchQuery, setSearchQuery] = useState(query || '');
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
            {/* <Card style={{marginLeft:"auto",marginRight:"auto", marginTop:"3%", width: '18rem' }}>
                <Card.Body>
                    <Card.Title>Steve Rogers</Card.Title>
                    <Card.Link href="#">Safety Plan</Card.Link>
                </Card.Body>
            </Card>
            <Card style={{marginLeft:"auto",marginRight:"auto", marginTop:"3%", width: '18rem' }}>
                <Card.Body>
                    <Card.Title>Tony Stark</Card.Title>
                    <Card.Link href="#">Safety Plan</Card.Link>
                </Card.Body>
            </Card>
            <Card style={{marginLeft:"auto",marginRight:"auto", marginTop:"3%", width: '18rem' }}>
                <Card.Body>
                    <Card.Title>Bruce Banner</Card.Title>
                    <Card.Link href="#">Safety Plan</Card.Link>
                </Card.Body>
            </Card>
            <Card style={{marginLeft:"auto",marginRight:"auto", marginTop:"3%", width: '18rem' }}>
                <Card.Body>
                    <Card.Title>Peter Parker</Card.Title>
                    <Card.Link href="#">Safety Plan</Card.Link>
                </Card.Body>
            </Card>
            <Card style={{marginLeft:"auto",marginRight:"auto", marginTop:"3%", width: '18rem' }}>
                <Card.Body>
                    <Card.Title>Clint Barton</Card.Title>
                    <Card.Link href="#">Safety Plan</Card.Link>
                </Card.Body>
            </Card> */}
        </div>
        </>
    )
}

export default Users
