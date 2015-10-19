// Given an image represented by an NxN matrix, where each pixel in the image is 4
//  bytes, write a method to rotate the image by 90 degrees. Can you do this in place?

package q06

import (
	"strings"
)

type Matrix [][]Pixel

func RandMatrix(n int) Matrix {
	matrix := make(Matrix, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]Pixel, n)
		for j := 0; j < n; j++ {
			matrix[i][j] = randpixel()
		}
	}

	return matrix
}

func (ma Matrix) String() string {
	il := len(ma)
	si := make([]string, 0, il)
	for i := 0; i < il; i++ {
		jl := len(ma[i])
		sj := make([]string, 0, jl)
		for j := 0; j < jl; j++ {
			sj = append(sj, ma[i][j].String())
		}
		si = append(si, strings.Join(sj, ", "))
	}

	return strings.Join(si, "\n")
}
