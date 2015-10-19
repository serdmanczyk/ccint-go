// Write an algorithm such that if an element in an MxN matrix is 0, its entire row
//  and column are set to 0.

package q07

func ZeroInPlace(mat *Matrix) {
	ma := *mat
	m := len(ma)
	n := len(ma[0])

	// We're going to use the top row and left
	// column to store which rows/columns
	// need to be zeroed.  However, we still
	// need a place to store data about the
	// top row and column.  These two variables
	// store that data.
	toprow := 1
	leftcolumn := 1

	// Check if any cells in left column are zero
	for i := 0; i < m; i++ {
		if ma[i][0] == 0 {
			leftcolumn = 0
			break
		}
	}

	// Check if any cells in top row are zero
	for j := 0; j < n; j++ {
		if ma[0][j] == 0 {
			toprow = 0
			break
		}
	}

	// Check if any remaining cells are zero.
	// Store data in top row and left column
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if ma[i][j] == 0 {
				ma[i][0] = 0
				ma[0][j] = 0
			}
		}
	}

	// Clear any rows w/zeroes
	for i := 1; i < m; i++ {
		if ma[i][0] == 0 {
			for j := 0; j < n; j++ {
				ma[i][j] = 0
			}
		}
	}

	// Clear any columns w/zeroes
	for j := 1; j < n; j++ {
		if ma[0][j] == 0 {
			for i := 0; i < m; i++ {
				ma[i][j] = 0
			}
		}
	}

	// Clear top row if applicable
	if toprow == 0 {
		for j := 0; j < n; j++ {
			ma[0][j] = 0
		}
	}

	// Clear left column if applicable
	if leftcolumn == 0 {
		for i := 0; i < m; i++ {
			ma[i][0] = 0
		}
	}
}

func ZeroLessWork(mat *Matrix) {
	ma := *mat
	m := len(ma)
	n := len(ma[0])

	// Make arrays for columns and rows.
	// If any value in any column or row is
	// zero that value will be set in its
	// associated array
	rows := make([]int, m)
	columns := make([]int, n)

	// Check every cell, if it is zero
	// set the value in the associated row
	// and column arrays to denote that row/
	// /column should be zeroed.
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if ma[i][j] == 0 {
				// slices are initialized to 0
				// so we use 1's to save time.
				rows[i] = 1
				columns[j] = 1
			}
		}
	}

	// Clear any rows with zeroes
	for i := 0; i < m; i++ {
		if rows[i] == 1 {
			for j := 0; j < n; j++ {
				ma[i][j] = 0
			}
		}
	}

	// Clear any columns with zeroes
	for j := 0; j < n; j++ {
		if columns[j] == 1 {
			for i := 0; i < m; i++ {
				ma[i][j] = 0
			}
		}
	}
}
