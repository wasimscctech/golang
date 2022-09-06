package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	// Name     string
	// Price    int
	// Platform string
	// Password string
	// Tags     []string
	Name     string   `json:"coursename"` //all the names in the json should be in lower case hence we are renaming the values inside json data
	Price    int      //note: space should not be there inside the double quotes.
	Platform string   `json:"website"`
	Password string   `json:"-"`              //this minus/ dash sign will help us to hide the data which we dont want to show(i.e password)
	Tags     []string `json:"tags,omitempty"` //omitempty will omit the empty data and don't show "null" in json data
}

func EncodeJson() {
	myCourses := []course{ //a slice of type course that is a struct
		{"ReactJS Bootcamp", 299, "reactJS.in", "abc123", []string{"web-dev", "JS"}},
		{"Angular Bootcamp", 199, "Angular.in", "xyz123", []string{"full-stack", "JS"}},
		{"MERN Bootcamp", 299, "MERN.in", "uvw123", nil},
	}

	// package this data as JSON data

	// finalJson, err := json.Marshal(myCourses) //produces json
	finalJson, err := json.MarshalIndent(myCourses, "", "\t") //produces json

	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}

func DecodeJson() {
	jsonDatafromWeb := []byte(`
	{
		"coursename": "ReactJS Bootcamp",
		"Price": 299,
		"website": "reactJS.in",
		"tags": ["web-dev","JS"]
	}
	`)

	var myCourse course //course is the structure we created

	checkValid := json.Valid(jsonDatafromWeb) //will return the values in bool

	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDatafromWeb, &myCourse) //we are not sending the copy of a data hence used the reference
		fmt.Printf("%#v\n", myCourse)              //it will print the values in the original format which is Name for coursename, platform for website etc
	} else {
		fmt.Println("JSON was not Valid")
	}

	//some cases where you just want to add data to key value pair

	var myOnlineData map[string]interface{} //we have write interface{} because the key is a string that is confirmed but we don't know what type of value we are getting.
	json.Unmarshal(jsonDatafromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("key is %v and value is %v and the type is: %T\n", k, v, v)
	}
}

func main() {
	fmt.Println("Welcome to JSON conversion")
	// EncodeJson()
	DecodeJson()
}
