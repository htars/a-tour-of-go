package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(b []byte) (n int, e error) {
	n, e = r.r.Read(b)
	for i, v := range b {
		switch {
		case (v >= 'A' && v < 'M') || (v > 'a' && v < 'm'):
			b[i] = v + 13
		case (v >= 'M' && v <= 'Z') || (v > 'm' && v <= 'z'):
			b[i] = v - 13
		default:
			continue
		}
	}
	return 
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
