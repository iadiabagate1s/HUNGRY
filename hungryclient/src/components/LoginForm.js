import React, {useState} from 'react'
import axios from 'axios'




const INITIAL_FORM = {
    "email": "",
    "password": "",	
}

const LoginForm = () => {



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
        let res = await axios.post('http://localhost:8080/login',loginobj)
        console.log('res^^', res.data)
        return res.data
    }
    const handleSubmit = (e) => {
        e.preventDefault()
        console.log(formData)
        let send = loginat(formData)
        console.log('end', send)
    }

    return (

        <form>
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

export default LoginForm