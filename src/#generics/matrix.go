/*
 * Matrix utilities.
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
