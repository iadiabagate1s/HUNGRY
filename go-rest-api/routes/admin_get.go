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
	"github.com/dgrijalva/jwt-go"
	"os"
	

)
type Token struct {
   
    Token string `form:"token" json:"token" bson:"token"`

  
}
type Res struct {
   
    name string `form:"name"  bson:"name"`
    admin bool `form:"admin"  bson:"admin"`
    

  
}

func Admin_get() gin.HandlerFunc{

	fmt.Println("inside Admin route")

	
	return func(c *gin.Context){
		var tok Token
		//binds jsson body to struct or returns message 
		if err := c.ShouldBindJSON(&tok); err != nil {
			c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
			return
		 }

		 fmt.Println("token in admin", tok)
		 tokenString := tok.Token
		 //sets map to hold decoded json web token payload 
		 claims := jwt.MapClaims{}
		 //parse encoded Json web token and decodes is
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
    return []byte(os.Getenv("ACCESS_SECRET")), nil
})
if err != nil {
	c.JSON(http.StatusUnprocessableEntity, "Invalid Token or expired Token")
	return 
 }
fmt.Println(token,err)
fmt.Println("claims amp>>>",claims["admin"])
//checks to see if payload variable for admin is true or not
//if its true user is allowed if not user is NOT allowed
if claims["admin"] == true {
	c.JSON(http.StatusOK, map[string]string{
		"Admin only": "welcome you are allowed",
	})

}else {
	c.JSON(http.StatusOK, map[string]string{
		"Admin only": "Protected !!! You are not allowed",
	})
}
		
	


	}
}