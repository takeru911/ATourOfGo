package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	var z float64 = 1
	for math.Abs(z*z-x) > 0.0000000000001 {
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(15))
	fmt.Println(math.Sqrt(15))
}
