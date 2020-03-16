// Implement Pic.
// It should return a slice of length dy, each element of
// which is a slice of dx 8-bit unsigned integers.
// When you run the program, it will display your picture,
// interpreting the integers as grayscale (well, bluescale)
// values.
//
// The choice of image is up to you. Interesting functions
// include (x+y)/2, x*y, and x^y.
//
// (You need to use a loop to allocate each []uint8 inside
// the [][]uint8.)
//
// (Use uint8(intValue) to convert between types.)
package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	var pic [][]uint8
	for y := 0; y < dy; y++ {
		var inner_pic []uint8
		for x := 0; x < dx; x++ {
			math_function := (x + y) / 2 // x*y, x^y
			inner_pic = append(inner_pic, uint8(math_function))
		}
		pic = append(pic, inner_pic)
	}
	return pic
}

func main() {
	pic.Show(Pic)
}
