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
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
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




func Signinuser_post() gin.HandlerFunc{

	








	// with the anon function dependencies can be passed and parameters 
	//can be validated before executing function
	return func(c *gin.Context){

		// Get the user fields from the request and bind them to a user struct
		var user User
		
		c.BindJSON(&user)

		var userpp []UserProfile
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
	   "mongodb+srv://HUNGRY2021:HUNGRY2021@cluster0.ntcly.mongodb.net/Cluster0?retryWrites=true&w=majority",
	))
	if err != nil { log.Fatal(err) }
	
	err=client.Ping(ctx,readpref.Primary() )
	fmt.Println("db connected in signin")
	databases, err:= client.ListDatabaseNames(ctx, bson.M{})
	
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("inside create****",databases)
	dbobj := client.Database("HUNGRY")
	fmt.Println("inside sign in user obj",dbobj, reflect.TypeOf(dbobj), user.Email)

	userobj := dbobj.Collection(user.Email)
	cursor , err := userobj.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var userp UserProfile
		cursor.Decode(&userp)
		userpp = append(userpp, userp)
	}
	newpass := userpp[0].Password
	fmt.Println("inside sign in user obj*****",userpp, reflect.TypeOf(userpp), userpp[0].Password)




		
		
		// var email = user.Email
		fmt.Println(user.Password)
		var oldpass = user.Password

	
		fmt.Println(user)

		

		var res = CheckPasswordHash(oldpass, newpass)
		
		fmt.Println(res)



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