package main

import "fmt"

func main() {
	var myArr [4]string                   //array declaration
	fmt.Println("the empty array", myArr) // this will print 4 spaces(empty array)

	myArr[0] = "mango"
	myArr[1] = "dates"
	myArr[3] = "figs"
	fmt.Println("the array", myArr)              //prints [mango dates  figs] includes the space at 2nd index
	fmt.Println("the size of array", len(myArr)) //this will always be the mentioned size even if we put only one value

	var myArr2 = [4]string{"marriage", "home", "vehicle"}
	fmt.Println("my array 2 is ", myArr2)
	fmt.Println("my array2 size is ", len(myArr2))
	fmt.Printf("the type of myarr2 is: %T", myArr2)

}
