import React from 'react'
import {Card} from 'react-bootstrap';
import 'bootstrap/dist/css/bootstrap.min.css';
import Header from './Header';
import Search from './Search';
import {withRouter, Link, useHistory } from 'react-router-dom';
import './Profile.css'

function Profile(props) {
    const history = useHistory();
    let firstname = "";
    let lastname = "";
    let email = "";
    let username = "";
    let address = "";
    let phonenum = "";
    if(props.location.state===undefined){
    }
    else{
        firstname = props.location.state.Data.Data.FirstName;
        lastname = props.location.state.Data.Data.LastName;
        email = props.location.state.Data.Data.Email;
        username = props.location.state.Data.Data.UserName;
        address = props.location.state.Data.Data.Address;
        phonenum = props.location.state.Data.Data.PhoneNumber;
    }
    const backClick = () =>{
        if(props.location.state.Role==1){
          history.push({
            pathname: '/clientHome',
            state: props.location.state.prev
        })
        }
        else{
            if(props.location.state.w==0){
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
            <div>
            {props.location.state.Role==1 &&
            <Header header="Client's Profile" role={props.location.state.Role} oldData={props.location.state.oldData}/>
            }
            {props.location.state.Role==2 &&
            <Header header="Clinician's Profile" role={props.location.state.Role} oldData={props.location.state.oldData}/>
            }
            <div style={{textAlign:"center", marginTop:"5rem"}}>
                <div className="card">
                    <h5>Profile Information</h5>
                    <br></br>
                    <h6>First Name: {firstname}</h6>
                    <h6>Last Name: {lastname}</h6>
                    <h6>Email: {email}</h6>
                    <h6>Username: {username}</h6>
                    <h6>Address: {address}</h6>
                    <h6>Phone number: {phonenum}</h6>
                    <br></br>
                    <Card.Link style={{'cursor': 'pointer'}} onClick={backClick}>Go back</Card.Link> 
                </div>
            </div>
            </div>
    )
}

export default Profile
