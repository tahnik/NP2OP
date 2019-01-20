package api

type Campaign struct {
	// u.name, u.username , u.email, u.phone, u.country, ut.name as "userType", ut.description, co.name as "courseName",
	// co.description as "courseDescr", co.cost, co.length, co.requirements, co.email as "courseEmail", sc.name as "schoolName"
	Name              string
	Title             string
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
