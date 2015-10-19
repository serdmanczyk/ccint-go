// Given an image represented by an NxN matrix, where each pixel in the image is 4
//  bytes, write a method to rotate the image by 90 degrees. Can you do this in place?

package q06

import (
	"time"
	"math/rand"
)

type Pixel [4]byte

var gen *rand.Rand

func init() {
	seed := time.Now().UnixNano()
	source := rand.NewSource(seed)
	gen = rand.New(source)
}

func randpixel() Pixel {
	var b Pixel

	for i := 0; i < 4; i++ {
		b[i] = byte((gen.Int() % 26) + 97)
	}

	return b
}

func (p Pixel) String() string {
	return string(p[:])
}
