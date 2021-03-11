package routes

import (
	"fmt"
	"net/http"
	// "context"
	
	// "time"
	// "log"
"github.com/gin-gonic/gin"
// "golang.org/x/crypto/bcrypt"
// "go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/mongo/readpref"
	

)

func Admin_get() gin.HandlerFunc{

	fmt.Println("inside Admin route")

	// with the anon function dependencies can be passed and parameters 
	//can be validated before executing function
	return func(c *gin.Context){

		c.JSON(http.StatusOK, map[string]string{
			"Admin only": "Protected !!!",
		})
	


	}
}