package isupov

func SolutionSquareGenerator(start int, n int) []int {
	if n < 1 {
		panic("n must be greater than 0")
	}
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = start * start
		start++
	}
	return res
}