package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

// rot13 applies rot13 to b
func rot13(b byte) byte {
	switch {
	// between A and M or between a and m
	case (b > 64 && b < 78) || (b > 96 && b < 110):
		b += 13
	// between N and Z or between n and z
	case (b > 77 && b < 91) || (b > 109 && b < 123):
		b -= 13
	}
	return b
}

// Read implements a Reader that uses io.Reader and applies a rot13 to it
func (rr *rot13Reader) Read(b []byte) (int, error) {
	n, err := rr.r.Read(b)
	for i := range b {
		b[i] = rot13(b[i])
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
