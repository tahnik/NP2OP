package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//NewRouter returns a router pointer
func (s *ServerState) NewRouter() *gin.Engine {
	r := gin.Default()

	Users := r.Group("/user/")
	{
		//placeholder handler functions demonstrating the grouping of the API
		Users.POST("/", s.addUser) //localhost:8080
		Users.GET("/", s.placeholder)
		Users.GET("/:id", s.placeholder) //localhost:8080/user/sdakjfbdshfbsdihvb
		Users.PUT("/:id", s.placeholder)
		Users.DELETE("/:id", s.placeholder)
	}

	Posts := r.Group("/posts/")
	{
		//placeholder handler functions demonstrating the grouping of the API
		Posts.POST("/", s.addUser) //localhost:8080
		Posts.GET("/", s.placeholder)
		Posts.GET("/:id", s.placeholder) //localhost:8080/user/sdakjfbdshfbsdihvb
		Posts.PUT("/:id", s.placeholder)
		Posts.DELETE("/:id", s.placeholder)
	}

	return r
}

//Placeholder function demonstrating the gin grouping
func (s *ServerState) placeholder(c *gin.Context) {

	//s.DB.DoStuff()
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func (s *ServerState) addUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
