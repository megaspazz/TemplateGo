
func Min(args ...int) int {
	min := args[0]
	for i := 1; i < len(args); i++ {
		if args[i] < min {
			min = args[i]
		}
	}
	return min
}

func Max(args ...int) int {
	max := args[0]
	for i := 1; i < len(args); i++ {
		if args[i] > max {
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

func GCD(a, b int64) int64 {
	if a < b {
		return GCD(b, a)
	}

	r := a % b
	if r == 0 {
		return b
	}

	return GCD(b, r)
}

/*
 * Takes the GCD of all input arguments.
 * To use on arrays, change the argumen to be an array instead of variadic parameters.
 */
func GCDAll(args ...int64) int64 {
	ans := args[0]
	for i := 1; i < len(args); i++ {
		ans = GCD(ans, args[i])
	}
	return ans
}
