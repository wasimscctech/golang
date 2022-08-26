package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("this is file handling")
	content := "This will be our content in myfiles"

	file, err := os.Create("./myfiles.txt") //this will create a file
	// if err != nil {
	// 	panic(err)
	// }
	checkNilErr(err)

	length, err := io.WriteString(file, content) //io package will be used to write on a file.
	checkNilErr(err)

	fmt.Println("The length is: ", length)
	defer file.Close()
	readFile("./myfiles.txt")
}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename) //ioutil is used to read and manipulate the file
	checkNilErr(err)
	// fmt.Println("the text data inside the file is: ", databyte) //this will print the binary data [116 104 105 115 32 110 101 101 100 115 32 116 111 32 103 111 32 105 110 32 97 32 102 105 108 101 32 45 32 109 121 102 105 108 101 115] but we need a string
	fmt.Println("the text data inside the file is: \n", string(databyte))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
