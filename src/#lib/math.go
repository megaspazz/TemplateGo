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
	if x < 0 {
		return -x
	} else {
		return x
	}
}

func Min64(args ...int64) int64 {
	min := args[0]
	for i := 1; i < len(args); i++ {
		if args[i] < min {
			min = args[i]
		}
	}
	return min
}

func Max64(args ...int64) int64 {
	max := args[0]
	for i := 1; i < len(args); i++ {
		if args[i] > max {
			max = args[i]
		}
	}
	return max
}

func Abs64(x int64) int64 {
	if x < 0 {
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

func GCDAll(args ...int64) int64 {
	ans := args[0]
	for i := 1; i < len(args); i++ {
		ans = GCD(ans, args[i])
	}
	return ans
}

func ComputeCombinations(n int, mod int64) [][]int64 {
	combos := make([][]int64, n+1)
	combos[0] = []int64{1, 0}
	for r := 1; r <= n; r++ {
		combos[r] = make([]int64, r+2)
		combos[r][0] = 1
		for c := 1; c <= r; c++ {
			combos[r][c] = (combos[r-1][c-1] + combos[r-1][c]) % mod
		}
	}
	return combos
}
