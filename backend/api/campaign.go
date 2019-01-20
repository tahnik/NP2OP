package api

type Campaign struct {
	// u.name, u.username , u.email, u.phone, u.country, ut.name as "userType", ut.description, co.name as "courseName",
	// co.description as "courseDescr", co.cost, co.length, co.requirements, co.email as "courseEmail", sc.name as "schoolName"
	Name string
	Username string
	Email string
	Phone string
	Country string
	UserType string
	Description string
	CourseName string
	CourseDescription string
	Cost float32
	Length int
	Requirements string
	CourseEmail string
	SchoolName string
}
