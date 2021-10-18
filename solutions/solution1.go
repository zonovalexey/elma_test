package solutions

func solutionTask1(A []int, K int) []int {
	if K > len(A) {
		K = K % len(A)
	}
	if K < len(A) && K > 0 {
		mass1 := A[:(len(A) - K)]
		mass2 := A[(len(A) - K):]
		A = append(mass2, mass1...)
	}
	return A
}
