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
                    if(response.status==200){
                        alert("Successfully deleted Appointment")
                        history.push({
                            pathname: '/clinicianHome',
                            state: {"Data": props.location.state.oldData, "Token": props.location.state.Token, "Role":props.location.state.Role, "oldData":props.location.state.oldData}
                        })
                    }
                    else{
                        alert("Error while deleting the Appointment, try again!")
                    }
                    }, (error) => {
                        alert("Error while deleting the Appointment, try again!")
                        console.log("Error"+error)
                    }
                );

    }

    function modifyAppointment(){
        const appdata = {"Appointment": props.location.state.Data[props.location.state.Index].AppointmentID, "AppointmentTime":props.location.state.Data[props.location.state.Index].AppointmentTime.replace("Z", "").slice(0, -3)
        , "AppointmentMedium":props.location.state.Data[props.location.state.Index].AppointmentMedium, "ClientName": props.location.state.ClientName,
        "ClinicianID": props.location.state.Data[props.location.state.Index].ClinicianID, "ClientID": props.location.state.Data[props.location.state.Index].ClientID
        }
        history.push({
            pathname: '/modifyAppointment',
            state: {"Data":appdata, "Token":props.location.state.Token, "AppointmentID":props.location.state.ID, "Role":props.location.state.Role, "oldData":props.location.state.oldData, "prev":props.location.state}
        })

    }

    const backClick = () =>{
        history.push({
          pathname: '/appointments',
          state: props.location.state.prev
      })
    }
    return (
        <>
        <Header header="Appointment Information" role={props.location.state.Role} oldData={props.location.state.oldData}/>
        <div style={{textAlign:"center", marginTop:"5rem"}}>
            
            <h3>Details</h3>
            <Card id="my_div" style={{marginLeft:"auto",marginRight:"auto", marginTop:"3%", width: '25rem' }}>
                    <Card.Body>
                    <Card.Title>{"Appointment "+props.location.state.Data[props.location.state.Index].AppointmentID}</Card.Title>
                    <p>{"Appointment Time: "+(props.location.state.Data[props.location.state.Index].AppointmentTime).replace("T", " ").replace("Z", "")}</p>
                    <p>{"Appointment Medium: "+props.location.state.Data[props.location.state.Index].AppointmentMedium}</p>
                    {props.location.state.Role == 2 &&
                    <p>{"Client's Name: "+props.location.state.ClientName}</p>}
                    <p>{"Clinician's Name: "+props.location.state.ClinicianName}</p>
                    {props.location.state.Role == 2 &&
                    <Card.Link style={{'cursor': 'pointer'}} onClick={() => deleteAppointment()}>Delete Appointment</Card.Link>
                    }
                    {props.location.state.Role == 2 &&
                    <Card.Link style={{'cursor': 'pointer'}} onClick={() => modifyAppointment()}>Modify Appointment</Card.Link>
                    }
                    <br></br>
                    <Card.Link style={{'cursor': 'pointer'}} onClick={backClick}>Go Back</Card.Link> 
                    </Card.Body>
            </Card>
        </div>
        </>
    )
}

export default AppointmentDetails
