package solutions

func solutionTask3(A []int) int {
	var m = make(map[int]int, len(A)+2)

	for _, v := range A {
		m[v] = 1
	}

	for i, v := range m {
		if v == 0 && i != 0 {
			return i
		}
	}

	return -1
}
