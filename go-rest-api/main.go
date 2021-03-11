package main

import ("fmt"
	// "context"
	// "reflect"
	// "time"
	// "log"
	"go-rest-api/routes"
	"github.com/gin-gonic/gin"
	// "go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/mongo/readpref"

)

func main(){
fmt.Println("working !!!!!")

// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// defer cancel()
// client, err := mongo.Connect(ctx, options.Client().ApplyURI(
//    "mongodb+srv://HUNGRY2021:HUNGRY2021@cluster0.ntcly.mongodb.net/Cluster0?retryWrites=true&w=majority",
// ))
// if err != nil { log.Fatal(err) }

// err=client.Ping(ctx,readpref.Primary() )
// fmt.Println("db connected")
// databases, err:= client.ListDatabaseNames(ctx, bson.M{})

// if err != nil {
//     log.Fatal(err)
// }
// fmt.Println(databases)
// fmt.Println(reflect.TypeOf(client))
r := gin.Default()


	
	r.GET("/working",  routes.Working_get())
	r.POST("/createuser", routes.Createuser_post())
	r.POST("/login", routes.Signinuser_post())
	


	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}