import React, {useState} from 'react'
import Header from './Header';
import {withRouter, Link, useHistory } from 'react-router-dom';
import axios from 'axios';
import './ModifyAppointment.css'
import {Card} from 'react-bootstrap';
import 'bootstrap/dist/css/bootstrap.min.css';

function ModifyAppointment(props) {
    console.log("Modify Appointment props:", props)
    const clients = props.location.state.Data;
    const history = useHistory();
    const [time, setTime] = useState(props.location.state.Data.AppointmentTime);
    const [medium, setMedium] = useState(props.location.state.Data.AppointmentMedium);

    const modifyAppointment = event => {
        event.preventDefault();
        console.log("Click")
        console.log(time, medium);
        const appointment = { 'AppointmentTime': time.replace("T", " ")+":00", 'AppointmentMedium':medium,'ClientID':props.location.state.Data.ClientID, 'ClinicianID':props.location.state.Data.ClinicianID, 'AppointmentID': props.location.state.AppointmentID};
        
            console.log(appointment)
            axios({ method: 'post', url: 'http://127.0.0.1:3000/appointment/update', headers: { 'Authorization': 'Bearer ' + props.location.state.Token } ,data: JSON.stringify(appointment)})
            .then((response) => {
                console.log(response);
                // alert(response.data)
                if(response.status==200){
                    alert("Successfully modified Appointment")
                    history.push({
                        pathname: '/clinicianHome',
                        state: {"Data": props.location.state.oldData, "Token": props.location.state.Token, "Role":props.location.state.Role, "oldData":props.location.state.oldData}
                    })
                }
                else{
                    alert("Error, could not modify the Appointment! Try again")
                }
                }, (error) => {
                    console.log("Error"+error)
                    alert("Error, could not modify the Appointment! Try again")

                }
            );

            
    }
    const backClick = () =>{
        history.push({
          pathname: '/appointmentDetails',
          state: props.location.state.prev
      })
    }
    return (
        <>
            <Header header="Appointment" role={props.location.state.Role} oldData={props.location.state.oldData}/>
            <div style={{marginTop:"5rem"}} className='loginForm'>

                <div className='container'>
                    <h1>Modify Appointment</h1>

                    <form>
                        <h5 style={{paddingBottom:"5px"}}>Appointment {props.location.state.Data.Appointment}</h5>
                        <h5 style={{paddingBottom:"5px"}}>Client Name: {props.location.state.Data.ClientName}</h5>
                        <h5>Appointment Time:</h5>
                        <input type='datetime-local' id="appointmenttime" value={time} onChange={event => setTime(event.target.value)} />
                        {/* <input type='text' value={time} onChange={event => setTime(event.target.value)} /> */}

                        <h5>Appointment Medium:</h5>
                        <input type='text' value={medium} onChange={event => setMedium(event.target.value)} />

                        <button type='submit' onClick={modifyAppointment} className='signInButton'>Modify Appointment</button>
                    </form>

                </div>
                <br></br>
                <Card.Link style={{'cursor': 'pointer'}} onClick={backClick}>Go Back</Card.Link> 
            </div>
        </>
    )
}

export default ModifyAppointment
