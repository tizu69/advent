package c

func Abs(in int) int {
	if in < 0 {
		return -in
	}
	return in
}

func Delta(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
