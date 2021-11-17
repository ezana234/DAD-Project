import React, { useState } from 'react';
import './Login.css'
import { Link, useHistory } from "react-router-dom";
import { auth } from "./firebase";
import axios from 'axios';
import $ from 'jquery';
import jwt from 'jwt-decode'

function Login() {
    const history = useHistory();
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [user, setUser] = useState('');
    const [token, setToken] = useState('');
    const signIn = event => {
        event.preventDefault();
        
        const credentials = { 'email': email, 'password':password};

        console.log('AJAX')
        
        console.log(credentials)
        $.ajax({
            type: 'post',
            url: 'http://localhost:3000/login',
            data: JSON.stringify(credentials),
            contentType: "application/json; charset=utf-8",
            traditional: true,  
            success: function (data) {
                console.log(data);
                console.log(data['token'])
                setToken(data['token'])
                
                const tokenData = jwt(data['token']);
                console.log(tokenData);
                if(tokenData.authorized){
                    if(tokenData.role =='1'){
                        console.log("Yes");
                        console.log(token);
                        console.log(data['token'])
                        console.log("Second request")

                        console.log("Bearer "+data['token'])

                        let url = "http://127.0.0.1:3000/client"

                        
                        const AuthStr = 'Bearer '.concat(data['token']);

                        axios({ method: 'get', url: 'http://127.0.0.1:3000/client', headers: { 'Authorization': 'Bearer ' + data['token'] } })
                        .then((response) => {
                                        console.log("FINAL", response)
                                        if(response.status  == 200){
                                            history.push({
                                                pathname: '/home',
                                                state: response.data
                                            })
                                        }
                                    }, (error) => {
                                        console.log("Error"+error)
                                    }
                                );
                    }
                    else if (tokenData.role =='2'){
                        console.log("Clinician data", data)
                        history.push({
                            pathname: '/clinicianHome',
                            state: data['token']
                        })
                    }
                }
                else{
                    alert("You are not an authorized user");
                }

                },
            error: function (xhr, ajaxOptions, thrownError) {
                alert("Incorrect email or password!");
            }

            });

            
    }

    const signUp = event => {
        event.preventDefault();

        
    }

    const onChangeValue = (event) =>{
        console.log("HERE")
        setUser(event.target.value)
        //console.log(event.target.value)
    }

    return (
        <div className='loginForm'>
                <img
                    className="logo"
                    src='data:image/jpeg;base64,/9j/4AAQSkZJRgABAQAAAQABAAD/2wCEAAkGBw8TEA8QEA8WEBIWEBAVGRIVFxAXFxAWFxUWFhUdGBcYHSggGBolGxYVITIiJikrLi4vFx8zODMsNygtLisBCgoKDg0OGhAQGysmHiIrLTA1LS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0tLS0rLS0tLS0tLS0tLTctLf/AABEIAOAA4AMBIgACEQEDEQH/xAAcAAEAAgMBAQEAAAAAAAAAAAAAAQYEBQcIAgP/xABAEAACAgEBBQUDBg0EAwAAAAAAAQIDEQQFBhIhMQdBUWFxEyKRMjWBobGzFCMzQlJicnN0gpLB0SVTsvAVouH/xAAaAQEAAwEBAQAAAAAAAAAAAAAAAgMEBQEG/8QAJREBAAICAgIBBAMBAAAAAAAAAAECAxEEIRIxURMUIkEjMjOh/9oADAMBAAIRAxEAPwDuIAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEZAkEZPzlqa08OcU/DKA/UHzxeZ9IAAAAAAAAAAAAIbGQJBGQmBIAAAAAAAAAAENkTklnPLzOL9oPaLO2U9LobHClNxndHlK7uag+6Hmub9Os6Y5vOoQvkikblfN5O0LQaRyg5u+1daqsS4X+tL5MfTOfI5ztrtX19raojXpYdzS9pZ/VL3f/UoAN1OPSvvtivntLYa/besv/L6q2zPVSnLhf8q5fUa7BJ9+zl+i/gy6Kx8KZt8y2Gxdv6vSzU9NfKGPzMtwn5Sg+T+09C7obdjrNHTqUuFyypRX5k4vEl6Z+po8zncOxLP/AI+3w/C7P+FeTNyqxrf7aePafLToYAMLaAEZAkEZHEBIIyQ2BU+0He1aDTxcUpX2OSri+iwvelJd8VlfS0cYV+1doWNRlfqp9XGLlwQ+hYhH6izdt1reuojnlHSxa8nKyef+K+BdOx3SwjsyucV71lt0pPvbUnBfCMUa66x44truWS28mTx305nodl7b02oohFajTSsthBSzKVeW0vew3DksvD7kz0HUnhJvLSXPxPpkmfJk8/0vpj8P2kENjJBYkEcRIAAAACGBzzti3ilRpoaWqWLL88TXWNUccX9Taj6cRxAuna/qZT2rbF9K6qYL0cPaP65spZ0sFdUhzs1vK8hs9mbJlZiUvdh498vT/J87H0HtJ5l8iPN+fgi0pclg24sXl3Lm8rk+H419vw0+iqgvcgl59X8WZGQDVFYhy7XtPcyxtToKp/Kgs+K5P4o6Z2XVV1aKNClmXtLJvPLKk+XrySRzzBcNFxVqvheJRUefmkYubji1dR7b+ByLUtufTo6JMTZurVtcZrv6rwfeZTOHManT6esxaNwwtsbUp01Nl99ns4QWW+/ySXe33I4pvN2n629yjp5PSU88Y4fayX60/wA1+Ufi+pkdsm3ZW6taSL/F0KLa7pWyjlv6IyS+mRvOyncur2UNfqYKyU3mqElmNcU3ieOjk+7wWPE1UrWlPOzNa1r38auf0aDa969pGGsuT/PzqHn0bfNDT7c2rorEvb6iiX+3bxuL/ksymvNI9J8Jr9t7G0+qqlTqK1ZF/GD7nF9YteKPPuIn3WNH28x6ntVez/f+OtfsL0qtSk2sfIuS6uOejXXh+nxxj9rmi19kdJ+BRuk07eL2LmsJqOOLha8zK3Q7ONNo7FfZN6i6Mm4Sa4Y1LmliOecuF82/owXhortatb7otitppqzy7tnTauuxR1kbY2cCaVzk5cOXjHE3yzkztk7N2tZUpaWGplVmSTqlYoZT97CTx1LJ22fOFX8JX95aXvsfX+lU/vdR97I1Wy6xxbTLXFu812qHZ5sza9e0aJ6qGqjSlbxO2Vrhzrko5TeOuDqG8e3aNHRK++WIrkor5Vku6MV3v7OptGjz92o7dlqNfbBS/FUN1RWeXEvyj9eLK9Ioz1ic1+19v4qdPveLtJ2hqG1VN6WrOFCt++13cVnXPpg1VWytr2r2iq1lifPift+fnzeWdR7MNyqqaKtXfWp6iyCnHiWfYRksxST6Taxl+eDoXCTtmrSdVhCMNrxu0vN+h3n2ro7OFai6El1pv45LHnCzml6YOxbhb716+EoSj7LUQjmVfVSXTig+9eK6rPfyb2u8+7mn1tMqroJvD4bEveql3OL/ALdGajcvcLT6Fq1t36jha9q1hQT6qEc8vV5f2EL3peu9alOlL0t76XIAFDQEMkAcG7ZNE4bSdnddTXLPi4r2b+qMfiUZLPI752qbtS1ek46o5upbnFLrOL/KR9cLK84rxOJbDo4ro56RzL4dPraOlxp86xDmcr+OZtKx6DTKuuMO/HPzfeZAB2IjUafNWtNpmZACYxbaSWW+7xPXjN2Rp+KxPuj7z9e4sZi7O0irgl+c+bfmZRz8t/KzoYaeNe1h3Su/Kw9JL7H/AGLIyrbpx/GWPuUEvi//AIWrByc+vOdPoOJMzijbzLvrOT2htBy6/hV3wUmo/Ukeid3a4x0mljDlFaehJLuShHBxPta2RKjaNlmPcvSsi/1klGa+Kz/MjofZTvLXfo69NKSV9EFBx75Vx5VyXjyST815lub8sdZhHF+OSYlewQmfnffGEZTnJRjFNuT5KKXVtvojK17fqDR7v716LWcS09ylKLacHmMsJ4yovrHzXibwTEx7eRMT6cN7bfnCr+Er+8tL52PfNNP73UfeyKH22/OFX8JX95aXzse+aaf3uo+9kasn+NWan+0rqeWNX72ps4u/UT4vpm8nqc83b/7IlptoamDTUJ2Stg/GNj4uXpJyj9B5xp7mDkx1EvR9cUkkuiSS9D7Kr2fbyw1mkrbkvb1wjC2PfxJY4seEsZXw6otKZntExOpaKzExuEgxtfrqqa522zUIQi5Sk+iSMDd7ePSayvj01qnjrB8pw/ai+aGp9vdxvTcAA8egAAho5rvhurTDUSv08eCdqbnDkot5WZLwb557m+Z0plT3rf46H7H92aOLaa5ImGPnVicMxLm1lcovEk4vwZ8lvnBPqk/VJn5LSVf7cf6UdmOR8w+cni99SrVGnnN4jFv7F9JvdnbNVfvS96fj3R9P8mcl4ciSu+abdQtx4Ir2AG92JsZtqy1Yj1UX1fm14Ga94pG5a8eKck6hs93dG4VZksSm8+i7l/3xNuQkScy07nbu0pFKxENFvdu1TrtPKmz3ZJ8ULF1rnjGV4rDw13nB9s7B2hs65SnGdTjLMNRU3wS81NdP2Xj0PSp8WVKSaksp9zw0WY8006/SGTFF+/24No+1TasIqMp1XfrTrXE/XgcU/ganbm9+0dbiu25yjJrFFUcRk/2Y+9P0bZ3a/c/Zk3xS0FDfiqq1n1wuZmbP2JpKPyGmqpfjCEIt+rSyy362OO4r2r+jeept05v2Y7g212Q12sg65Ry6qXykm1jisx05N4j55fRHVwkSZ73m87ldSkUjUOG9tnzhV/CV/eWl77Hn/pVP73UfeyLPrdh6O6SnfpabpJY4rK6ptLrjMk3jr8TI0Wiqpgq6a41QTbUIRjGKzzeFFYJ2ybpFfhCuOYvNn7lZ353Rq19Kj8i6GXXb+jnrGXjB4WfTJZwVxMxO4W2iJjUvM+r0O0dm3qUlZpbFlRsi/dmvKS92S8n9KRvtP2rbUjHhbpsePlSrw/p4Wl8Ed1u08JpxnFTi+qkk0/VM09m5uy5PL2fRnyrgvsRo+vW396s/0LV/rLg2194tobQnGFtkrm37tFccRz3YhFc/V5Z03sw3Fs0r/DNUuG+UXGFSf5KL6uTXJyfhzx6l80GydNQsUUV0p9fZwhHPrhczMwRvm3HjWNQlTDqfK07kRIBQvAABBp94NmO2MZQ+XHPL9JM3JGD2tprO4QyUi9fGXPLapReJRcX4NYPzydFcPFZIjVFdIpeiRq+6n4YJ4HfVv+KDVpbJfJhKXonj4my0u790vlYrXnzfwX+S3YJwQtybT6WU4NI99tXoNh1V4eOOXjLu9F3G0wSCi1pt7bKUrSNVgAB4kAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAD/9k=' 
                />

            <div className='container'>
                <h1>SignIn</h1>

                <form>
                    <h5>Username:</h5>
                    <input type='text' value={email} onChange={event => setEmail(event.target.value)} />

                    <h5>Password:</h5>
                    <input type='password' value={password} onChange={event => setPassword(event.target.value)} />
                    {/* <p>You are:</p>
                    <div onChange={onChangeValue}>
                        <input type="radio" id="client" name="radio" value="client"/>
                        <label for="client">Client</label>
                        &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
                        <input type="radio" id="clinician" name="radio" value="clinician"/>
                        <label for="clinician">Clinician</label>
                        &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
                        <input type="radio" id="Other" name="radio" value="Other"/>
                        <label for="Other">Other</label>
                    </div> */}
                    <button type='submit' onClick={signIn} className='signInButton'>Sign In</button>
                </form>

                <p>
                    By signing-in you agree to the CFC Conditions & Policies. Please
                    see our Privacy Notice.
                </p>

                <button onClick={signUp} className='signUpButton'>Sign Up</button>
            </div>
        </div>
    )
}

export default Login