package api

type Campaign struct {
	// u.name, u.username , u.email, u.phone, u.country, ut.name as "userType", ut.description, co.name as "courseName",
	// co.description as "courseDescr", co.cost, co.length, co.requirements, co.email as "courseEmail", sc.name as "schoolName"
	Name              string
	Username          string
	Email             string
	Phone             string
	Country           string
	UserType          string `db:"userType"`
	Description       string
	CourseName        string `db:"courseName"`
	CourseDescription string `db:"courseDescr"`
	Cost              float32
	Length            string `db:"length"`
	Requirements      string
	CourseEmail       string `db:"courseEmail"`
	SchoolName        string `db:"schoolName"`
}

type InsertCampaign struct {
	Title       string
	Description string
	UserID      int
	CourseID    int
}

type AddFund struct {
	ID     int
	Amount int
}

type IndividualCampaign struct {
	/*
			        u.username ,
		        u.email,
		        u.phone,
		        u.country,
		        ut.name as "userType",
		        ut.description as "userTypeDescription",
		        ca.title as "campaignTitle",
		        ca.description as "campaignDescription",
		        ca.verified, co.name as "courseName",
		        co.description as "courseDescr",
		        co.cost, co.length,
		        co.requirements,
		        co.email as "courseEmail",
		        org.name as "organisationName"
	*/
	Name              string `db:"name" json:"name"`
	Username          string `json:"username"`
	Email             string `json:"email"`
	Phone             int    `json:"phone"`
	Country           string `json:"country"`
	UserType          string `db:"userType" json:"userType"`
	UserTypeDesc      string `db:"userTypeDescription" json:"userTypeDesc"`
	CampaignID        int    `json:"campaignID"`
	CampaignTitle     string `db:"campaignTitle" json:"campaignTitle"`
	CampaignDesc      string `db:"campaignDescription" json:"campaignDesc"`
	CampaignVerified  int    `db:"verified" json:"campaignVerified"`
	CourseName        string `db:"courseName" json:"courseName"`
	CourseDescription string `db:"courseDescr" json:"courseDesc"`
	CourseCost        int    `db:"cost" json:"courseCost"`
	CourseLength      string `db:"length" json:"courseLength"`
	CourseReq         string `db:"requirements" json:"courseReq"`
	CourseEmail       string `db:"courseEmail" json:"courseEmail"`
	OrgName           string `db:"organisationName" json:"organisatioName"`
}
