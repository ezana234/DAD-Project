import React, {useState} from 'react'
import {Card} from 'react-bootstrap';
import 'bootstrap/dist/css/bootstrap.min.css';
import Header from './Header'
import {withRouter, Link, useHistory } from 'react-router-dom';
import axios from 'axios';
import './AppointmentDetails.css'

function AppointmentDetails(props) {
    const history = useHistory();
    console.log(props);

    function deleteAppointment(){
        axios({ method: 'post', url: 'http://127.0.0.1:3000/appointment/delete', headers: { 'Authorization': 'Bearer ' + props.location.state.Token }, data: JSON.stringify({'appointmentID': props.location.state.ID})})
        .then((response) => {
                    console.log("Delete appointment", response)
                    }, (error) => {
                        console.log("Error"+error)
                    }
                );

    }

    return (
        <>
        <Header header="Appointment Details"/>
        <br></br>
        <br></br>
        <br></br>
        <br></br>
        <div style={{textAlign:"center"}}>
            
            <h3>Appointment Details</h3>
            <br></br>
            <Card id="my_div" style={{marginLeft:"auto",marginRight:"auto", marginTop:"3%", width: '18rem' }}>
                    <Card.Body>
                    <Card.Title>{"Appointment ID"+props.location.state.Data[props.location.state.Index].AppointmentID}</Card.Title>
                    <p>{"Time: "+props.location.state.Data[props.location.state.Index].AppointmentTime}</p>
                    <p>{"Medium: "+props.location.state.Data[props.location.state.Index].AppointmentMedium}</p>
                    <p>{"Client's Name: "+props.location.state.ClientName}</p>
                    <p>{"Clinician's Name: "+props.location.state.ClinicianName}</p>
                    <Card.Link onClick={() => deleteAppointment()}>Delete Appointment</Card.Link>
                    {/* <Card.Link onClick={() => modifyAppointment()}>Modify Appointment</Card.Link> */}
                    </Card.Body>
            </Card>
        </div>
        </>
    )
}

export default AppointmentDetails
