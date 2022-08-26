package main

import (
	"fmt"
	"net/url"
)

const myUrls string = "https://lco.dev:3000/learn?coursename=reactJS&paymentid=ghdawjdhaj2313"

func main() {
	fmt.Println("handling urls")
	fmt.Println(myUrls)

	//parsing url:  (URL Parsing. The URL parsing functions focus on splitting a URL string into its components, or on combining URL components into a URL string) i.e extracting the info
	result, err := url.Parse(myUrls)

	if err != nil {
		panic(err)
	}

	fmt.Println(result.Scheme)   //prints https
	fmt.Println(result.Host)     //prints lco.dev:3000
	fmt.Println(result.Path)     //prints  /learn
	fmt.Println(result.Port())   //prints 3000
	fmt.Println(result.RawQuery) //prints coursename=reactJS&paymentid=ghdawjdhaj2313

	queryPara := result.Query()                                   //query function stores all the query parameters of the url into a better format.
	fmt.Printf("the type of query parameter is: %T\n", queryPara) //prints: "url.Values" (i.e the data is in key value pairs and that is a dictionary)

	fmt.Println(queryPara["coursename"]) //coursename is the key in our URL and the value is reactJS
	fmt.Println(queryPara["paymentid"])  //& in URL is a comma and query in are url is in the form of a map so we can get the data accordingly

	for _, val := range queryPara {
		fmt.Println("parameter is: ", val) //it will print the query parameters values only.
	}
	fmt.Println(queryPara) //it will print "map[coursename:[reactJS] paymentid:[ghdawjdhaj2313]]"

	// incase we want to create a URL using some different parts of url
	partsOfUrl := &url.URL{ //& is must url is the package and URL is the method of that package.
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "tutcss",
		RawPath: "user=hitesh",
	}
	fmt.Printf("the type is %T\n", partsOfUrl) //the type is *url.URL, we need a string hence need to convert using the below syntax
	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)
	fmt.Printf("the type %T", anotherUrl) //the type is a string
}
