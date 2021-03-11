import React, {useState, useEffect} from 'react'
import axios from 'axios'

export default function AdminCom({user}) {

    const [auth , setAuth]=useState(null)
    console.log('auth',auth)
    console.log(user)
    useEffect(() => {
      return () => {
        axios.post('http://localhost:8080/admin', {token: user.token}).then(u => setAuth(u.data))
      };
    }, [])

    return (
        <div>
            <h1>Are you an admin ?</h1>
            
            
        </div>
    )
}
