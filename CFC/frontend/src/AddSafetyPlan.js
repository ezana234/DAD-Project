import React, {useState} from 'react'
import './AddSafetyPlan.css'
import Header from './Header';
import {withRouter, Link, useHistory } from 'react-router-dom';
import axios from 'axios';



function AddSafetyPlan(props) {
    console.log("Add safetyplan props:", props)
    const history = useHistory();
    const [triggers, setTriggers] = useState('');
    const [warningSigns, setWarningSigns] = useState('');
    const [destrctiveBehaviors, setDestrctiveBehaviors] = useState('');
    const [internalStratergies, setInternalStratergies] = useState('');

    const createSafetyPlan = event => {
        event.preventDefault();
        console.log("Click")
        axios({ method: 'get', url: 'http://127.0.0.1:3000/client/userid', headers: { 'Authorization': 'Bearer ' + props.location.state.Token }, params: {"userID": props.location.state.UserID} })
        .then((response) => {
            console.log("res")
            let iD = response.data.ClientID;
            const safetyplan = { 'Triggers': triggers, 'WarningSigns':warningSigns, 'DestructiveBehaviors':destrctiveBehaviors,'InternalStrategies':internalStratergies,'ClientID':iD};
        
            console.log(safetyplan)
            axios({ method: 'post', url: 'http://127.0.0.1:3000/clinician/safetyplan/add', headers: { 'Authorization': 'Bearer ' + props.location.state.Token } ,data: JSON.stringify(safetyplan)})
            .then((response) => {
                console.log(response);
                // alert(response.data)
                }, (error) => {
                    console.log("Error"+error)
                }
            );
            }, (error) => {
                console.log("Error"+error)
            }
        );
        

            
    }

    return (
        <>
            <Header/>
            <br></br>
            <br></br>
            <br></br>
            <div className='loginForm'>

                <div className='container'>
                    <h1>Add Safety Plan</h1>

                    <form>
                        <h5>Triggers:</h5>
                        <input type='text' value={triggers} onChange={event => setTriggers(event.target.value)} />

                        <h5>Warning Signs:</h5>
                        <input type='text' value={warningSigns} onChange={event => setWarningSigns(event.target.value)} />

                        <h5>Destructive Behaviors:</h5>
                        <input type='text' value={destrctiveBehaviors} onChange={event => setDestrctiveBehaviors(event.target.value)} />

                        <h5>Internal Coping Stratergies:</h5>
                        <input type='text' value={internalStratergies} onChange={event => setInternalStratergies(event.target.value)} />

                        <button type='submit' onClick={createSafetyPlan} className='signInButton'>Create Safety Plan</button>
                    </form>

                </div>
            </div>
        </>
    )
}

export default AddSafetyPlan
