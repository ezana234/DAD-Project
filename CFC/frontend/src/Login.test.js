import axios from 'axios';
import MockAdapter from 'axios-mock-adapter';
import Login_Test from './Login_Test'


    test('Test login route', async() => {
        var mock = new MockAdapter(axios);
        const data = {
            "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJlbWFpbCI6ImNsaWVudEBnbWFpbC5jb20iLCJleHAiOjE2Mzc4MTIwNTIsInJvbGUiOiIxIiwidXNlcklEIjoxMDA1fQ.cxX5sESNtFB9WG3D43XnboWSrxYtbJ8dzYwhKp1Y9mM"
        };
        mock.onPost('http://127.0.0.1:3000/login').reply(200, data);

        const resp = await Login_Test({"email": "client@gmail.com", "password": "cliepassword"});
        
        expect(resp).toEqual(data);

    });
