import React from 'react'
import axios from 'axios';

const Client_SafetyPlan_Test = (token) => {
    return axios({ method: 'get', url: 'http://127.0.0.1:3000/client/safetyplan', headers: {'Authorization': 'Bearer ' + token}})
    .then(response => response.data);
        
    
}

export default Client_SafetyPlan_Test