package routes

import (
	"fmt"
	"context"
	
	"time"
	"log"
"github.com/gin-gonic/gin"
"golang.org/x/crypto/bcrypt"
"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	

)

// Profile is the model for the profile table.
type Profile struct {
   
    FirstName string `form:"first_name" json:"first_name" bson:"first_name"`
    LastName  string `form:"last_name" json:"last_name" bson:"last_name"`
    Password  string `form:"password" json:"password" bson:"password"`
	Email  string `form:"email" json:"email" bson:"email"`
	Admin bool `form:"admin" json:"admin" bson:"admin"`
  
}

func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

var client *mongo.Client

func Createuser_post() gin.HandlerFunc{

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
	   "mongodb+srv://HUNGRY2021:HUNGRY2021@cluster0.ntcly.mongodb.net/Cluster0?retryWrites=true&w=majority",
	))
	if err != nil { log.Fatal(err) }
	
	err=client.Ping(ctx,readpref.Primary() )
	fmt.Println("db connected in Create ")
	databases, err:= client.ListDatabaseNames(ctx, bson.M{})
	
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("inside create****",databases)





	// with the anon function dependencies can be passed and parameters 
	//can be validated before executing function
	return func(c *gin.Context){

		// Get the user fields from the request and bind them to a user struct
		var profile Profile
		
		c.BindJSON(&profile)
		
	
		fmt.Println(profile.Password)
		newpass, err := HashPassword(profile.Password)
		fmt.Println(newpass, err)
		profile.Password = newpass
		fmt.Println(profile)

		collection := client.Database("HUNGRY").Collection(profile.Email)

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		result, err := collection.InsertOne(ctx, profile)

		fmt.Println(result)

	
        c.JSON(200, gin.H{"status create": profile}) // Your custom response here


	// c.JSON(http.StatusOK, map[string]string{
	// 	"post": "post got",
	// })

	}
}