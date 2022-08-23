package main

import (
	"fmt"
	"time"
)

func main() {
	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006 monday"))

	createdDate := time.Date(2022, time.February, 11, 12, 8, 0, 0, time.UTC)
	fmt.Println(createdDate)
}
