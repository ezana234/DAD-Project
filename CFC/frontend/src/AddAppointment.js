import React, {useState} from 'react'
import Header from './Header';
import {withRouter, Link, useHistory } from 'react-router-dom';
import axios from 'axios';
import './AddAppointment.css'
import {Card} from 'react-bootstrap';
import 'bootstrap/dist/css/bootstrap.min.css';

function AddAppointment(props) {
    console.log("Add Appointment props:", props)
    const clients = props.location.state.Data;
    const history = useHistory();
    const [id, setId] = useState(0);
    const [time, setTime] = useState('');
    const [medium, setMedium] = useState('');

    const createAppointment = event => {
        event.preventDefault();
        console.log("Click")
        console.log(id, time, medium);
        axios({ method: 'get', url: 'http://127.0.0.1:3000/client/userid', headers: { 'Authorization': 'Bearer ' + props.location.state.Token }, params: {"userID": id} })
        .then((response) => {
            console.log("res")
            let clientID = response.data.ClientID;
            const appointment = { 'AppointmentTime': time.replace("T", " ")+":00", 'AppointmentMedium':medium,'ClientID':clientID, 'ClinicianID':props.location.state.ClinicianID};
        
            console.log(appointment)
            axios({ method: 'post', url: 'http://127.0.0.1:3000/appointment/add', headers: { 'Authorization': 'Bearer ' + props.location.state.Token } ,data: JSON.stringify(appointment)})
            .then((response) => {
                console.log(response);
                // alert(response.data)
                if(response.status==200){
                    alert("Successfully added Appointment")
                    history.push({
                        pathname: '/clinicianHome',
                        state: {"Data": props.location.state.oldData, "Token": props.location.state.Token, "Role":props.location.state.Role, "oldData":props.location.state.oldData}
                    })
                }
                else{
                    alert("Error, could not add the Appointment! Try again")
                }
                }, (error) => {
                    console.log("Error"+error)
                    alert("Error, could not add the Appointment! Try again")

                }
            );
            }, (error) => {
                console.log("Error"+error)
            }
        );

            
    }
    function handleSelectChange(event) {
        console.log(event.target.value)
        setId(event.target.value);
    }
    const backClick = () =>{
        history.push({
          pathname: '/appointments',
          state: props.location.state.prev
      })
    }
    return (
        <>
            <Header header="New Appointment" role={props.location.state.Role} oldData={props.location.state.oldData}/>
            <div style={{marginTop:"5rem"}} className='loginForm'>

                <div className='container'>
                    <h1>Add Appointment</h1>

                    <form>
                        <h5>Client Name:</h5>
                        <select onChange={handleSelectChange}>
                            {clients.map(client => {
                            return (
                                <option value={client.UserID}> {client.FirstName + " " + client.LastName} </option>
                            )
                            })}
                        </select>
                        <br></br><br></br>
                        <h5>Appointment Time:</h5>
                        <input type='datetime-local' id="appointmenttime" value={time} onChange={event => setTime(event.target.value)} />
                        {/* <input type='text' value={time} onChange={event => setTime(event.target.value)} /> */}

                        <h5>Appointment Medium:</h5>
                        <input type='text' value={medium} onChange={event => setMedium(event.target.value)} />

                        <button type='submit' onClick={createAppointment} className='signInButton'>Create Appointment</button>
                    </form>

                </div>
                <br></br>
                <Card.Link style={{'cursor': 'pointer'}} onClick={backClick}>Go Back</Card.Link> 
            </div>
        </>
    )
}

export default AddAppointment
