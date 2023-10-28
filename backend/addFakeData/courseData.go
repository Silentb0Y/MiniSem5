package main

func main() {
	var courses []CourseData
	// Print the list of courses
	addCourses(&courses)
	// for _, course := range courses {
	// 	fmt.Println(course.CourseName)
	// 	fmt.Println(course.CourseImagePath)
	// 	fmt.Println(course.Link)
	// 	fmt.Println()
	// }
}

type CourseData struct {
	CourseID       int    `bson:"_id,omitempty" json:"courseId"`
	CourseName     string `bson:"courseName" json:"courseName"`
	CourseMetaData string `bson:"courseMetaData" json:"courseMetaData"`
	// CourseImage     Image  `bson:"courseImage" json:"courseImage"`
	CourseImagePath string `json:"courseImageLink"`
	Link            string `json:"courseLink"`
}

func addCourses(courses *[]CourseData) {
	*courses = append(*courses, CourseData{
		CourseName:      "Full stack developer",
		CourseMetaData:  "Welcome to web-development Course In this course we are providing all required resource to You Click Now",
		CourseImagePath: "./images/web-development-png-web-development-1100.png",
		Link:            "./courses/web/Web Development.html",
	}, CourseData{
		CourseName:      "Android Developer",
		CourseMetaData:  "Welcome to Android development Course In this course we are providing all required resource to You Click Now",
		CourseImagePath: "./images/android.png",
		Link:            "./courses/android/androiddev.html",
	}, CourseData{
		CourseName:      "Machine Learning",
		CourseMetaData:  "Welcome to Machine Learning Course In this course we are providing all required resource to You Click Now",
		CourseImagePath: "./images/machine.png",
		Link:            "./courses/Machine/machine.html",
	}, CourseData{
		CourseName:      "Cyber Security",
		CourseMetaData:  "Welcome to Cyber Security Course In this course we are providing all required resource to You Click Now",
		CourseImagePath: "./images/hacking.png",
		Link:            "./courses/CyberSecurity/Cybercon.html",
	}, CourseData{
		CourseName:      "Data Science",
		CourseMetaData:  "Welcome to Data Science Course In this course we are providing all required resource to You Click Now",
		CourseImagePath: "./images/ds.png",
		Link:            "./courses/ds/dscon.html",
	}, CourseData{
		CourseName:      "Data Structures & Algorithms",
		CourseMetaData:  "Welcome to Data Structures and Algorithm Course In this course we are providing all required resource to You Click Now",
		CourseImagePath: "./images/dsa.png",
		Link:            "./courses/DSA/dsa.html",
	}, CourseData{
		CourseName:      "Business Foundation",
		CourseMetaData:  "Welcome to Business Foundation Course In this course we are providing all required resource to You Click Now",
		CourseImagePath: "./images/oracle.png",
		Link:            "./courses/Busl/busscon.html",
	}, CourseData{
		CourseName:      "DevOps for SDLC",
		CourseMetaData:  "Welcome to DevOps Course In this course we are providing all required resource to You Click Now",
		CourseImagePath: "./images/cloud-data.png",
		Link:            "./courses/devops/dev.html",
	}, CourseData{
		CourseName:      "DevOps for SDLC",
		CourseMetaData:  "Welcome to DevOps Course In this course we are providing all required resource to You Click Now",
		CourseImagePath: "./images/cloud-data.png",
		Link:            "./courses/devops/dev.html",
	})
}
