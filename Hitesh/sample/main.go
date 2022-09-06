package main

import (
	"fmt"
	"strconv"
)

var newfilename2 *string

func replaceunderscore(myFile string) {
	for index := 0; index < len(myFile); index++ {
		mynum, _ := strconv.Atoi(string(myFile[index]))
		// fmt.Println(mynum)

		if mynum != 0 {
			// fmt.Println("this is index: ", index)
			newFilename := myFile[:index] + myFile[index+1:]
			newfilename2 = &newFilename
		} else if string(myFile[index]) == "_" {
			// fmt.Println("this is index: ", index)
			newFilename := myFile[:index] + " " + myFile[index+1:]
			newfilename2 = &newFilename
		}
	}
	fmt.Println(*newfilename2)
}

// func replaceNumbers(){
// 	mynum, _ := strconv.Atoi(myStr)
// 	// fmt.Println(mynum)
// 	if mynum != 0 {
// 		myStr = ""
// 	}
// 	fmt.Println(myStr)
// }

func main() {
	replaceunderscore("Wasim_Shai7k7h.txt")
	// myFile := "W22"
	// var newfilename2 *string

	// for index := 0; index < len(myFile); index++ {
	// 	mynum, _ := strconv.Atoi(string(myFile[index]))
	// 	fmt.Println(mynum)
	// 	if mynum != 0 {
	// 		// fmt.Println("this is index: ", index)
	// 		newFilename := myFile[:index] + myFile[index+1:]
	// 		newfilename2 = &newFilename
	// 	}
	// }
	// fmt.Println(*newfilename2)
	// newString := strings.ReplaceAll(myFile, "_", " ")
}
