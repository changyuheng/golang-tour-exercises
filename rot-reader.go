// +build ignore

package main

// Exercise: rot13Reader
// A common pattern is an io.Reader that wraps another io.Reader,
// modifying the stream in some way.

// For example, the gzip.NewReader function takes an io.Reader
// (a stream of compressed data) and returns a *gzip.Reader that also implements
// io.Reader (a stream of the decompressed data).

// Implement a rot13Reader that implements io.Reader and reads from an
// io.Reader, modifying the stream by applying the rot13 substitution cipher
// to all alphabetical characters.

// The rot13Reader type is provided for you. Make it an io.Reader by
// implementing its Read method.

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func rot13(b byte) byte {
	switch {
	case 'A' <= b && b <= 'M':
		fallthrough
	case 'a' <= b && b <= 'm':
		b += 13
	case 'N' <= b && b <= 'Z':
		fallthrough
	case 'n' <= b && b <= 'z':
		b -= 13
	}
	return b
}

func (r rot13Reader) Read(bs []byte) (int, error) {
	n, err := r.r.Read(bs)
	for i := range bs {
		bs[i] = rot13(bs[i])
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
