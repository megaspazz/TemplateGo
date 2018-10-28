func MakeIntMatrix(rows, cols int) [][]int {
	mat := make([][]int, rows)
	for i := range mat {
		mat[i] = make([]int, cols)
	}
	return mat
}

func MakeLongMatrix(rows, cols int) [][]int64 {
	mat := make([][]int64, rows)
	for i := range mat {
		mat[i] = make([]int64, cols)
	}
	return mat
}

func MakeBoolMatrix(rows, cols int) [][]bool {
	mat := make([][]bool, rows)
	for i := range mat {
		mat[i] = make([]bool, cols)
	}
	return mat
}

func MakeStringMatrix(rows, cols int) [][]string {
	mat := make([][]string, rows)
	for i := range mat {
		mat[i] = make([]string, cols)
	}
	return mat
}

/*
 * Generic function to create a matrix.
 *
 * RESERVED GENERIC NAMES:
 *   - TYPE: type of matrix to create.
 */
func MakeMatrix(rows, cols int) [][]TYPE {
	mat := make([][]TYPE, rows)
	for i := range mat {
		mat[i] = make([]TYPE, cols)
	}
	return mat
}
