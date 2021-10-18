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

// func solutionTask4a(A []int) int {
// 	temp_mass := make([]int, len(A))
// 	copy(temp_mass, A)
// 	sort.Ints(temp_mass)
// 	for i := 0; i < len(temp_mass)-1; i++ {
// 		if temp_mass[i]+1 != temp_mass[i+1] {
// 			return 0
// 		}
// 	}
// 	return 1
// }
