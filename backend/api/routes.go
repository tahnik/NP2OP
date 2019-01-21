package api

import (
	"fmt"
	"net/http"
	"strconv"

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
		Users.GET("/:email", s.getUser) //localhost:8080/user/sdakjfbdshfbsdihvb
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
		Campaigns.PUT("/:id/funding", s.updateCampaignFunding)
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
			return
		}
		if post.Description == "" {
			c.JSON(400, gin.H{"status": "Post.Description cannot be blank"})
			return
		}
		if post.Description == "" || len(post.Description) < 100 {
			c.JSON(400, gin.H{"status": "Post.Description cannot be less than 100 characters"})
			return
		}
		if post.Cost == 0 || post.Cost < 0 {
			c.JSON(400, gin.H{"status": "post.Goal is less than 0 or a negative amount"})
			return
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
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": fmt.Sprintf("User inserted successfully, ID is %d", id)})
}

//Placeholder function demonstrating the gin grouping
func (s *ServerState) getPosts(c *gin.Context) {
	query := `select u.name, u.username , p.description, p.timestamp, p.total_campaign_snapshot, co.name, co.cost, sc.name
	from ht.post p, ht.campaign ca, ht.course co, ht.organisation sc, ht.user u
	where p.campaign_id = ca.id AND ca.course_id = co.id AND co.organisation_id = sc.id AND ca.user_id = u.id;`

	//curl --header "Content-Type: application/json" --request GET http://localhost:8080/posts/

	var posts []Post

	err := s.DB.Select(&posts, query)
	if err != nil {
		c.JSON(500, gin.H{"status": err})
		return
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
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": u})
}

//curl --request GET "http://localhost:8080/user/visa@visa.com"
func (s *ServerState) getUser(c *gin.Context) {
	var u []User
	q := fmt.Sprintf("SELECT * FROM user WHERE email='%s'", c.Param("email"))
	fmt.Println(q)
	if err := s.DB.Select(&u, q); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"ok": false, "status": err.Error()})
		return
	}
	if len(u) == 0 {
		c.JSON(http.StatusOK, gin.H{"ok": false, "status": "No user by that name"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"ok": true})
}

// curl --request POST --data '{"name":"james","email":"j@j.com", "Phone":"07875","usertype_id":0, "country":"UK"}' "http://localhost:8080/user/"
func (s *ServerState) addUser(c *gin.Context) {
	var u User

	err := c.BindJSON(&u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	phoneInt, err := strconv.Atoi(u.Phone)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	q := "INSERT INTO `user` (`password`, `name`, `email`, `phone`, `usertype_id`, `country`, `username`) VALUES ('', ?, ?, ?, ?, ?, '')"
	_, err = s.DB.Exec(q, u.Name, u.Email, phoneInt, u.UserTypeId, u.Country)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
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
			return
		}
		if ins.Description == "" {
			c.JSON(400, gin.H{"status": "ins.Description cannot be blank"})
			return
		}
		if ins.CourseID == 0 {
			c.JSON(400, gin.H{"status": "ins.Description cannot be zero"})
			return
		}
		if ins.Title == "" {
			c.JSON(400, gin.H{"status": "ins.Title is blank"})
			return
		}
	} else {
		c.JSON(400, gin.H{"status": fmt.Sprintf("Failed to bind campaing.insert, error: %s", err.Error())})
		return
	}

	res := s.DB.MustExec(query, ins.Title, ins.Description, 0, ins.UserID, ins.CourseID)
	id, err := res.LastInsertId()
	if err != nil {
		c.JSON(500, gin.H{"status": "Failure returning last inserted ID"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": fmt.Sprintf("Campaign inserted successfully, ID is %d", id)})
}

func (s *ServerState) getCampaigns(c *gin.Context) {
	query := `
select 
	u.name, 
	ca.title,
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
from ht.campaign ca, ht.course co, ht.organisation sc, ht.user u, ht.usertype ut
where ca.course_id = co.id AND co.organisation_id = sc.id AND ca.user_id = u.id AND ut.id = u.usertype_id;`
	var campaigns []Campaign

	err := s.DB.Select(&campaigns, query)
	if err != nil {
		panic(err)
	}
	fmt.Println(campaigns)

	c.JSON(http.StatusOK, campaigns)
}

func (s *ServerState) getCampaign(c *gin.Context) {
	q := fmt.Sprintf("select %s %s %s from %s where %s",
		"u.name,u.username, u.email,u.phone,u.country,ut.name as `userType`,ut.description as `userTypeDescription`,",
		"ca.title as `campaignTitle`, ca.description as `campaignDescription`, ca.verified, co.name as `courseName`,co.description as `courseDescr`,",
		"co.cost, co.length, co.requirements, co.email as `courseEmail`, org.name as `organisationName`",
		"ht.campaign ca, ht.course co, ht.organisation org, ht.user u, ht.usertype ut",
		"ca.id = ? AND ca.course_id = co.id AND co.organisation_id = org.id AND ca.user_id = u.id AND ut.id = u.usertype_id",
	)
	fmt.Println(q)

	var camp IndividualCampaign
	str := c.Param("id")

	id, err := strconv.Atoi(str)

	camp.CampaignID = id
	//Validate
	if camp.CampaignID == 0 {
		c.JSON(400, gin.H{"status": "ins.ID cannot be zero"})
		return
	}

	err = s.DB.Get(&camp, q, string(id))
	if err != nil {
		c.JSON(500, gin.H{"status": fmt.Sprintf("Failed to get getCampaign, error: %s", err.Error())})
		return
	}

	c.JSON(http.StatusOK, camp)
}

func (s *ServerState) updateCampaign(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func (s *ServerState) updateCampaignFunding(c *gin.Context) {
	var add AddFund

	/*
		curl --header "Content-Type: application/json" --request PUT --data \
		    '{"amount": 100, "id": 4}' \
		     http://localhost:8080/campaigns/5/funding
	*/

	if err := c.ShouldBind(&add); err == nil {
		//Validate
		if add.ID == 0 {
			c.JSON(400, gin.H{"status": "add.ID cannot be zero"})
			return
		}
		if add.Amount == 0 || add.Amount < 0 {
			c.JSON(400, gin.H{"status": "add.amount cannot be zero or less than zero"})
			return
		}
	} else {
		c.JSON(400, gin.H{"status": fmt.Sprintf("Failed to bind updateCampaign.add, error: %s", err.Error())})
		return
	}

	query := `UPDATE campaign SET total_received = total_received + ? WHERE id = ?`
	_, err := s.DB.Exec(query, add.Amount, add.ID)
	if err != nil {
		c.JSON(500, gin.H{"status": fmt.Sprintf("Failed to bind updateCampaignFunding, error: %s", err.Error())})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func (s *ServerState) approveCampaign(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
