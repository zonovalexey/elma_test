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

// func solutionTask3a(A []int) int {
// 	temp_mass := make([]int, len(A))
// 	copy(temp_mass, A)
// 	sort.Ints(temp_mass)
// 	for i := 0; i < len(temp_mass)-1; i++ {
// 		if temp_mass[i]+1 != temp_mass[i+1] {
// 			return temp_mass[i] + 1
// 		}
// 	}
// 	if temp_mass[0] == 0 {
// 		return len(temp_mass)
// 	}
// 	if temp_mass[len(temp_mass)-1] == len(temp_mass) {
// 		return 0
// 	}
// 	return -1
// }
