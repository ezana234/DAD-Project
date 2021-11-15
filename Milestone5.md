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

