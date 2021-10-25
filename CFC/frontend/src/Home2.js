import React from 'react';
import {Card} from 'react-bootstrap';
import 'bootstrap/dist/css/bootstrap.min.css';

function Home2() {
    return (
        <div style={{textAlign:"center"}}>
            <br></br>
            <h1>Users</h1>
            <Card style={{marginLeft:"auto",marginRight:"auto", marginTop:"3%", width: '18rem' }}>
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
            </Card>
        </div>
    )
}

export default Home2;
