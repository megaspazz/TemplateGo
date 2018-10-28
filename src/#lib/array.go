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
 * Generic function to fill an array.
 *
 * RESERVED GENERIC NAMES:
 *   - TYPE: type of the array to fill.
 */
func FillArray(arr []TYPE, val TYPE) {
	for i := range arr {
		arr[i] = val
	}
}
