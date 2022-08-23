package main

import (
	"fmt"
	"sort"
)

func main() {
	var mySlice = []string{}                           //creating an empty slice
	var mySlice1 = []string{"apple", "mango", "date"}  //creating a slice
	fmt.Printf("the type of mySlice is %T\n", mySlice) //prints []string
	fmt.Printf("the value of mySlice is %v\n", mySlice)
	fmt.Printf("the type of mySlice1 is %T\n", mySlice1)
	fmt.Printf("the value of mySlice1 is %v\n", mySlice1)

	mySlice1 = append(mySlice1, "banana", "fig") //adding elements in slices
	fmt.Printf("the value of mySlice1 is %v\n", mySlice1)

	mySlice1 = append(mySlice1[1:4]) //new slice with indexed items
	fmt.Println(mySlice1)

	marks := make([]int, 4) //initializing a slice using make function
	marks[0] = 78
	marks[1] = 82
	marks[2] = 75
	marks[3] = 98
	// marks[4] = 99 //error message"panic: runtime error: index out of range [4] with length 4 " this will throw an error as the slice size mentioned is out of bounds. but using append method we can actually reallocate the memory of the slice
	fmt.Println(marks)
	marks = append(marks, 99, 59, 100) //here the memory of slice will be reallocated and adds up the elements at the end of the slice
	fmt.Println(marks)

	sort.Ints(marks) //sorts the int type slice in increasing order
	fmt.Println(marks)
	fmt.Println(sort.IntsAreSorted(marks)) //returns the bool value (true) as we have already sorted the slice

	subjects := []string{"english", "physics", "chemistry", "biology", "maths"}
	fmt.Println(subjects)
	index := 2
	subjects = append(subjects[:index], subjects[index+1:]...) //it will skip chemistry(element at index 2)
	// subjects = append(subjects[:index], subjects...) //this will create a slice with two enlish and physics elements
	fmt.Println(subjects)
}
