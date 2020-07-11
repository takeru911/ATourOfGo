package main

import (
	"fmt"
	"runtime"
)
var Switch_test_var string = "test"
func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.", os)
	}
}

func TestPublicFunc(x int) int {
	return 0
}
