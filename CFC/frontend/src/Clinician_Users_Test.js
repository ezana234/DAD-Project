import React from 'react'
import axios from 'axios';

const Clinician_Users_Test = (token) => {
    return axios({ method: 'get', url: 'http://127.0.0.1:3000/clinician/clients', headers: {'Authorization': 'Bearer ' + token}})
    .then(response => response.data);
        
    
}

export default Clinician_Users_Test