package main

import (
	"fmt"
	"time"

)

func main() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	fmt.Println(time.Saturday)
	fmt.Println(today-1)
}
