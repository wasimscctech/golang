package main

import "fmt"

func main() {
	var arr1 [3]int                   //declaration and default value will be [0 0 0]
	var arr2 [3]int = [3]int{1, 2, 4} //declaration + initialization
	arr3 := [3]int{5, 6, 7}
	staticArr := [3]int{0, 0, 7}
	dynamicArr := [...]int{5, 6, 7, 8, 9, 0, 1, 10} //we can change the size by using triple dots inside []
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(staticArr)
	fmt.Println(dynamicArr)

	myArr := [...]int{5, 6, 7, 8, 9, 0, 1, 10, 1, 2, 4, 5}
	fmt.Println(len(myArr)) //will print the size of array
	fmt.Println(myArr[4])   //will print the value at 4th index (at 5th position).
	myArr[5] = 11           //we can also change the value at a perticular position
	fmt.Println(myArr)
	myArr2 := myArr
	myArr2[5] = 7
	fmt.Println(myArr2)
	fmt.Println(myArr2[3:])
	fmt.Println(myArr2[:3])
	fmt.Println(myArr2[3:6])

	//multidimensional array
	matrix := [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(matrix[1][1]) //prints 5
}
