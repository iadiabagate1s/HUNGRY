import React from 'react'
import Navbar from 'react-bootstrap/Navbar'
import {
    BrowserRouter as Router,
    Switch,
    Route,
    Link
  } from "react-router-dom"
  import Nav from 'react-bootstrap/Nav'


export default function NavB() {
    return (
        <Navbar bg="dark" variant="dark">
        <Navbar.Brand href="#home">Navbar</Navbar.Brand>
        <Nav className="mr-auto">
          <Nav.Link href="/">Home</Nav.Link>
          <Nav.Link href="/login">Login</Nav.Link>
          <Nav.Link href="/signup">Register</Nav.Link>
          <Nav.Link href="/admin">Admin</Nav.Link>
        </Nav>
      
      </Navbar>
    )
}
