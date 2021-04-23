package isupov

func SolutionRotate(A []int, k int) []int {
	sliceLen := len(A)
	if sliceLen < 2 || k%sliceLen == 0 {
		return A
	}
	shift := sliceLen - k%sliceLen
	return append(A[shift:], A[:shift]...)
}
