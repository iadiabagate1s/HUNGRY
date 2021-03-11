import React, {useState} from 'react'
import axios from 'axios'
import Form from 'react-bootstrap/Form'
import Button from 'react-bootstrap/Button'




const INITIAL_FORM = {
    "email": "",
    "password": "",	
}

const LoginForm = ({setuser}) => {



    const [formData, setFormData] = useState(INITIAL_FORM)

    const handleChange = e => {
        
        console.log( e.target)
        const {name, value} = e.target
        setFormData( data => ({
            ...data,
            [name] : value
        }))


    }

    async function loginat(loginobj) {
        try{
        let res = await axios.post('http://localhost:8080/login',loginobj)
        console.log('res^^^', res.data['success'])
        setuser(res.data['success'])
        return res.data
        }
        catch(err){
            alert("incorrect Email/password")
        }
    }
    const handleSubmit = (e) => {
        e.preventDefault()
        console.log(formData)
        let send = loginat(formData)
        console.log('end', send)
    }

    return (

        // <form>
        // <label htmlFor='email'> Email</label>
        // <input 
        // onChange = {handleChange}
        // type="text" 
        // name="email" 
        // id="email"
        // value={formData.email}/>
        // <label htmlFor='email'> Password</label>
        // <input 
        // onChange = {handleChange}
        // type="password" 
        // name="password" 
        // id="password"
        // value={formData.password}/>

        // <button onClick={handleSubmit} type="submit">submit</button>



        // </form>
        <Form style={{marginTop:'50px', width:'50vw'}}>
  <Form.Group  controlId="formBasicEmail">
    <Form.Label>Email address</Form.Label>
    <Form.Control name='email' value={formData.email} onChange={handleChange} type="email" placeholder="Enter email" />
    <Form.Text className="text-muted">
      We'll never share your email with anyone else.
    </Form.Text>
  </Form.Group>

  <Form.Group controlId="formBasicPassword">
    <Form.Label>Password</Form.Label>
    <Form.Control name="password"  value={formData.password} onChange={handleChange} type="password" placeholder="Password" />
  </Form.Group>
  
  <Button onClick={handleSubmit} variant="primary" type="submit">
    Login
  </Button>
</Form>




    )


}

export default LoginForm