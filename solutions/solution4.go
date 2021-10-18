package solutions

func solutionTask4(A []int) int {
	max := 0
	for _, v := range A {
		if v > max {
			max = v
		}
	}

	if max == len(A) {
		return 1
	}

	return 0
}
