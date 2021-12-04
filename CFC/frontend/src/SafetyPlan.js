import React from 'react';
import {Card} from 'react-bootstrap';
import 'bootstrap/dist/css/bootstrap.min.css';
import Header from './Header';
import axios from 'axios';
import {withRouter, Link, useHistory } from 'react-router-dom';

function SafetyPlan(props) {
    const history = useHistory();
    console.log("Safety plans props", props)
    let data = {};
    let token = props.location.state.Token;
    let id = props.location.state.UserID;
    let safetyPlanObject = {}
    console.log(props.location.state==null)
    if(props.location.state.Data == null){
        safetyPlanObject =  {"":"", "": "", "":"", "":""}
    }
    else{
        safetyPlanObject = props.location.state.Data[0];
        console.log(safetyPlanObject)
        data = props.location.state.Data[0];
        
        console.log("Values", data, token);
    }

    function modifySafetyPlan(props){
        history.push({
                pathname: '/modifysafetyplan',
                state: {"Data":data, "Token":token}
            })

    }

    function deleteSafetyPlan(){
        axios({ method: 'post', url: 'http://127.0.0.1:3000/clinician/safetyplan/delete', headers: { 'Authorization': 'Bearer ' + props.location.state.Token }, data: JSON.stringify({'SafetyID': data.SafetyID})})
        .then((response) => {
                    console.log("Delete plan clinician", response)
                    }, (error) => {
                        console.log("Error"+error)
                    }
                );

    }

    function addSafetyPlan(){
        history.push({
            pathname: '/addsafetyplan',
            state: {"Data":data, "Token":token, "UserID":id}
        })

    }

    return (
        <>
            <Header header="Safety Plan"/>
            <br></br>
            <br></br>
            <br></br>
            <div style={{textAlign:"center"}}>
                    <br></br>
                    <h4>Safety Plan Details:</h4>
                    {Object.keys(safetyPlanObject).map((key) => (
                        (key=="Triggers" || key == "WarningSigns" || key == "DestructiveBehaviors" || key == "InternalStrategies") &&
                            <Card style={{marginLeft:"auto",marginRight:"auto", marginTop:"3%", width: '18rem' }}>
                            <Card.Body>
                                <Card.Title>{key}</Card.Title>
                                <p>{safetyPlanObject[key]}</p>
                                
                            </Card.Body>
                            </Card>
                    ))
                    }
                    {
                         props.location.state.Data==null&& <Card.Link onClick={() => addSafetyPlan()}>Add Safety Plan</Card.Link>
                    }
                    {
                         props.location.state.Data!=null&& <Card.Link onClick={() => modifySafetyPlan()}>Modify Safety Plan</Card.Link> 
                    }
                    {
                         props.location.state.Data!=null&& <Card.Link onClick={() => deleteSafetyPlan()}>Delete Safety Plan</Card.Link> 
                    }
                    
                </div>
            </>
    )
}

export default SafetyPlan
