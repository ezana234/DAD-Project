import React from 'react'
import Header from './Header';

function Profile(props) {
    
    console.log("Profile", props)
    let firstname = "";
    let lastname = "";
    let email = "";
    let username = "";
    let address = "";
    let phonenum = "";
    if(props.location.state===undefined){
        console.log("Undefined")
    }
    else{
        console.log("Not undefined")
        firstname = props.location.state.Data.FirstName;
        lastname = props.location.state.Data.LastName;
        email = props.location.state.Data.Email;
        username = props.location.state.Data.UserName;
        address = props.location.state.Data.Address;
        phonenum = props.location.state.Data.PhoneNumber;
    }
    return (
            <div>
            <Header header="Client's Profile"/>
            <div style={{textAlign:"center"}}>
            <div className="card">
                <br></br>
                <br></br>
                <br></br>
                <br></br>
                <h5>Profile Info</h5>
                <br></br>
                <h6>Fisrt Name: {firstname}</h6>
                <h6>Last Name: {lastname}</h6>
                <h6>Email: {email}</h6>
                <h6>Username: {username}</h6>
                <h6>Address: {address}</h6>
                <h6>Phone number: {phonenum}</h6>
            </div>
            </div>
            </div>
    )
}

export default Profile
