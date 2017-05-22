
func Min(args ...int64) int64 {
	min := args[0]
	for i := 1; i < len(args); i++ {
		if (args[i] < min) {
			min = args[i]
		}
	}
	return min
}

func Max(args ...int64) int64 {
	max := args[0]
	for i := 1; i < len(args); i++ {
		if (args[i] > max) {
			max = args[i]
		}
	}
	return max
}

func Abs(x int) int {
	if (x < 0) {
		return -x
	} else {
		return x
	}
}
