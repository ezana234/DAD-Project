import axios from 'axios';
import MockAdapter from 'axios-mock-adapter';
import Client_SafetyPlan_Test from './Client_SafetyPlan_Test'


    test('Test for safetyplan of a client', async() => {
        var mock = new MockAdapter(axios);
        const token = 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJlbWFpbCI6ImNsaWVudEBnbWFpbC5jb20iLCJleHAiOjE2Mzc4MTI5MDAsInJvbGUiOiIxIiwidXNlcklEIjoxMDA1fQ.bmqBPT73VCE4NT0bIoN_jfrYwDBxQgwEVeYpo3UgtuE';
        const data = [
            {
                "SafetyID": 3,
                "Triggers": "Puppies",
                "WarningSigns": "Manic Behavior",
                "DestructiveBehaviors": "N/A",
                "InternalStrategies": "Playing with cats",
                "UpdatedClinician": 341,
                "UpdatedDatetime": "2021-11-16T22:51:29Z",
                "ClientID": 334,
                "ClinicianID": 341
            }
        ];
        mock.onGet('http://127.0.0.1:3000/client/safetyplan',{'Authorization': 'Bearer ' + token}).reply(200, data);

        const resp = await Client_SafetyPlan_Test(token);
        
        expect(resp).toEqual(data);

    });