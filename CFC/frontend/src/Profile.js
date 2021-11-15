import React from 'react'
import Header from './Header';

function Profile(props) {
    
    console.log(props)
    return (
            <div>
            <Header header="Client's Profile"/>
            <div style={{textAlign:"center"}}>
            <div className="card">
                <br></br>
                <br></br>
                <br></br>
                <br></br>
                <h5>My Profile</h5>
                <br></br>
                <h6>Fisrt Name: {props.location.state.FirstName}</h6>
                <h6>Last Name: {props.location.state.LastName}</h6>
                <h6>Email: {props.location.state.Email}</h6>
                <h6>Username: {props.location.state.UserName}</h6>
                <h6>Address: {props.location.state.Address}</h6>
                <h6>Phone number: {props.location.state.PhoneNumber}</h6>
            </div>
            </div>
            </div>
    )
}

export default Profile
