package solutions

func solutionTask2(A []int) int {
	var m = make(map[int]int, len(A))

	for _, v := range A {
		m[v]++
	}

	for i, v := range m {
		if v%2 == 1 {
			return i
		}
	}

	return 0
}
