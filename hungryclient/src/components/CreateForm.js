import React, {useState} from 'react'
import axios from 'axios'


const INITIAL_FORM = {
    "first_name": "",
    "last_name":"",
    "email": "",
    "password": "",
    "admin": ""
        
}
const CreateForm = () => {



   

    const [formData, setFormData] = useState(INITIAL_FORM)

    const handleChange = e => {
        
        console.log( e.target)
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

        <form onSubmit={handleSubmit}>
            <label htmlFor='first_name'> First Name</label>
            <input 
            onChange = {handleChange}
            type="text"
            name ='first_name' 
            placeholder='Enter Name'
            id='first_name' 
            value={formData.first_name}/>

        <label htmlFor='last_name'> Last Name</label>
        <input 
        onChange = {handleChange}
        type="text" 
        name="last_name" 
        id="last_name"
        value={formData.last_name}/>

        <label htmlFor='email'> Email</label>
        <input 
        onChange = {handleChange}
        type="text" 
        name="email" 
        id="email"
        value={formData.email}/>
        <label htmlFor='email'> Password</label>
        <input 
        onChange = {handleChange}
        type="password" 
        name="password" 
        id="password"
        value={formData.password}/>

        <button onClick={handleSubmit} type="submit">submit</button>



        </form>




    )


}

export default CreateForm