package som

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type UserData struct {
	Username string
	Password string
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	connStr := "user=yash password=yash dbname=codeforprogress sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &PostgresStore{
		db: db,
	}, nil
}

type Storage interface {
	CreateCourseInDB(*CourseData) error
	DeleteCourseInDB(*CourseData) error
	GetAllCourses() (*[]CourseData, error)
	GetCourseData(int) (*CourseData, error)
}

func (s *PostgresStore) CreateCourseInDB(newCourse *CourseData) error {
	query := `
	insert into "CourseData" (courseName, courseMetaData, courseImagePath, courseLink, ytcourseLink1, ytcourseLink2, ytcourseLink3, ytcourseLink4)
	values ( $1, $2, $3, $4, $5, $6, $7, $8)`

	_, err := s.db.Exec(query,
		newCourse.CourseName,
		newCourse.CourseMetaData,
		newCourse.CourseImagePath,
		newCourse.Link,
		newCourse.YtCourseLinks[0],
		newCourse.YtCourseLinks[1],
		newCourse.YtCourseLinks[2],
		newCourse.YtCourseLinks[3])

	if err != nil {
		return err
	}
	return nil
}

func (s *PostgresStore) DeleteCourseInDB(*CourseData) error {
	return nil
}

func (s *PostgresStore) GetAllCourses() (*[]CourseData, error) {
	var Courses []CourseData

	query := `select * from "CourseData"`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var course CourseData
		// funki.CourseData.YtCourseLinks[2] = 1
		var unNecessaryArray [4]string
		if err := rows.Scan(&course.CourseID, &course.CourseName, &course.CourseMetaData, &course.CourseImagePath, &course.Link, &unNecessaryArray[0], &unNecessaryArray[1], &unNecessaryArray[2], &unNecessaryArray[3]); err != nil {
			return nil, err
		}
		course.YtCourseLinks = append(course.YtCourseLinks, unNecessaryArray[0], unNecessaryArray[1], unNecessaryArray[2], unNecessaryArray[3])
		Courses = append(Courses, course)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &Courses, nil
}

func (s *PostgresStore) GetCourseData(id int) (*CourseData, error) {
	return nil, nil
}
