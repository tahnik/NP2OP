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
		Campaigns.PUT("/:id/funding", s.getCampaignFunding)
		Campaigns.PUT("/:id/approve", s.approveCampaign)
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
		if post.Cost == 0 || post.Cost < 0 {
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
	query := `select u.name, u.username , p.description, p.timestamp, p.total_campaign_snapshot, co.name, co.cost, sc.name
	from ht.post p, ht.campaign ca, ht.course co, ht.school sc, ht.user u
	where p.campaign_id = ca.id AND ca.course_id = co.id AND co.school_id = sc.id AND ca.user_id = u.id;`

	//curl --header "Content-Type: application/json" --request GET http://localhost:8080/posts/

	var posts []Post

	err := s.DB.Select(&posts, query)
	if err != nil {
		panic(err)
	}
	fmt.Println(posts)

	c.JSON(http.StatusOK, gin.H{"status": posts})
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
	if err := c.ShouldBind(&u); err != nil {
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
	var ins InsertCampaign
	query := `INSERT INTO campaign (title, description, total_received, user_id, course_id) VALUES (?, ?, ?, ?, ?);`

	/* MODIFY DATA DUE TO KEY CONSTRAINTS

	curl --header "Content-Type: application/json" --request POST --data \
	    '{"title":"This is my Title", "description":"DESCRIPTION GOES HERE","userid": 5, "courseid": 3}' \
	     http://localhost:8080/campaigns/
	*/
	if err := c.ShouldBind(&ins); err == nil {
		//Validate
		if ins.UserID == 0 {
			c.JSON(400, gin.H{"status": "ins.ID cannot be zero"})
		}
		if ins.Description == "" {
			c.JSON(400, gin.H{"status": "ins.Description cannot be blank"})
		}
		if ins.CourseID == 0 {
			c.JSON(400, gin.H{"status": "ins.Description cannot be zero"})
		}
		if ins.Title == "" {
			c.JSON(400, gin.H{"status": "ins.Title is blank"})
		}
	} else {
		c.JSON(400, gin.H{"status": fmt.Sprintf("Failed to bind campaing.insert, error: %s", err.Error())})
	}

	res := s.DB.MustExec(query, ins.Title, ins.Description, 0, ins.UserID, ins.CourseID)
	id, err := res.LastInsertId()
	if err != nil {
		c.JSON(500, gin.H{"status": "Failure returning last inserted ID"})
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{"status": fmt.Sprintf("Campaign inserted successfully, ID is %d", id)})
}

func (s *ServerState) getCampaigns(c *gin.Context) {
	query := `
select 
	u.name, 
	u.username, 
	u.email, 
	u.phone, 
	u.country, 
	ut.name as "userType", 
	ut.description, 
	co.name as "courseName", 
	co.description as "courseDescr", 
	co.cost, 
	co.length, 
	co.requirements, 
	co.email as "courseEmail", 
	sc.name as "schoolName"
from ht.campaign ca, ht.course co, ht.school sc, ht.user u, ht.usertype ut
where ca.course_id = co.id AND co.school_id = sc.id AND ca.user_id = u.id AND ut.id = u.usertype_id;`
	var campaigns []Campaign

	err := s.DB.Select(&campaigns, query)
	if err != nil {
		panic(err)
	}
	fmt.Println(campaigns)

	c.JSON(http.StatusOK, campaigns)
}

func (s *ServerState) getCampaign(c *gin.Context) {
	//var cmpgn Campaign
	//if err := c.ShouldBindUri(&cmpgn); err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"status": err})
	//}
	//
	//if err := s.DB.Select(&cmpgn, "select "+string(cmpgn.Id)+" from campaign"); err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"status": err})
	//}
	//
	//c.JSON(http.StatusOK, gin.H{"campaign": cmpgn})
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
