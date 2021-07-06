package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot *rot13Reader) Read(b []byte) (int, error) {
	n, err := rot.r.Read(b)
	for i, c := range b {
		if c >= 65 && c <= 90 { // A - 65, Z- 90
			tmp := c + 13
			if tmp > 90 {
				tmp = tmp%90 + 64
			}
			b[i] = tmp
		} else if c >= 97 && c <= 122 { // a - 97, z - 122
			tmp := c + 13
			if tmp > 122 {
				tmp = tmp%122 + 96
			}
			b[i] = tmp
		}
	}

	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
