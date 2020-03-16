// A common pattern is an io.Reader that wraps another io.Reader,
// modifying the stream in some way.
//
// For example, the gzip.NewReader function takes an io.Reader
// (a stream of compressed data) and returns a *gzip.Reader that
// also implements io.Reader (a stream of the decompressed data).
//
// Implement a rot13Reader that implements io.Reader and reads
// from an io.Reader, modifying the stream by applying the rot13
// substitution cipher to all alphabetical characters.
//
// The rot13Reader type is provided for you. Make it an io.Reader
// by implementing its Read method.
package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot rot13Reader) Read(p []byte) (n int, err error) {
	b, err := rot.r.Read(p)
	for i, char := range p {
		if char > 64 && char < 91 { // A-Z
			p[i] = 65 + ((char - 65 + 13) % 26)
		} else if char > 96 && char < 123 { // a-z
			p[i] = 97 + ((char - 97 + 13) % 26)
		}
	}
	return b, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	s_after_rot13 := make([]byte, s.Len())
	r.Read(s_after_rot13)
	io.Copy(os.Stdout, &r)
	fmt.Println(string(s_after_rot13))
}
