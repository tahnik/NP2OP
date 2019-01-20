package api

type Campaign struct {
	Id            int    `db:"id"`
	TotalReceived string `db:"total_received"`
	UserId        string `db:"user_id"`
	CourseId      string `db:"course_id"`
}
