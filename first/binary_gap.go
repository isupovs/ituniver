package isupov

func SolutionBinaryGap(N int) int {
	max, current := 0, 0
	count := false

	for ; N != 0; N >>= 1 {
		if N&1 == 1 {
			if !count {
				count = true
				continue
			}

			if max < current {
				max = current
			}
			current = 0
			continue
		}

		if count {
			current++
		}
	}

	return max
}
