package api

import (
	"google.golang.org/genproto/googleapis/ads/googleads/v0/enums"
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
		Users.GET("/", s.getUsers)
		Users.GET("/:id", s.getUser) //localhost:8080/user/sdakjfbdshfbsdihvb
		//Users.PUT("/:id", s.placeholder)
		//Users.DELETE("/:id", s.placeholder)
	}

	Posts := r.Group("/posts/")
	{
		//placeholder handler functions demonstrating the grouping of the API
		//Posts.POST("/", s.addUser) //localhost:8080
		//Posts.GET("/", s.getUsers)
		//Posts.GET("/:id", s.getUser) //localhost:8080/user/sdakjfbdshfbsdihvb
		Posts.PUT("/:id", s.placeholder)
		Posts.DELETE("/:id", s.placeholder)
	}

	return r
}

//Placeholder function demonstrating the gin grouping
func (s *ServerState) getUsers(c *gin.Context) {
	var u []User
	if err := s.DB.Select(&u, "select * from users"); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err})
	}

	c.JSON(http.StatusOK, gin.H{"user": u})
}

func (s *ServerState) getUser(c *gin.Context)  {
	var u User
	if err := c.ShouldBindUri(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err})
	}

	if err := s.DB.Select(&u, "select "+string(u.Id)+" from users"); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err})
	}

	c.JSON(http.StatusOK, gin.H{"user": u})
}

func (s *ServerState) addUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
