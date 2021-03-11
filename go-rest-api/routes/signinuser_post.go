package routes

import (
	"fmt"
	"context"
	"reflect"
	"net/http"
	"time"
	"log"
"github.com/gin-gonic/gin"
"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/mongo/readpref"
	// "golang.org/x/crypto/bcrypt"

	"github.com/dgrijalva/jwt-go"
	"os"

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

func CreateToken(userid,name string, isadmin bool) (string, error) {
	var err error
	//Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userid
	atClaims["name"] = name
	atClaims["admin"] = isadmin
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
	   return "", err
	}
	return token, nil
  }




func Signinuser_post(DB *mongo.Client, ctx context.Context, err error) gin.HandlerFunc{

	





	// with the anon function dependencies can be passed and parameters 
	//can be validated before executing function
	return func(c *gin.Context){

		// Get the user fields from the request and bind them to a user struct
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
			return
		 }
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
			token, err := CreateToken(user.Email,userpp[0].FirstName,userpp[0].Admin)
  			if err != nil {
     		c.JSON(http.StatusUnprocessableEntity, err.Error())
     		return
  				}
			c.JSON(200, gin.H{"success" : map[string]string{
				"first name": userpp[0].FirstName ,
				"last name": userpp[0].LastName,
				"email": userpp[0].Email,
				"token": token,
		} })// Your custom response here
		}else {
		
			c.JSON(http.StatusUnauthorized, "Please provide valid login details")
		}

	// c.JSON(http.StatusOK, map[string]string{
	// 	"post": "post got",
	// })

	}
}