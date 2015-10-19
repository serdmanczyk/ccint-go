
package q07

import (
	"strings"
	"fmt"
	"math/rand"
	"time"
)

type Matrix [][]int

var gen *rand.Rand

func init() {
	seed := time.Now().UnixNano()
	source := rand.NewSource(seed)
	gen = rand.New(source)
}

func RandMatrix(m, n int) Matrix {
	matrix := make(Matrix, m)
	for i := 0; i < m; i++ {
		matrix[i] = make([]int, n)
		for j := 0; j < n; j++ {
			matrix[i][j] = gen.Int() % 10
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
			sj = append(sj, fmt.Sprintf("%d", ma[i][j]))
		}
		si = append(si, strings.Join(sj, ", "))
	}

	return strings.Join(si, "\n")
}
