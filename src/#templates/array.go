
func FillArray(arr []int, val int) {
	for i := range arr {
		arr[i] = val
	}
}

func MakeMatrix(rows, cols int) [][]int {
	mat := make([][]int, rows)
	for r := range mat {
		mat[r] = make([]int, cols)
	}
	return mat
}
