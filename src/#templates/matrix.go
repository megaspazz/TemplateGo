
func MakeMatrix(rows, cols int) [][]int {
	mat := make([][]int, rows)
	for i := range mat {
		mat[i] = make([]int, cols)
	}
	return mat
}
