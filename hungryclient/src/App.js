import logo from './logo.svg';
import './App.css';
import {
  BrowserRouter as Router,
  Switch,
  Route,
  Link
} from "react-router-dom";

import CreateForm from './components/CreateForm'
import LoginForm from './components/LoginForm'
function App() {
  return (
  <Router>

    
    <Link to="/login"><h6>Log in</h6></Link>
    <Link  to="/signup"><h6>Sign Up</h6></Link>

    <Link  to="/"><h6>HOME</h6></Link>
    <Link  to="/admin"><h6>Admin Only</h6></Link>



      <Route path="/admin">
            <h1>Admin Only</h1>
          </Route>

      <Route exact path="/login">
            <LoginForm/>
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
