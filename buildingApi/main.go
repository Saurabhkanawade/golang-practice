package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// main function

func main() {
	fmt.Println("Welcome to the goland building api")

	r := mux.NewRouter()

	// getAll controller
	r.HandleFunc("/courses", getAllCourse).Methods("GET")

	// getById controller
	r.HandleFunc("/courses/{id}", getOneCourse).Methods("GET")

	//  create course
	r.HandleFunc("/courses", createCourse).Methods("POST")

	// update course 
	r.HandleFunc("/courses",updateOneCourse).Methods("PUT")

	//delete course
	r.HandleFunc("/courses/{id}",deleteOneCourse).Methods("DELETE")

	http.ListenAndServe(":9000", r)
	//

}

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"courseprice"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake db
var courses []Course

//middleware , helper-file

func (c *Course) IsEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h2>Hello welcome to the golang route</h2>"))
}

func getAllCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get All Courses")
	w.Header().Set("Content-Tyep", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get One Course")
	w.Header().Set("Content-Type", "application/json")

	//grab id from request
	params := mux.Vars(r)

	//loops through course , find matching id and return the response

	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("No Course Found with given courseId ")
	return
}

// second create course method

func createCourse(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	// if body is empty

	if r.Body == nil {

		json.NewEncoder(w).Encode("Please some data")

	}

	var course Course

	_ = json.NewDecoder(r.Body).Decode(&course)

	if course.IsEmpty() {

		json.NewEncoder(w).Encode("No data Inside json")

	}

	rand.Seed(time.Now().UnixMicro())

	course.CourseId = strconv.Itoa(rand.Intn(100))

	courses = append(courses, course)

	json.NewEncoder(w).Encode(course)

	return

}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("update the course")
	w.Header().Set("Content-Type", "application/json")

	//first grab id from req

	params := mux.Vars(r)

	//loops , id , remove , add with my id

	for index, course := range courses {

		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)

			var course Course
			json.NewDecoder(r.Body).Decode(&course)
			return
		}

	}
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request){
	log.Println("Delete one course")
	w.Header().Set("Content-Type","application/json")

	params :=mux.Vars(r)

	for index,course:=range courses {

		if course.CourseId== params["id"]{
			courses=append(courses[:index],courses[index+1:]...)
			break
		}
	}
}