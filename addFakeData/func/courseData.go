package funki

type CourseData struct {
	CourseID       int    ` json:"courseId"`
	CourseName     string ` json:"courseName"`
	CourseMetaData string `json:"courseMetaData"`
	// CourseImage     Image  `bson:"courseImage" json:"courseImage"`
	CourseImagePath string   `json:"courseImageLink"`
	Link            string   `json:"courseLink"`
	YtCourseLinks   []string `json:"ytCourseLinks"`
}

func AddCourses(courses *[]CourseData) {
	*courses = append(*courses, CourseData{
		CourseName:      "Full stack developer",
		CourseMetaData:  "Welcome to web-development Course In this course we are providing all required resource to You Click Now",
		CourseImagePath: "./images/web-development-png-web-development-1100.png",
		Link:            "./courses/web/Web Development.html",
		YtCourseLinks:   []string{"fake1", "fake2", "fake3", "fake4"},
	}, CourseData{
		CourseName:      "Android Developer",
		CourseMetaData:  "Welcome to Android development Course In this course we are providing all required resource to You Click Now",
		CourseImagePath: "./images/android.png",
		Link:            "./courses/android/androiddev.html",
		YtCourseLinks:   []string{"fake1", "fake2", "fake3", "fake4"},
	}, CourseData{
		CourseName:      "Machine Learning",
		CourseMetaData:  "Welcome to Machine Learning Course In this course we are providing all required resource to You Click Now",
		CourseImagePath: "./images/machine.png",
		Link:            "./courses/Machine/machine.html",
		YtCourseLinks:   []string{"fake1", "fake2", "fake3", "fake4"},
	}, CourseData{
		CourseName:      "Cyber Security",
		CourseMetaData:  "Welcome to Cyber Security Course In this course we are providing all required resource to You Click Now",
		CourseImagePath: "./images/hacking.png",
		Link:            "./courses/CyberSecurity/Cybercon.html",
		YtCourseLinks:   []string{"fake1", "fake2", "fake3", "fake4"},
	}, CourseData{
		CourseName:      "Data Science",
		CourseMetaData:  "Welcome to Data Science Course In this course we are providing all required resource to You Click Now",
		CourseImagePath: "./images/ds.png",
		Link:            "./courses/ds/dscon.html",
		YtCourseLinks:   []string{"fake1", "fake2", "fake3", "fake4"},
	}, CourseData{
		CourseName:      "Data Structures & Algorithms",
		CourseMetaData:  "Welcome to Data Structures and Algorithm Course In this course we are providing all required resource to You Click Now",
		CourseImagePath: "./images/dsa.png",
		Link:            "./courses/DSA/dsa.html",
		YtCourseLinks:   []string{"fake1", "fake2", "fake3", "fake4"},
	}, CourseData{
		CourseName:      "Business Foundation",
		CourseMetaData:  "Welcome to Business Foundation Course In this course we are providing all required resource to You Click Now",
		CourseImagePath: "./images/oracle.png",
		Link:            "./courses/Busl/busscon.html",
		YtCourseLinks:   []string{"fake1", "fake2", "fake3", "fake4"},
	}, CourseData{
		CourseName:      "DevOps for SDLC",
		CourseMetaData:  "Welcome to DevOps Course In this course we are providing all required resource to You Click Now",
		CourseImagePath: "./images/cloud-data.png",
		Link:            "./courses/devops/dev.html",
		YtCourseLinks:   []string{"fake1", "fake2", "fake3", "fake4"},
	}, CourseData{
		CourseName:      "DevOps for SDLC",
		CourseMetaData:  "Welcome to DevOps Course In this course we are providing all required resource to You Click Now",
		CourseImagePath: "./images/cloud-data.png",
		Link:            "./courses/devops/dev.html",
		YtCourseLinks:   []string{"fake1", "fake2", "fake3", "fake4"},
	})
}
