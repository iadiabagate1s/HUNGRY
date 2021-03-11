package main

import ("fmt"
	"context"
	// "reflect"
	"time"
	"log"
	"go-rest-api/routes"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/mongo/readpref"
	

)
var DB *mongo.Client


func main(){
fmt.Println("working !!!!!")

ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()
DB, err := mongo.Connect(ctx, options.Client().ApplyURI(
   "mongodb+srv://HUNGRY2021:HUNGRY2021@cluster0.ntcly.mongodb.net/Cluster0?retryWrites=true&w=majority",
))
if err != nil { log.Fatal(err) }


// fmt.Println(reflect.TypeOf(client))
r := gin.Default()

r.Use(cors.Default())//allow all cors
	
	r.GET("/working",  routes.Working_get(DB,ctx,err))
	r.POST("/admin",  routes.Admin_get())
	r.POST("/createuser", routes.Createuser_post(DB,ctx,err))
	r.POST("/login", routes.Signinuser_post(DB,ctx,err))
	

	
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}