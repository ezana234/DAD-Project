import React from 'react';
import {Card} from 'react-bootstrap';
import 'bootstrap/dist/css/bootstrap.min.css';
import Header from './Header';

function SafetyPlan() {
    return (
        <div style={{textAlign:"center"}}>
                <br></br>
                <h1>My Safety Plan</h1>
                <Card style={{marginLeft:"auto",marginRight:"auto", marginTop:"3%", width: '18rem' }}>
                    <Card.Body>
                        <Card.Title>Warning Signs</Card.Title>
                        <Card.Link href="#">Link</Card.Link>
                    </Card.Body>
                </Card>
                <Card style={{marginLeft:"auto",marginRight:"auto", marginTop:"3%", width: '18rem' }}>
                    <Card.Body>
                        <Card.Title>Internal Coping Stratergies</Card.Title>
                        <Card.Link href="#">Link</Card.Link>
                    </Card.Body>
                </Card>
                <Card style={{marginLeft:"auto",marginRight:"auto", marginTop:"3%", width: '18rem' }}>
                    <Card.Body>
                        <Card.Title>External Coping Stratergies</Card.Title>
                        <Card.Link href="#">Link</Card.Link>
                    </Card.Body>
                </Card>
                <Card style={{marginLeft:"auto",marginRight:"auto", marginTop:"3%", width: '18rem' }}>
                    <Card.Body>
                        <Card.Title>Emergency Contacts</Card.Title>
                        <Card.Link href="#">Link</Card.Link>
                    </Card.Body>
                </Card>
                <Card style={{marginLeft:"auto",marginRight:"auto", marginTop:"3%", width: '18rem' }}>
                    <Card.Body>
                        <Card.Title>Request Immediate Help</Card.Title>
                        <Card.Link href="#">Link</Card.Link>
                    </Card.Body>
                </Card>
            </div>
    )
}

export default SafetyPlan
