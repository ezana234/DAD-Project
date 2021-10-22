import './App.css';
import {useState} from 'react';
import {Link} from 'react-router-dom';
function App() {

  const [inputField , setInputField] = useState({
    email: '',
    password: '',
})
  const handleSubmit = () => {
    console.log(inputField.email);
    console.log(inputField.email);
  }
  const handleChange = (event) => {
    setInputField( {[event.target.name]: event.target.value} )
  }

  return (
    <div className="App">
      <h1 style={{"textAlign":"center"}}>Login Page</h1>
      <br/>
      <form>
        <label>
          Email: &nbsp;
          <input name="email" type="gmail" onChange={handleChange} value={inputField.email} />
        </label>
        &nbsp; &nbsp;
        <label>
          Password: &nbsp;
          <input name="password" type="text" onChange={handleChange} value={inputField.password} />
        </label>
        <br></br><br></br><br></br>
        <Link to="/home"><button onClick={handleSubmit}>Login</button></Link>
      </form>
    </div>
  );
}

export default App;
