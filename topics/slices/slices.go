package main

import "fmt"

func main() {
	//array
	var arr1 [4]int = [4]int{4, 3, 2, 1} //array fixed sized
	arr2 := [4]int{4, 3, 2, 1}           //array fixed size
	arr3 := [...]int{4, 3, 2, 1, 10}     //array dynamic size

	// slice
	arr4 := []int{1, 4, 6, 7, 8, 9}    //this is the 1st major difference between array and slice
	var arr5 []int = make([]int, 2, 4) //another way to declare + define a slice, this is an empty slice, "2" represent the length(size) and "4" represent the capacity.

	fmt.Printf("this is an array %v\n", arr1)
	fmt.Printf("this is an array %v\n", arr2)
	fmt.Printf("this is an array %v\n", arr3)
	fmt.Printf("this is a slice %v\n", arr4)
	fmt.Printf("this is a slice %v\n", arr5)

	// length and capacity of a slice
	fmt.Printf("this is legnth of arr4: %v\n", len(arr4)) //by default the length and capacity would be same
	fmt.Printf("this is capacity of arr4: %v\n", cap(arr4))
	fmt.Printf("this is legnth of arr5: %v\n", len(arr5))
	fmt.Printf("this is capacity of arr5: %v\n", cap(arr5))

	fmt.Println(arr4[0])
	arr4[0] = 5
	fmt.Println(arr4[0])

	arr6 := arr4 //in array if we assign an array's value with other array then all the elements get copied, but in slices the memory location of other slice get copied. this one is the 2nd differnce.
	fmt.Println(arr6)

	arr6[5] = 18 //value of arr4 will also get changed and both the slice will be same.
	fmt.Println(arr4)
	fmt.Println(arr6)

	// adding extra element at the end
	myslice := []int{1, 2, 3, 4}
	myslice2 := append(myslice, 10)
	fmt.Println(myslice2)
	myslice3 := append(myslice, 10, 50)
	fmt.Println(myslice3)
	// myslice4 := append(myslice2, myslice3) //this will throw an error that is why we need to put 3 dots as below line
	myslice4 := append(myslice2, myslice3...) //away to add the values of other slices into the different one.
	fmt.Println(myslice4)

}
