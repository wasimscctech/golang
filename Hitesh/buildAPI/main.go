package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// model for course - file
type Course struct {
	CourseID    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	FullName string `json:"fullname"`
	Website  string `json:"website"`
}

// fake DB
var courses []Course

// middleware/helper
func (c *Course) IsEmpty() bool {
	// return c.CourseID == "" && c.CourseName == ""
	return c.CourseName == ""
}

func main() {

}

//controllers - file

// serve home router

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to my API</h1>"))
}

func getAllCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(courses)

}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("content-type", "application/json")

	// grab id from request
	params := mux.Vars(r)

	// loop through course, find matching id and return the response

	for _, course := range courses {
		if course.CourseID == params["id"] {
			json.NewEncoder(w).Encode(course)
		}
		return
	}
	json.NewEncoder(w).Encode("no course found with given ID")

}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("content-type", "application/json")

	//what if: body of the respose is empty

	if r.Body == nil {
		json.NewEncoder(w).Encode("please send some data")
	}

	// what if: the response is empty json i.e {}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside the JSON")
		return
	}

	//generate unique id, string
	// append course into courses

	rand.Seed(time.now().UnixNano())
	course.CourseID = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return

}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("update one course")
	w.Header().Set("content-type", "application/json")
	// first- grab id from request
	params := mux.Vars(r)

	// loop, id , remove, add with my id
	for index, course := range courses {
		if course.CourseID == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseID = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	//todo: send a response when ID is not found
	
}
