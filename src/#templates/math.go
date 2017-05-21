
func Min(lhs, rhs int) int {
	if lhs < rhs {
		return lhs
	} else {
		return rhs
	}
}

func Max(lhs, rhs int) int {
	if lhs > rhs {
		return lhs
	} else {
		return rhs
	}
}

func Abs(x int) int {
	if (x < 0) {
		return -x
	} else {
		return x
	}
}
