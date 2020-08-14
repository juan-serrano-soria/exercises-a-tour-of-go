package main

import "golang.org/x/tour/pic"

// Pic returns a two dimensional slice of size dy x dx
// filled with values calculated using a concrete function
func Pic(dx, dy int) [][]uint8 {
	
	// Allocate two dimensional slice and fill it with values calculated using xXORy
	s := make([][]uint8, dy)
	for i := range s {
		s[i] = make([]uint8, dx)
		for j := range s[i] {
			s[i][j] = uint8(i^j)
		}
	}
	
	return s
}

func main() {
	pic.Show(Pic)
}
