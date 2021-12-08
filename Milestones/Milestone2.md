# Development Book

Group 1, ISTE.432.01

Fall 2021

## **Table of Contents**

**[Team Members and Roles:](#_weu71772umyy) 3**

**[Background:](#_vo2jif98s2qt) 3**

**[Project Description:](#_l1s5s8362ac0) 3**

**[Project Requirements:](#_7m8vliial958) 3**

**[Business Rules:](#_lqph00cwewet) 4**

**[Technologies Used:](#_40gpt1hy2irx) 4**

**[Design Pattern:](#_djspdxmuhigd) 4**

**[Entity Relationship Diagram:](#_f8024x3x3c8g) 4**

**[Timeline:](#_h6euul65kw8z) 4**

## Team Members and Roles:

- Siddhesh Periaswami, Front-End
- Ezana Kalekristos, Database Administrator
- Ryan Weiss, Back-End
- Afzal Ali, UI/UX Designer

## Background:

Catholic Family Center works to address the barriers preventing people from moving to self-sufficiency in today&#39;s environment. Programs are collaborative and focused on systemic change. Clinicians work with high risk clients to develop a safety plan on paper. That plan is printed and handed to the client or, in the case of phone or Zoom encounters, printed and mailed to the client.The plan, being printed, is not interactive and may easily be lost or destroyed or stored somewhere without ready access by the client.

## Project Description:

The current system involves printing the safety plan for clients on paper and then handing that plan to the clients. If the client is on the phone or over Zoom, then the safety plan is printed and mailed to the client. This is not ideal because it can be easily lost, destroyed, or not easily accessible to the client. In addition, the paper plan is not interactive and thus may not be as helpful as it could. The main goal of this project is to streamline the process of accessing and editing the safety plan using the clinician&#39;s desktop/phone and the client&#39;s phone.

## Project Requirements:

- Safe and secure storage of safety plans uploaded by the clinician.
- Clinicians should be able to copy/paste the plan into another application.
- The application should ask for access to the contact list to get a friend&#39;s phone number for calls
- The users should be presented with choices to create or modify safety plans, and see when a plan was edited and by who.
- The application should have search bar for the users to search for client&#39;s safety plans
- The app should allow both clinician and client to enter and delete items since clients may have trouble typing during phone calls.
- Add resources that the clients can access in the app (help lines, on call clinician)
- Signing up. logging in, and logging out
- The clinician should be able to add a client

## Business Rules:

- You must login before entering details in the application.
- You must be an authorized user of the system to access (Clinician, Client, Family of Client, etc.)
- When a user signs up for the first time, they should fill up their personal information before proceeding with any other steps.
- Only clients and clinicians can create appointments
- The professionals/agencies that we recommend must be verified to be credible before using with clients
- Only a clinician can invite another clinician to support a client
- The date, time and person who edited a safety plan must be recorded
- All users must be able to logout
- Clinicians must be able to to view all clients in the system
- Has to have full functionality on mobile and desktop

## Technologies Used:

**Front-End Technologies:** React, Javascript

**Back-End Technologies:** Go

**Database:** MySQL

**UI/UX:** Figma

## Design Pattern:

We will incorporate the design pattern MVC into our project for various reasons. One reason would be that it can help enforce our business rules within the application. Within the controller, we can create and enable files that will both create and manage the business logic and the functionality for the application. Another reason is that it would allow us to handle different changes/modifications within the application with minimal changes to other sections of the application. The MVC pattern supports loose coupling, and since we are anticipating later refactoring to code, it will help significantly later in the timeline. This pattern also aids with support for multiple views. Due to the fact that the CFC project is intended to have support for both the Desktop and mobile devices, MVC will allow us to have two different views without needing to adjust the controller or the model at all.

Model: Solely responsible for interacting with the database and handling all data logic. It creates, retrieves, updates, and deletes any data from the database. The Model only communicates with the controller and provides it with whatever information the controller needs.

View: The frontend graphical user interface that the user interacts with. The View only communicates with the controller. The View is used to display the data using the Model object.

Controller: The intermediary between the View and Model. It requests information from the Model and sends it to the View, or it can get information from the View and send it to the Model. This will handle all of the logic for the business rules.

## Entity Relationship Diagram:

![](https://github.com/ezana234/DAD-Project/blob/main/Milestone%202%20ERD.png)

## Timeline:

- Milestone3 (Layering) - October 8, 2021
- Milestone4 (ExceptionHandling) - October 22, 2021
- Milestone5 (Refactoring) - November 5, 2021
- Milestone6 (Testing) - November 19, 2021
- Milestone7 (Packaging) - November 29, 2021
- Finalized Project Code - December 6, 2021
