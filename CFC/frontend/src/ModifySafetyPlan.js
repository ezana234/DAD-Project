import React, {useState} from 'react'
import './ModifySafetyPlan.css'
import Header from './Header';
import {withRouter, Link, useHistory } from 'react-router-dom';
import axios from 'axios';

function ModifySafetyPlan(props) {
    console.log("Modify safetyplan props:", props)
    const triggers_val = props.location.state.Data.Triggers;
    const warning_signs_val = props.location.state.Data.WarningSigns;
    const destructive_behaviors_val = props.location.state.Data.DestructiveBehaviors;
    const internal_statergies_val = props.location.state.Data.InternalStrategies;

    const history = useHistory();
    const [triggers, setTriggers] = useState(triggers_val);
    const [warningSigns, setWarningSigns] = useState(warning_signs_val);
    const [destrctiveBehaviors, setDestrctiveBehaviors] = useState(destructive_behaviors_val);
    const [internalStratergies, setInternalStratergies] = useState(internal_statergies_val);

    const modifySafetyPlan = event => {
        event.preventDefault();
        
        const safetyplan = {'SafetyID': props.location.state.Data.SafetyID,'Triggers': triggers, 'WarningSigns':warningSigns, 'DestructiveBehaviors':destrctiveBehaviors,'InternalStrategies':internalStratergies,'ClientID':props.location.state.Data.ClientID, 'ClinicianID': props.location.state.Data.ClinicianID};
        
        console.log(safetyplan)
        axios({ method: 'post', url: 'http://127.0.0.1:3000/clinician/safetyplan/update', headers: { 'Authorization': 'Bearer ' + props.location.state.Token } ,data: JSON.stringify(safetyplan)})
        .then((response) => {
            console.log(response);
            // alert(response.data)
            if(response.status==200){
                alert("Successfully modified Safety Plan")
            }
            else{
                alert("Could not modify the Safety Plan")
            }
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
                    <h1>Modify Safety Plan</h1>

                    <form>
                        <h5>Triggers:</h5>
                        <input type='text' value={triggers} onChange={event => setTriggers(event.target.value)} />

                        <h5>Warning Signs:</h5>
                        <input type='text' value={warningSigns} onChange={event => setWarningSigns(event.target.value)} />

                        <h5>Destructive Behaviors:</h5>
                        <input type='text' value={destrctiveBehaviors} onChange={event => setDestrctiveBehaviors(event.target.value)} />

                        <h5>Internal Coping Stratergies:</h5>
                        <input type='text' value={internalStratergies} onChange={event => setInternalStratergies(event.target.value)} />

                        <button type='submit' onClick={modifySafetyPlan} className='signInButton'>Modify Safety Plan</button>
                    </form>

                </div>
            </div>
        </>
    )
}

export default ModifySafetyPlan
