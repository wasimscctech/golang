package main

import "fmt"

func main() {
	//if else
	i := 3
	if i > 3 && i == 5 {
		fmt.Println("Hey!")
	} else if i < 4 {
		fmt.Println("Hi!")
	} else {
		fmt.Println("Bye!")
	}

	// // switch case
	// a := 90
	// switch a {
	// case 90:
	// 	fmt.Println("Awesome!")
	// 	fallthrough //this keyword is used when we want to print the next case after the true case. here case 90 will print if a == 90 and due to "fallthrough" case 60 will also get printed
	// case 60:
	// 	fmt.Println("good!")
	// case 40:
	// 	fmt.Println("OK!")
	// 	break // there is no break statement in switch case as it is implicitly added by go compiler. we still can use it in between two statement within same case
	// 	fmt.Println("keep working!")
	// default:
	// 	fmt.Printf("oops!")
	// }

	// switch i := 11; i { //we can initiaze here also
	// case 5:
	// 	fmt.Println("hello wasim")
	// case 6, 7, 8: //multiple cases can be write
	// 	fmt.Println("Hi Wasim")
	// case 10, 20, 30:
	// 	fmt.Println("Hey, Wasim!")
	// default:
	// 	fmt.Println("bye!")
	// }

	var x interface{} = "Wasim"
	switch x.(type) {
	case int:
		fmt.Println("integer")
	case float32:
		fmt.Println("float32")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("None")
	}
}
