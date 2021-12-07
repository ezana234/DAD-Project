import React, {useState} from 'react'
import {Card} from 'react-bootstrap';
import 'bootstrap/dist/css/bootstrap.min.css';
import Header from './Header'
import {withRouter, Link, useHistory } from 'react-router-dom';
import axios from 'axios';
import './ViewAppointments.css'

function ViewAppointments(props) {
    let clinicianID = 0
    let arr=[]    

    const history = useHistory();
    if(props.location.state.Data!=null){
        clinicianID = props.location.state.Data[0].ClinicianID
        arr = props.location.state.Data;
    }

    function viewAppointmentDetails(appointmentId, index){
        if(appointmentId!=null){
            let clientname = ""
            let clinicianname = ""
            axios({ method: 'get', url: 'http://127.0.0.1:3000/appointment', headers: { 'Authorization': 'Bearer ' + props.location.state.Token }, params: {"appointmentID": appointmentId} })
            .then((response) => {
                        if(response.status==200){
                            axios({ method: 'get', url: 'http://127.0.0.1:3000/clientname', headers: { 'Authorization': 'Bearer ' + props.location.state.Token }, params: {"clientID": response.data.ClientID} })
                            .then((res) => {
                                        if(res.status==200){
                                            clientname = res.data.FirstName+" "+res.data.LastName;
                                            axios({ method: 'get', url: 'http://127.0.0.1:3000/clinicianname', headers: { 'Authorization': 'Bearer ' + props.location.state.Token }, params: {"clinicianID": response.data.ClinicianID} })
                                            .then((re) => {
                                                        if(re.status==200){
                                                            clinicianname = re.data.FirstName+" "+re.data.LastName
                                                            history.push({
                                                                pathname: '/appointmentDetails',
                                                                state: {"Data": arr, "Role":props.location.state.Role, "Token":props.location.state.Token, "ID":appointmentId, "ClientName":clientname, "ClinicianName": clinicianname, "Index":index, "oldData": props.location.state.oldData, "prev":props.location.state}
                                                            })
                                                        }
                                                        }, (error) => {
                                                            console.log("Error"+error)
                                                        }
                                                    );
                                            
                                        }
                                        }, (error) => {
                                            console.log("Error"+error)
                                        }
                                    );
                            
                        }
                        }, (error) => {
                            console.log("Error"+error)
                        }
                    );
        }
        

    }

    function addAppointment(){
        axios({ method: 'get', url: 'http://127.0.0.1:3000/clinician/clients', headers: { 'Authorization': 'Bearer ' + props.location.state.Token } })
                    .then((response) => {
                    if(response.status  == 200){
                        history.push({
                            pathname: '/addAppointment',
                            state: {"Data": response.data, "Token": props.location.state.Token, "Role":props.location.state.Role, "ClinicianID": clinicianID, "oldData": props.location.state.oldData, "prev":props.location.state}
                        })
                    }
                }, (error) => {
                    console.log(error)
                }
            );
    }

    const backClick = () =>{
        if(props.location.state.Role==1){
          history.push({
            pathname: '/clientHome',
            state: props.location.state.prev
        })
        }
        else{
          history.push({
            pathname: '/clinicianHome',
            state: props.location.state.prev
    
        })
        }
      }

    return (
        <>
        <Header header="List of Appointments" role={props.location.state.Role} oldData={props.location.state.oldData}/>
        <div style={{textAlign:"center", marginTop:"5rem"}}>
            
            <h3>Your Appointments</h3>
            <br></br>
            <div className="DivWithScroll">
            {arr.map((appointment, index) => (
                    <Card key={index} id="my_div" style={{marginLeft:"auto",marginRight:"auto", marginTop:"3%", width: '18rem', height: '40%'}}>
                    <Card.Body>
                    <Card.Title>{"Appointment "+appointment.AppointmentID}</Card.Title>
                    <p>{"Time: "+(appointment.AppointmentTime.replace("T", " ").replace("Z", ""))}</p>
                    <Card.Link style={{'cursor': 'pointer'}} onClick={() => viewAppointmentDetails(appointment.AppointmentID, index)}>View Details</Card.Link>
                    </Card.Body>
                </Card>
                ))}
            </div>
            {props.location.state.Role == 2 &&
            <Card.Link style={{'cursor': 'pointer'}} onClick={() => addAppointment()}>Add Appointment</Card.Link>}
            <br></br>
            <Card.Link style={{'cursor': 'pointer'}} onClick={backClick}>Go back</Card.Link> 
        </div>
        </>
    )
}

export default ViewAppointments
