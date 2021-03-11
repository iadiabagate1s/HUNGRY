package routes

import (
	"fmt"
	"context"
	"reflect"
	"time"
	"log"
"github.com/gin-gonic/gin"
"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/mongo/readpref"
	// "golang.org/x/crypto/bcrypt"

)

// Profile is the model for the profile table.
type User struct {
   
	Email  string `form:"email" json:"email"`
    Password  string `form:"password" json:"password"`
   
  
}
type UserProfile struct {
   
    FirstName string `form:"first_name" json:"first_name" bson:"first_name"`
    LastName  string `form:"last_name" json:"last_name" bson:"last_name"`
    Password  string `form:"password" json:"password" bson:"password"`
	Email  string `form:"email" json:"email" bson:"email"`
	Admin bool `form:"admin" json:"admin" bson:"admin"`
  
}
type UserRes struct {
   
    FirstName string `form:"first_name" json:"first_name" bson:"first_name"`
    LastName  string `form:"last_name" json:"last_name" bson:"last_name"`
	Email  string `form:"email" json:"email" bson:"email"`

  
}




func Signinuser_post(DB *mongo.Client, ctx context.Context, err error) gin.HandlerFunc{

	





	// with the anon function dependencies can be passed and parameters 
	//can be validated before executing function
	return func(c *gin.Context){

		// Get the user fields from the request and bind them to a user struct
		var user User
		c.BindJSON(&user)
		//create a slice of user profile
		var userpp []UserProfile

	
	//connect to database and find the collection by email
	dbobj := DB.Database("HUNGRY")
	fmt.Println("inside sign in user obj",dbobj, reflect.TypeOf(dbobj), user.Email)
	userobj := dbobj.Collection(user.Email)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	//once thats found return all data inside the collection
	cursor , err := userobj.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)
	//decrypting binary json
	for cursor.Next(ctx) {
		var userp UserProfile
		cursor.Decode(&userp)
		//append data  obj to slice made earlier
		userpp = append(userpp, userp)
	}
	//get saved hashed password
	savedhashedpass := userpp[0].Password
	fmt.Println("inside sign in user obj*****",userpp, reflect.TypeOf(userpp), userpp[0].Password)




		
		
		// var email = user.Email
		fmt.Println(user.Password)
		//entered password
		var enteredpass = user.Password

	
		fmt.Println(user)

		
		//check is entered password is the same as the saved hashed password 
		//if true returns true 
		var res = CheckPasswordHash(enteredpass, savedhashedpass)
		
		fmt.Println(res)


		//if the passwords match return user information 
		//if false return message
		if res {
			c.JSON(200, gin.H{"status success" : map[string]string{
				"first name": userpp[0].FirstName ,
				"last name": userpp[0].LastName,
				"email": userpp[0].Email,
		} })// Your custom response here
		}else {
			c.JSON(200, gin.H{"status failed" : map[string]string{
					"worked": "false" ,
					"message": "incorect email or password",
			} })
		}

	// c.JSON(http.StatusOK, map[string]string{
	// 	"post": "post got",
	// })

	}
}