package routes

import (
	"fmt"
	"context"
	"net/http"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
"github.com/gin-gonic/gin"
)


func Working_get(DB *mongo.Client, ctx context.Context, err error) gin.HandlerFunc{

	fmt.Println("working route ", DB)
	err=DB.Ping(ctx,readpref.Primary())
	fmt.Println("db connected main")
	databases, err:= DB.ListDatabaseNames(ctx, bson.M{})

// if err != nil {
//     log.Fatal(err)
// }
fmt.Println("inside working route DB",databases)


	// with the anon function dependencies can be passed and parameters 
	//can be validated before executing function
	return func(c *gin.Context){
	c.JSON(http.StatusOK, map[string]string{
		"working": "router !!!",
	})

	}
}