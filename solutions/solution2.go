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

// func solutionTask2(A []int) int {
// 	lonely := 0

// 	for i, v := range A {
// 		lonely = 0
// 		for i2, v2 := range A {
// 			if v == v2 && i != i2 {
// 				lonely = 1
// 				break
// 			}
// 		}
// 		if lonely == 0 {
// 			return A[i]
// 		}
// 	}
// 	return -1
// }
