package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("Cannot Sqrt negative number: %.2f", e)
}

// Abs returns the absolute value of x
func Abs(x float64) float64 {
	if x < 0 {
		x = -x
	}
	return x
}

// Sqrt calculates the root of a given x number using Newton's method
// We use two vars, z and y, to show one iteration and the previous, so we can stop the algorithm
// If x is negative, it returns an error
func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	
	z := 1.0
	var y float64
	iters := 0
	sens := 1e-10
	fmt.Println("Printing iterations of Newton's Method:")
	for {
		y -= (z*z - x) / (2*z)
		if Abs(z - y) < sens  {
			break
		}
		iters++
		z = y
		fmt.Println(z)
	}
	
	fmt.Println("Number of iterations:", iters)
	return z, nil
}


func main() {
	if v, err := Sqrt(2); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(v)
	}
	
	if v, err := Sqrt(-2); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(v)
	}

}
