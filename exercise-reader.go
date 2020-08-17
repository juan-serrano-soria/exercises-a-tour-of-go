package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// Read implements a Reader that emits an infinite stream of
// the ASCII character 'A'.
func (r MyReader) Read(b []byte) (int, error) {
	// ASCII for A is 65
	for i := range(b) {
		b[i] = 65
	}
	return len(b), nil
}

func main() {
	reader.Validate(MyReader{})
}
