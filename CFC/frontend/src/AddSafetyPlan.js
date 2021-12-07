import React, {useState} from 'react'
import './AddSafetyPlan.css'
import Header from './Header';
import {withRouter, Link, useHistory } from 'react-router-dom';
import axios from 'axios';
import {Card} from 'react-bootstrap';
import 'bootstrap/dist/css/bootstrap.min.css';



function AddSafetyPlan(props) {
    const history = useHistory();
    const [triggers, setTriggers] = useState('');
    const [warningSigns, setWarningSigns] = useState('');
    const [destrctiveBehaviors, setDestrctiveBehaviors] = useState('');
    const [internalStratergies, setInternalStratergies] = useState('');

    const createSafetyPlan = event => {
        event.preventDefault();
        axios({ method: 'get', url: 'http://127.0.0.1:3000/client/userid', headers: { 'Authorization': 'Bearer ' + props.location.state.Token }, params: {"userID": props.location.state.UserID} })
        .then((response) => {
            let iD = response.data.ClientID;
            const safetyplan = { 'Triggers': triggers, 'WarningSigns':warningSigns, 'DestructiveBehaviors':destrctiveBehaviors,'InternalStrategies':internalStratergies,'ClientID':iD};
            axios({ method: 'post', url: 'http://127.0.0.1:3000/clinician/safetyplan/add', headers: { 'Authorization': 'Bearer ' + props.location.state.Token } ,data: JSON.stringify(safetyplan)})
            .then((response) => {
                if(response.status==200){
                    alert("Safety Plan successfully added")
                    history.push({
                        pathname: '/users',
                        state: {"Data": props.location.state.Clients, "Token": props.location.state.Token, "Role":props.location.state.Role, "oldData":props.location.state.oldData}
                    })
                }
                else{
                    alert("Error while adding Safety Plan, please try again!")
                }
                }, (error) => {
                    alert("Error while adding Safety Plan, please try again!")
                    console.log("Error"+error)
                }
            );
            }, (error) => {
                console.log("Error"+error)
            }
        );
        

            
    }

    const backClick = () =>{
          history.push({
            pathname: '/safetyplan',
            state: props.location.state.prev
        })
      }

    return (
        <>
            <Header header="New Safety Plan" role={props.location.state.Role} oldData={props.location.state.oldData}/>
            <div style={{marginTop:"5rem"}} className='loginForm'>

                <div className='container'>
                    <h1>Add Safety Plan</h1>

                    <form>
                        <h5>Triggers:</h5>
                        <input type='text' value={triggers} onChange={event => setTriggers(event.target.value)} />

                        <h5>Warning Signs:</h5>
                        <input type='text' value={warningSigns} onChange={event => setWarningSigns(event.target.value)} />

                        <h5>Destructive Behaviors:</h5>
                        <input type='text' value={destrctiveBehaviors} onChange={event => setDestrctiveBehaviors(event.target.value)} />

                        <h5>Internal Coping Strategies:</h5>
                        <input type='text' value={internalStratergies} onChange={event => setInternalStratergies(event.target.value)} />

                        <button type='submit' onClick={createSafetyPlan} className='signInButton'>Create Safety Plan</button>
                    </form>

                </div>
                <br></br>
                    <Card.Link style={{'cursor': 'pointer'}} onClick={backClick}>Go Back</Card.Link> 
            </div>
        </>
    )
}

export default AddSafetyPlan
