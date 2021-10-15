package tasks

import (
	"fmt"
	"tasks/generators"
	"tasks/settings"
	"time"
)

var user_settings settings.Task_1
user_settings.Now()


func Solution1(A []int, K int) []int {
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
	//var user_settings generators.Task_1
	user_settings.New()
	user_settings.user_hello()
	start := time.Now()
	user_settings.Mass = user_settings.Generator()
	duration := time.Since(start)
	fmt.Println("\nИсходный массив:\n", mass)
	fmt.Println("Время создания массива размером", len(mass), "элементов:", duration)
	start = time.Now()
	mass = solution(mass, user_settings.User_shift)
	duration = time.Since(start)
	fmt.Println("\nРезультат:\n", mass)
	fmt.Println("Время смещения массива:", duration)
}

func (t generators.Task_2) user_hello() generators.Task_1 {

	for t.Mass_size < 1 || t.Mass_size > default_mass_size {
		fmt.Printf("\nВведите размер массива (1 - %v)? ", t.Default_mass_size)
		fmt.Scanf("%d\n", &user_settings.Mass_size)
		if t.Mass_size < 1 || t.Mass_size > t.Default_mass_size {
			fmt.Println("\nРазмер массива должен быть в пределах от 1 до", t.Default_mass_size)
		}
	}

	for t.User_shift < 0 || t.User_shift > default_mass_size {
		fmt.Printf("\nВведите смещение массива (1 - %v)? ", t.Default_mass_size)
		fmt.Scanf("%d\n", &t.User_shift)
		if t.User_shift < 1 || t.User_shift > t.Default_mass_size {
			fmt.Println("\nЗначение смещения должно быть в пределах от 0 до", t.Default_mass_size)
		}
	}

	return user_settings
}
