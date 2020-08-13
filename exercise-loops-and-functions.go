package main

import (
	"fmt"
	"math"
)

// Abs returns the absolute value of x
func Abs(x float64) float64 {
	if x < 0 {
		x = -x
	}
	return x
}

// Sqrt calculates the root of a given x number using Newton's method
// We use two vars, z and y, to show one iteration and the previous, so we can stop the algorithm
func Sqrt(x float64) float64 {
	
	z := 1.0
	var y float64
	iters := 0
	
	sens := 1e-20
	
	fmt.Println("Printing iterations of Newton's Method:")
	
	for {
		y -= (z*z - x) / (2*z)
		
		//fmt.Println(Abs(z -y))
		if Abs(z - y) < sens  {
			break
		}
		
		iters++
		z = y
		fmt.Println(z)
	}
	
	fmt.Println("Number of iterations:", iters)
	return z
}

func main() {
	x := 4.0
	fmt.Println("Our funtcion's value:", Sqrt(x))
	fmt.Println("fmt default function value:", math.Sqrt(x))
}
