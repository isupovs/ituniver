package isupov

func SolutionUniq(A []int) int {
	sliceLen := len(A)

	if sliceLen == 0 {
		return 0
	}

	if sliceLen == 1 {
		return 1
	}

	uniqMap := make(map[int]struct{}, sliceLen)
	for _, v := range A {
		uniqMap[v] = struct{}{}
	}
	return len(uniqMap)
}
