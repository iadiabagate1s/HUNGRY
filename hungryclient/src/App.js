import logo from './logo.svg';
import React, {useState} from 'react'
import './App.css';
import 'bootstrap/dist/css/bootstrap.min.css';
import {
  BrowserRouter as Router,
  Switch,
  Route,
  Link
} from "react-router-dom";


import CreateForm from './components/CreateForm'
import LoginForm from './components/LoginForm'
import NavB from './components/NavB'
import AdminCom from './components/AdminCom'
function App() {

  const [user, setUser]=useState(null)
  console.log('this is user profile', user)

  return (
  <Router>
    <NavB/>
    
    {/* <Link to="/login"><h6>Log in</h6></Link>
    <Link  to="/signup"><h6>Sign Up</h6></Link>

    <Link  to="/"><h6>HOME</h6></Link>
    <Link  to="/admin"><h6>Admin Only</h6></Link> */}



      <Route path="/admin">
            <AdminCom user={user}/>
          </Route>

      <Route exact path="/login">
            <LoginForm setuser={setUser}/>
          </Route>

      <Route exact path="/signup">
           <CreateForm/>
          </Route>

      <Route exact path="/">
           <h1>Home directory</h1>
          </Route>





  </Router>
     
  
  );
}

export default App;
