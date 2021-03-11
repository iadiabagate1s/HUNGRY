import React, {useState} from 'react'
import axios from 'axios'
import Form from 'react-bootstrap/Form'
import Button from 'react-bootstrap/Button'

const INITIAL_FORM = {
    "first_name": "",
    "last_name":"",
    "email": "",
    "password": "",
    "admin": false
        
}
const CreateForm = () => {



   

    const [formData, setFormData] = useState(INITIAL_FORM)
    console.log("form data^^^",formData)
    const handleChange = e => {
        
      
        const {name, value} = e.target
        setFormData( data => ({
            ...data,
            [name] : value
        }))

    }
    async function createat(createobj) {
        let res = await axios.post('http://localhost:8080/createuser',createobj)
        console.log('res^^', res.data)
        return res.data
    }
    const handleSubmit = (e) => {
        e.preventDefault()
        console.log(formData)
        
        let send = createat(formData)
        console.log('end', send)
    }
   

    return (

        // <form onSubmit={handleSubmit}>
        //     <label htmlFor='first_name'> First Name</label>
        //     <input 
        //     onChange = {handleChange}
        //     type="text"
        //     name ='first_name' 
        //     placeholder='Enter Name'
        //     id='first_name' 
        //     value={formData.first_name}/>

        // <label htmlFor='last_name'> Last Name</label>
        // <input 
        // onChange = {handleChange}
        // type="text" 
        // name="last_name" 
        // id="last_name"
        // value={formData.last_name}/>

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
          <Form.Label>First Name</Form.Label>
          <Form.Control name='first_name' value={formData.first_name} onChange={handleChange} type="text" placeholder="Enter First Name" />
        
        </Form.Group>

        <Form.Group  controlId="formBasicEmail">
          <Form.Label>Last Name</Form.Label>
          <Form.Control name='last_name' value={formData.last_name} onChange={handleChange} type="text" placeholder="Enter Last Name" />
        
        </Form.Group>

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
          Create User
        </Button>
      </Form>




    )


}

export default CreateForm