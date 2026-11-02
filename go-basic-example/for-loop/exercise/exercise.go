package main

import (
	"fmt"
	"math"
)

// x is the number it will be squared
// y is the number against it will be tested

// Square root playing. It does the square root by using Newton's method:
// y -= (y*y - x) / (2 * y)
func Sqrt_mine(x, y float64) float64 {
	// y= y - (function)
	y -= (y*y - x) / (2 * y)
	fmt.Printf("value:%v\n", y)
	return y
}

// Runs the squareroot I created and checks if the diff is smaller
// then 0.0001, if so increments the counter. Thrice the diff is smaller then
// 0.0001 it exits and thens print the diff between my implementation an the
// math.Sqrt() (native go)
func main() {
	x := float64(2)
	y := float64(1)

	for count := 0; count < 3; {
		y_old := y
		fmt.Printf("counts of diff < 0.001: %v\n", count)
		y = Sqrt_mine(x, y)
		if y_old-y < 0.001 {
			count += 1
		}
	}
	fmt.Println("Func sqrt returned a diference smaller then 0.001")
	fmt.Printf("%v is the square root from my func\n", y)
	real_sqrt_root := math.Sqrt(x)
	diff := real_sqrt_root - y
	fmt.Printf("My implementation is %f away from math.Sqrt()\n", diff)
}
