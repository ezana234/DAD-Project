import React, {useState} from 'react'
import {Card} from 'react-bootstrap';
import 'bootstrap/dist/css/bootstrap.min.css';
import Header from './Header'
import {withRouter, Link, useHistory } from 'react-router-dom';
import axios from 'axios';
import './ViewAppointments.css'

function ViewAppointments(props) {
    const history = useHistory();
    let arr = []
    arr = props.location.state.Data;

    function viewAppointmentDetails(appointmentId, index){
        console.log("Inside VIEW Details")
        console.log(appointmentId);
        // console.log("Index: "+clients.indexOf(client))
        if(appointmentId!=null){
            let clientname = ""
            let clinicianname = ""
            axios({ method: 'get', url: 'http://127.0.0.1:3000/appointment', headers: { 'Authorization': 'Bearer ' + props.location.state.Token }, params: {"appointmentID": appointmentId} })
            .then((response) => {
                        if(response.status==200){
                            console.log("Appointment details", response.data)
                            axios({ method: 'get', url: 'http://127.0.0.1:3000/clientname', headers: { 'Authorization': 'Bearer ' + props.location.state.Token }, params: {"clientID": response.data.ClientID} })
                            .then((res) => {
                                        if(res.status==200){
                                            console.log("Client Name details", res.data)
                                            clientname = res.data.FirstName+" "+res.data.LastName;
                                        }
                                        }, (error) => {
                                            console.log("Error"+error)
                                        }
                                    );
                            axios({ method: 'get', url: 'http://127.0.0.1:3000/clinicianname', headers: { 'Authorization': 'Bearer ' + props.location.state.Token }, params: {"clinicianID": response.data.ClinicianID} })
                                .then((re) => {
                                            if(re.status==200){
                                                console.log("Clinician name details", re.data)
                                                clinicianname = re.data.FirstName+" "+re.data.LastName
                                            }
                                            }, (error) => {
                                                console.log("Error"+error)
                                            }
                                        );
                            history.push({
                                pathname: '/appointmentDetails',
                                state: {"Data": arr, "Token":props.location.state.Token, "ID":appointmentId, "ClientName":clientname, "ClinicianName": clinicianname, "Index":index}
                            })
                        }
                        }, (error) => {
                            console.log("Error"+error)
                        }
                    );
        }
        

    }

    return (
        <>
        <Header header="List of appointments"/>
        <br></br>
        <br></br>
        <br></br>
        <br></br>
        <div style={{textAlign:"center"}}>
            
            <h3>Your appointments</h3>
            <br></br>
            <div className="DivWithScroll">
            {arr.map((appointment, index) => (
                    <Card id="my_div" style={{marginLeft:"auto",marginRight:"auto", marginTop:"3%", width: '18rem' }}>
                    <Card.Body>
                    <Card.Title>{"Appointment "+appointment.AppointmentID}</Card.Title>
                    <p>{"Time: "+appointment.AppointmentTime}</p>
                    <p>{"Medium: "+appointment.AppointmentMedium}</p>
                    <Card.Link onClick={() => viewAppointmentDetails(appointment.AppointmentID, index)}>View Details</Card.Link>
                    {/* <Card.Link onClick={() => viewAppointmentDetails(appointment.AppointmentID)}>View Details</Card.Link> */}
                    </Card.Body>
                </Card>
                ))}
            </div>
        </div>
        </>
    )
}

export default ViewAppointments
