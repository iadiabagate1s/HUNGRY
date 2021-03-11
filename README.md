# Hungry-tech

## Backend 
#### golang
    used :
        Mongo-go driver 
        Gin-gonic web Framework
        bcrypt : for password hashing and encryption (authentication)
        Json web tokens: for encrypting tokens and payload. 
            mainly for communicating with frontend.

cd into go-rest-api directory

run commands to start the server

    make dev 
the server runs on 

    http://localhost:8080/
theres are 3 routes available 

### Create User: POST/createuser

accepts json structured like :

    {
    "first_name": "",
	"last_name":"",
	"email": "",
	"password": "",
    (optional- defaults to false)admin: bool
    }
### Log in: POST/login
#### returns a token needed to access admin route 
accepts json structured like:

       {
	"email": "",
	"password": "",
   
    }
### Admin: POST/admin
accepts token received from log in route 

Returns message letting you know wether you're and admin or not
simulating authorization

accepts json structures like:

    {
        "token": ""
    }

# ***Extra FrontEnd *** 
cd into hungryclient 
start front end 

    npm start

created a quick front end in React with 2 forms 
one for creating a user
one for Logging in 

the base functionality works so submitting and logging in sends those same 
requests to the API

## Still Needs - 
    -redirects after register and Log in

    - global store to keep user info and token when switching pages 

    -Admin page needs to connect to API to check if the user is an admin or not 
    and display a dynamic message according to results 



