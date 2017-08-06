
func FillIntArray(arr []int, val int) {
	for i := range arr {
		arr[i] = val
	}
}

func FillLongArray(arr []int64, val int64) {
	for i := range arr {
		arr[i] = val
	}
}

func FillBoolArray(arr []bool, val bool) {
	for i := range arr {
		arr[i] = val
	}
}

func FillStringArray(arr []string, val string) {
	for i := range arr {
		arr[i] = val
	}
}

/*
 * Replace T with the actual type (2 occurrences).
 */
func FillArray(arr []T, val T) {
	for i := range arr {
		arr[i] = val
	}
}
