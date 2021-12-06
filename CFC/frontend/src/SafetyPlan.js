import React from 'react';
import {Card} from 'react-bootstrap';
import 'bootstrap/dist/css/bootstrap.min.css';
import Header from './Header';
import axios from 'axios';
import {withRouter, Link, useHistory } from 'react-router-dom';

function SafetyPlan(props) {
    let clients = props.location.state.Clients;
    let role = props.location.state.Role;
    let olddata = props.location.state.oldData;
    let prevdata = props.location.state
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
                state: {"Data":data, "Token":token, "Role":role, "oldData": olddata, "Clients":clients, "prev":prevdata}
            })

    }

    function deleteSafetyPlan(){
        axios({ method: 'post', url: 'http://127.0.0.1:3000/clinician/safetyplan/delete', headers: { 'Authorization': 'Bearer ' + props.location.state.Token }, data: JSON.stringify({'SafetyID': data.SafetyID})})
        .then((response) => {
                    console.log("Delete plan clinician", response)
                    if(response.status==200){
                        alert("Successfully deleted Safety Plan")
                        history.push({
                            pathname: '/users',
                            state: {"Data": props.location.state.Clients, "Token": props.location.state.Token, "Role":props.location.state.Role, "oldData":props.location.state.oldData}
                        })
                    }
                    else{
                        alert("Error while deleting the Safety Plan, try again!")
                    }
                    }, (error) => {
                        console.log("Error"+error)
                        alert("Error while deleting the Safety Plan, try again!")

                    }
                );

    }

    function addSafetyPlan(){
        history.push({
            pathname: '/addsafetyplan',
            state: {"Data":data, "Token":token, "UserID":id, "Role":props.location.state.Role, "oldData": props.location.state.oldData, "Clients":props.location.state.Clients, "prev":props.location.state}
        })

    }

    const backClick = () =>{
        if(props.location.state.Role==1){
          history.push({
            pathname: '/clientHome',
            state: props.location.state.prev
        })
        }
        else{
            if(props.location.state.s==0){
                history.push({
                    pathname: '/clinicianHome',
                    state: props.location.state.prev
            
                })
            }
            else{
                history.push({
                    pathname: '/users',
                    state: props.location.state.prev
            
                })
            }
        }
      }

    return (
        <>
            <Header header="Safety Plan" role={props.location.state.Role} oldData={props.location.state.oldData}/>

            <div style={{textAlign:"center", marginTop:"5rem"}}>
                    <br></br>
                    <h4>Safety Plan Details:</h4>
                    {Object.keys(safetyPlanObject).map((key) => (
                        (key=="Triggers" || key == "WarningSigns" || key == "DestructiveBehaviors" || key == "InternalStrategies") &&
                            <Card style={{marginLeft:"auto",marginRight:"auto", marginTop:"3%", width: '18rem' }}>
                            <Card.Body>
                                {key == "WarningSigns" && <Card.Title>Warning Signs</Card.Title>}
                                {key == "Triggers" && <Card.Title>Triggers</Card.Title>}
                                {key == "DestructiveBehaviors" && <Card.Title>Destructive Behaviors</Card.Title>}
                                {key == "InternalStrategies" && <Card.Title>Internal Strategies</Card.Title>}

                                <p>{safetyPlanObject[key]}</p>
                                
                            </Card.Body>
                            </Card>
                    ))
                    }
                    <br></br>
                    {
                         props.location.state.Data==null&& props.location.state.Role==2 && <Card.Link style={{'cursor': 'pointer'}} onClick={() => addSafetyPlan()}>Add Safety Plan</Card.Link>
                    }
                    {
                         props.location.state.Data!=null&& props.location.state.Role==2 && <Card.Link style={{'cursor': 'pointer'}} onClick={() => modifySafetyPlan()}>Modify Safety Plan</Card.Link> 
                    }
                    {
                         props.location.state.Data!=null&& props.location.state.Role==2 && <Card.Link style={{'cursor': 'pointer'}} onClick={() => deleteSafetyPlan()}>Delete Safety Plan</Card.Link> 
                    }
                    <br></br>
                    <Card.Link style={{'cursor': 'pointer'}} onClick={backClick}>Go back</Card.Link> 
                </div>
            </>
    )
}

export default SafetyPlan
