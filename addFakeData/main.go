package main

import (
	"database/sql"
	"fmt"
	"log"

	funki "fakedata/func"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateUserInDB(*funki.UserData) error
	DeleteUserFromDB(*funki.UserData) error

	CreateCourseInDB(*funki.CourseData) error
	DeleteCourseInDB(*funki.CourseData) error

	CheckUservalidityAndGetUser(username string, password string) (*funki.UserData, error)
	GetAllCourses() (*[]funki.CourseData, error)
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

func (s *PostgresStore) Init() error {
	fmt.Println("init function was executed")
	if err := s.CreateUserTable(); err != nil {
		return err
	}
	if err := s.CreateCourseTable(); err != nil {
		return err
	}
	return nil
}

func (s *PostgresStore) CreateUserTable() error {
	query := `
		create table if not exists "UserData"(
		username varchar(225),
		password varchar(225)
	)`
	_, err := s.db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func (s *PostgresStore) CreateCourseTable() error {
	query := `
		create table if not exists "CourseData"(
			courseID serial primary key,
			courseName varchar(225),
			courseMetaData text,
			courseImagePath text,
			courseLink varchar(255),

			ytcourseLink1 varchar(255),
			ytcourseLink2 varchar(255),
			ytcourseLink3 varchar(255),
			ytcourseLink4 varchar(255)
	)`
	// courseLinks	Serial primary key
	_, err := s.db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func (s *PostgresStore) CreateUserInDB(newUser *funki.UserData) error {
	query := `
	insert into "UserData" (username, password)
	values ( $1, $2)`

	_, err := s.db.Exec(query, newUser.Username, newUser.Password)

	if err != nil {
		return err
	}
	return nil
}

func DeleteUserFromDB(*funki.UserData) error {
	return nil
}

func (s *PostgresStore) CreateCourseInDB(newCourse *funki.CourseData) error {
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

func (s *PostgresStore) DeleteCourseInDB(*funki.CourseData) error {
	return nil

}

func (s *PostgresStore) CheckUservalidityAndGetUser(username string, password string) (bool, error) {
	var userFound funki.UserData
	query := `
	select * from "UserData"
	where username=$1 `

	err := s.db.QueryRow(query, username).Scan(&userFound.Username, &userFound.Password)
	if err != nil {
		return false, err
	}
	if password == userFound.Password {
		return true, nil
	}
	return false, fmt.Errorf("wrong password")

}

func (s *PostgresStore) GetAllCourses() (*[]funki.CourseData, error) {
	var Courses []funki.CourseData

	query := `select * from "CourseData"`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var course funki.CourseData
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

func main() {

	store, err := NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", store)
	err1 := store.Init()
	if err1 != nil {
		log.Fatal(err1)
	}
	defer store.db.Close()
	var courses []funki.CourseData
	var users []funki.UserData
	funki.AddUsers(&users)
	funki.AddCourses(&courses)

	for _, user := range users {
		store.CreateUserInDB(&user)
	}
	for _, course := range courses {
		store.CreateCourseInDB(&course)
	}
	fmt.Println(store.GetAllCourses())
	// fmt.Println(store.CheckUservalidityAndGetUser("yash", "yash"))
}
