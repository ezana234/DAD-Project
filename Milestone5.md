# Milestone 5

## Project Updates

### Backend
- Changes to Database
    1. The Person table now has fields for DOB and password expiration
    2. The Clinician table now has a field for referral code, which will be used byt he client when they go to create an account on the site. The referral code is used to prevent spam accounts and to automatically assign that clinician to the client.
- Updated routes for login and for clinicians to get clients
    1. JWT Authentication tokens implemented for users
- Updated facades and DAOs for more functionality
    1. Mostly just adding the ability to get data from multiple tables within a facade (example: Able to get Safety Plans from the ClientFacade via clientID)

### Frontend
- The frontend allows only authorized users to sign in
- Modified Client's and Clinician's homepage to list all the options a user can do upon landing on the home page, also included a navbar.
- Clients can now view their profile by clicking on the View My Profile button on their homepage
- Clinicians can view/search for users by clicking on View Users button on their homepage
- Refactored the entire frontend code to separate View components and the actual components which contain the data. Following this stratergy made the code look cleaner, readable and easily understandable. Earlier two other components were grouped with the View component which made the code look very messy and cluttered. In the code snippet below, we can see one exmaple of how the code has been restructured.

```
return (
        <>
            <Header header="Client's Homepage"/>
            <div style={{textAlign:"center", marginLeft:"auto", marginRight:"auto"}}>
                <h4>Welcome {firstname}</h4>
                <br></br>
                <h5>What would you like to do today?</h5>
                <button onClick={viewProfile} class="myButton">View my profile</button>
                <br></br>
                <button class="myButton">View my safety plan</button>
                <br></br>
                <button class="myButton">Add emergency contact</button>
            </div>
            
        </>
    )
```

This is the code used to render Client's homepage. Notice how the Navbar/Header is split into a component of it's own and being used(in the second line inside return method) in this View component. A similar code structure will be noticeable in all the other components as well.
