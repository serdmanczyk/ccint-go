// Given an image represented by an NxN matrix, where each pixel in the image is 4
//  bytes, write a method to rotate the image by 90 degrees. Can you do this in place?

package q06

func (m Matrix) Rotate() {
	n := len(m)

	// Work our way inwards from the outer shell of the matrix
	for z := 0; z < n/2; z++ {
		// Move our index from where we'll work
		for i := z; i < n-z-1; i++ {
			// Rotate values at four corners
			// upper left  m[z][i]
			// upper right m[i][n-1-z]
			// lower right m[n-1-z][n-1-i]
			// lower left  m[n-1-i][n-1-z]
			h := m[z][i]
			m[z][i] = m[i][n-1-z]
			m[i][n-1-z] = m[n-1-z][n-1-i]
			m[n-1-z][n-1-i] = m[n-1-i][z]
			m[n-1-i][z] = h
		}
	}
}
