import logo from './logo.svg';
import './App.css';
import {
  BrowserRouter as Router,
  Switch,
  Route,
  Link
} from "react-router-dom";

function App() {
  return (
  <Router>

    
    <Link to="/login"><h1>Log in</h1></Link>

    <Link  to="/signup"><h1>Sign Up</h1></Link>

    <Link  to="/"><h1>HOME</h1></Link>
    <Link  to="/admin"><h1>Admin Only</h1></Link>



      <Route path="/admin">
            <h1>Admin Only</h1>
          </Route>

      <Route exact path="/login">
            <h1>Log in page Here</h1>
          </Route>

      <Route exact path="/signup">
           <h1>Sign up page here</h1>
          </Route>

      <Route exact path="/">
           <h1>Home directory</h1>
          </Route>





  </Router>
     
  
  );
}

export default App;
