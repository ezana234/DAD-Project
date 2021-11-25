import React from 'react';
import {Card} from 'react-bootstrap';
import 'bootstrap/dist/css/bootstrap.min.css';
import Header from './Header';

function SafetyPlan(props) {
    console.log("Safety plans props", props)
    let safetyPlanObject = {}
    console.log(props.location.state==null)
    if(props.location.state == null){
        safetyPlanObject =  {"Triggers":"", "WarningSigns": "", "DestructiveBehaviors":"", "InternalStratergies":""}
    }
    else{
        safetyPlanObject = props.location.state[0];
        console.log(safetyPlanObject)
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
                </div>
            </>
    )
}

export default SafetyPlan
