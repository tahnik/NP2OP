package api

import (
	"fmt"
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
		Posts.POST("/", s.addPost)
		Posts.GET("/", s.getPosts)
		Posts.GET("/:id", s.getPost)
		Posts.PUT("/:id", s.updatePost)
		Posts.DELETE("/:id", s.deletePost)
	}

	Campaigns := r.Group("/campaigns")
	{
		Campaigns.POST("/", s.addCampaign)
		Campaigns.GET("/", s.getCampaigns)
		Campaigns.GET("/:id", s.getCampaign)
		Campaigns.PUT("/:id", s.updateCampaign)
		Campaigns.GET("/funding", s.getCampaignFunding)
		Campaigns.GET("/approve", s.approveCampaign)
	}

	return r
}

//Placeholder function demonstrating the gin grouping
func (s *ServerState) addPost(c *gin.Context) {
	var post Post
	if err := c.ShouldBind(&post); err == nil {
		//Validate
		if post.ID == 0 {
			c.JSON(400, gin.H{"status": "Post.ID cannot be zero"})
		}
		if post.Description == "" {
			c.JSON(400, gin.H{"status": "Post.Description cannot be blank"})
		}
		if post.Description == "" || len(post.Description) < 100 {
			c.JSON(400, gin.H{"status": "Post.Description cannot be less than 100 characters"})
		}
		if post.Goal == 0 || post.Goal < 0 {
			c.JSON(400, gin.H{"status": "post.Goal is less than 0 or a negative amount"})
		}
	} else {
		c.JSON(400, gin.H{"status": fmt.Sprintf("Failed to bind post, error: %s", err.Error())})
	}

	query := `INSERT INTO 'post' ('description', 'total_campaign_snapshot') 
	VALUES (?, ?)`
	res := s.DB.MustExec(query, post.ID, post.Description)
	id, err := res.LastInsertId()
	if err != nil {
		c.JSON(500, gin.H{"status": "Failure returning last inserted ID"})
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{"status": fmt.Sprintf("User inserted successfully, ID is %d", id)})
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
func (s *ServerState) getUsers(c *gin.Context) {
	var u []User
	if err := s.DB.Select(&u, "select * from users"); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err})
	}

	c.JSON(http.StatusOK, gin.H{"user": u})
}

func (s *ServerState) getUser(c *gin.Context) {
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

func (s *ServerState) addCampaign(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func (s *ServerState) getCampaigns(c *gin.Context) {
	var cmpg []Campaign
	if err := s.DB.Select(&cmpg, "select * from campaign"); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err})
	}

	c.JSON(http.StatusOK, gin.H{"user": cmpg})
}

func (s *ServerState) getCampaign(c *gin.Context) {
	var cmpgn Campaign
	if err := c.ShouldBindUri(&cmpgn); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err})
	}

	if err := s.DB.Select(&cmpgn, "select "+string(cmpgn.Id)+" from campaign"); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err})
	}

	c.JSON(http.StatusOK, gin.H{"campaign": cmpgn})
}

func (s *ServerState) updateCampaign(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func (s *ServerState) getCampaignFunding(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func (s *ServerState) approveCampaign(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
