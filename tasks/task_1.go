package tasks

import (
	"fmt"
	"tasks/generators"
	"time"
)

var user_settings1 generators.Task_1

func solution1(t generators.Task_1) []int {
	A := t.Mass
	K := t.User_shift
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

func TaskSolution1() {
	user_settings1.SetSettings()
	user_hello1(&user_settings1)
	start := time.Now()
	user_settings1.Mass = user_settings1.Generator()
	duration := time.Since(start)
	fmt.Println("\nИсходный массив:\n", user_settings1.Mass)
	fmt.Println("Время создания массива размером", len(user_settings1.Mass), "элементов:", duration)
	start = time.Now()
	user_settings1.Mass = solution1(user_settings1)
	duration = time.Since(start)
	fmt.Println("\nРезультат:\n", user_settings1.Mass)
	fmt.Println("Время смещения массива:", duration)
}

func user_hello1(t *generators.Task_1) *generators.Task_1 {

	for t.Mass_size < 1 || t.Mass_size > t.Default_mass_size {
		fmt.Printf("\nВведите размер массива (1 - %v)? ", t.Default_mass_size)
		fmt.Scanf("%d\n", &t.Mass_size)
		if t.Mass_size < 1 || t.Mass_size > t.Default_mass_size {
			fmt.Println("\nРазмер массива должен быть в пределах от 1 до", t.Default_mass_size)
		}
	}

	for t.User_shift < 0 || t.User_shift > t.Default_mass_size {
		fmt.Printf("\nВведите смещение массива (1 - %v)? ", t.Default_mass_size)
		fmt.Scanf("%d\n", &t.User_shift)
		if t.User_shift < 1 || t.User_shift > t.Default_mass_size {
			fmt.Println("\nЗначение смещения должно быть в пределах от 0 до", t.Default_mass_size)
		}
	}

	return t
}
