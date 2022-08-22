package main

import (
	"encoding/json"
	"net/http"
)

func CreateEmployee(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Contect-Type", "application/json")
	var emp Employee
	json.NewDecoder(r.Body).Decode(&emp) //decoding data which is sent in the body
	Database.Create(&emp)                //saving decoded data in database
	json.NewEncoder(w).Encode(emp)       //returning back in json form
}

func GetEmployees(w http.ResponseWriter, r *http.Request) {

}

func GetEmployeeBYID(w http.ResponseWriter, r *http.Request) {

}
