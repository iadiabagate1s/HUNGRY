package routes

import (
	"fmt"
	"context"
	
	"time"
	// "log"
"github.com/gin-gonic/gin"
"golang.org/x/crypto/bcrypt"
"go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/mongo/readpref"
	

)

// Profile is the model for the profile table.
type Profile struct {
   
    FirstName string `form:"first_name" json:"first_name" bson:"first_name"`
    LastName  string `form:"last_name" json:"last_name" bson:"last_name"`
    Password  string `form:"password" json:"password" bson:"password"`
	Email  string `form:"email" json:"email" bson:"email"`
	Admin bool `form:"admin" json:"admin" bson:"admin"`
  
}
//hashing a password with bcrypt
func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

//checking a hashed password vs unhashed to see if its the same 
//returns true if it is 
func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}



func Createuser_post(DB *mongo.Client, ctx context.Context, err error) gin.HandlerFunc{


	// with the anon function dependencies can be passed and parameters 
	//can be validated before executing function
	return func(c *gin.Context){

		// Get the user fields from the request and bind them to a user struct
		var profile Profile
		
		//binds request body to strcut 
		c.BindJSON(&profile)
		
	
		fmt.Println(profile.Password)
		//hash password 
		newpass, err := HashPassword(profile.Password)
		fmt.Println(newpass, err)
		//set password in struct to the hashed password
		profile.Password = newpass
		fmt.Println(profile)

		//create a collection with he email as the name 
		collection := DB.Database("HUNGRY").Collection(profile.Email)
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		//insert profile data into collection 
		result, err := collection.InsertOne(ctx, profile)

		fmt.Println("added to db res",result)

	
        c.JSON(201, gin.H{"status create": "done"}) // Your custom response here


	}
}