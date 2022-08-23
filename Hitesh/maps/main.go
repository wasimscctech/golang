package main

import "fmt"

func main() {
	languages := make(map[string]string) //defining a map
	fmt.Printf("the type is %T\n", languages)
	languages["PY"] = "python" //adding the key value pairs
	languages["GO"] = "google"
	languages["JS"] = "JavaScript"
	fmt.Println("list of all languages", languages)
	fmt.Println("JS shorts for:  ", languages["JS"])

	delete(languages, "GO") //deleting a key
	fmt.Println("list of all languages", languages)

	// loops
	for key, value := range languages {
		fmt.Printf("For key %v, the value is: %v\n", key, value)
	}
}
