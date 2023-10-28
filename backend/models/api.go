package model

type ApiData struct {
	Username   string   `json:"username"`
	Password   string   `json:"password"`
	CourseId   int      `json:"courseId"`
	Course     Course   `json:"course"`
	Courselist []Course `json:"courseList"`
}
