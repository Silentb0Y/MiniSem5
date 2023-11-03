import "../home.css"
// import "../all.css"
import { Nav } from "../components/nav"
import { useEffect, useState } from "react"

interface ApiData  {
	Username   :string       
	Password   :string       
	CourseId   :number       
	Course     :CourseData   
	Courselist :CourseData[]
}

interface CourseData {
	courseId       :number   
	courseName     :string 
	courseMetaData :string 
	courseImageLink :string  
	courseLink            :string   
	ytCourseLinks   :string[]
}


interface CourseInfo {
    Courseid: number
    CourseName: string
    CourseMetaData: string
    CourseImage: string
    CoursePageLink: string
}

export function Home() {
    const [courseList, setCourseList] = useState<CourseInfo[]>([])
    const [courseListola, setCourseListola] = useState<CourseData[]>([])
    const [wasetbull, setbull] = useState(true)
    const newcourse: CourseInfo = {
        Courseid:1,
        CourseName:"Full stack developer",
        CourseMetaData:"Welcome to web-development Course In this course we are providing all required resource to You",
        CourseImage:"./images/web-development-png-web-development-1100.png",
        CoursePageLink:"./courses/web/Web Development.html"
    }
    useEffect(() => {
            // setCourseList([...courseList, newcourse])
            fetch("http://localhost:7002/getCourses", {
                method: "GET",
              })
                .then((res) => res.json())
                .then((data) => {
                    console.log(data)
                    setCourseListola(data.courseList)
                })
                .catch((error) => console.log("ERROR:", error));
    },[])
    useEffect(() => {
        for (const course of courseListola) {
            const newcourse2: CourseInfo = {
                Courseid: course.courseId,
                CourseName: course.courseName,
                CourseMetaData: course.courseMetaData,
                CourseImage: course.courseImageLink,
                CoursePageLink: course.courseLink 
            }
            setCourseList([...courseList, newcourse2])
        }
    }, [courseListola])

    const courseListDiv = courseList.map(course =>
        <div className="training-item" key={course.Courseid}>
            <img src={course.CourseImage} alt="" />
            <h3>{course.CourseName}</h3>
            <p>{course.CourseMetaData}<small>Click Now</small></p>
            <a href={course.CoursePageLink} target="blank"><button>Start</button></a>
        </div>
    )
    return (
        // <div>
            <div className="container"> 
            {courseListDiv}
            </div>
        // </div>

    )
}