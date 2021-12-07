import React from 'react'
import axios from 'axios';

const Login_Test = (credentials) => {
    return axios({ method: 'post', url: 'http://127.0.0.1:3000/login', data: JSON.stringify(credentials)})
    .then(response => response.data);
        
    
}

export default Login_Test
