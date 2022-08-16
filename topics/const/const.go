package main

import "fmt"

//emumarated const
const (
	i = iota //prints 0
	j = iota //prints 1 as both constants in a single group and iota will be increased from 0 to n
	k = iota //prints 2
	_        //if we want to skip the next number then we can use underscore
	w        //above underscore will skip number 3 and w will print 4
)
const (
	a = iota     //prints 0 as this is different const number group
	b = iota + 2 //prints 3 (1+2)
	x            //prints the next number after b, after the first constant we can just write the contant variable
	y
	z
)

func main() {
	fmt.Printf("%v, %v, %v, %v, %v, %v, %v, %v, %v", i, j, k, w, a, b, x, y, z)
}
