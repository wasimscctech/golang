package main

import "fmt"

func main() {
	days := []string{"monday", "tuesday", "wednesday", "friday", "sunday"}
	fmt.Println(days)
	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for index, day := range days { //index will start from zero to n and day will be the elements inside the slice
	// 	fmt.Printf("Index is %v and the day is %v\n", index, day)
	// }

	// for _, day := range days {
	// 	fmt.Printf("the day is %v\n", day)
	// }

	// similar to while loop
	i := 1
	for i < 10 {
		// break
		// if i == 5 {
		// 	break
		// }
		// continue
		if i == 5 {
			i++ //if we don't do the increment the loop will stuck in 4 and it will keep on running
			continue
		}
		// goto
		if i == 6 {
			goto myLabel //this will break the loop and prints the below label statement, it won't print 6.
		}

		fmt.Println(i)
		i++ //preincrement will be an error in golang
	}
myLabel:
	fmt.Println("This is a goto statement")
}
