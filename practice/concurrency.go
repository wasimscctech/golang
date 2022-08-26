package main

import (
	"fmt"
	"time"
)

func main() {

	go sampleRoutine() //due the sleep time the execution of sampleRoutine function will run in background and “started main” will be executed first
	fmt.Println("Started Main")
	time.Sleep(10 * time.Second) //this will hold the execution for 10 sec and then next line will get printed.
	fmt.Println("Finished Main")
}

func sampleRoutine() {
	fmt.Println("Inside Sample Goroutine")
}
