package som

type ApiData struct {
	Username   string       `json:"username"`
	Password   string       `json:"password"`
	CourseId   int          `json:"courseId"`
	Course     CourseData   `json:"course"`
	Courselist []CourseData `json:"courseList"`
}

type CourseData struct {
	CourseID       int    ` json:"courseId"`
	CourseName     string ` json:"courseName"`
	CourseMetaData string `json:"courseMetaData"`
	// CourseImage     Image  `bson:"courseImage" json:"courseImage"`
	CourseImagePath string   `json:"courseImageLink"`
	Link            string   `json:"courseLink"`
	YtCourseLinks   []string `json:"ytCourseLinks"`
}
