package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

// Rot13 applies rot13 to b
func Rot13(b byte) byte {
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

func (rr *rot13Reader) Read(b []byte) (int, error) {
	// TODO
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
