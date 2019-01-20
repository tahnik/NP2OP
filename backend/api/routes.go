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
		Posts.POST("/", s.addPost) //localhost:8080
		Posts.GET("/", s.getPosts)
		Posts.GET("/:id", s.getPost) //localhost:8080/user/sdakjfbdshfbsdihvb
		Posts.PUT("/:id", s.updatePost)
		Posts.DELETE("/:id", s.deletePost)
	}

	return r
}

//Placeholder function demonstrating the gin grouping
func (s *ServerState) addPost(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

//Placeholder function demonstrating the gin grouping
func (s *ServerState) getPosts(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

//Placeholder function demonstrating the gin grouping
func (s *ServerState) getPost(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

//Placeholder function demonstrating the gin grouping
func (s *ServerState) updatePost(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

//Placeholder function demonstrating the gin grouping
func (s *ServerState) deletePost(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

//Placeholder function demonstrating the gin grouping
func (s *ServerState) placeholder(c *gin.Context) {

	//s.DB.DoStuff()
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func (s *ServerState) addUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
