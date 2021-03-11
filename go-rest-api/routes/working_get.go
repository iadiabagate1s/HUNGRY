package routes

import ("net/http"
"github.com/gin-gonic/gin"
)


func Working_get() gin.HandlerFunc{

	// with the anon function dependencies can be passed and parameters 
	//can be validated before executing function
	return func(c *gin.Context){
	c.JSON(http.StatusOK, map[string]string{
		"working": "router !!!",
	})

	}
}