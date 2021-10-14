package test_1

import (
	"fmt"
	"math/rand"
	"time"
)

const default_el_size = 2000 //-1000 ... 1000
const default_mass_size = 100

type settings struct {
	user_shift, mass_size int
}

func solution(A []int, K int) []int {
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

func gen_mass(user_settings settings) []int {
	mass := make([]int, 0, user_settings.mass_size)
	generator := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < user_settings.mass_size; i++ {
		mass = append(mass, generator.Intn(default_el_size)+1-1000)
	}
	return mass
}

func TaskSolution() {
	user_settings := user_hello()
	start := time.Now()
	mass := gen_mass(user_settings)
	duration := time.Since(start)
	fmt.Println("\nИсходный массив:\n", mass)
	fmt.Println("Время создания массива размером", len(mass), "элементов:", duration)
	start = time.Now()
	mass = solution(mass, user_settings.user_shift)
	duration = time.Since(start)
	fmt.Println("Результат:\n", mass)
	fmt.Println("Время смещения массива:", duration)
}

func user_hello() settings {
	var user_settings settings
	user_settings.user_shift = -1
	user_settings.mass_size = 0

	for user_settings.mass_size < 1 || user_settings.mass_size > default_mass_size {
		fmt.Printf("\nВведите размер массива (1 - %v)? ", default_mass_size)
		fmt.Scanf("%d\n", &user_settings.mass_size)
		if user_settings.mass_size < 1 || user_settings.mass_size > default_mass_size {
			fmt.Println("\nРазмер массива должен быть в пределах от 1 до", default_mass_size)
		}
	}

	for user_settings.user_shift < 0 || user_settings.user_shift > default_mass_size {
		fmt.Printf("\nВведите смещение массива (1 - %v)? ", default_mass_size)
		fmt.Scanf("%d\n", &user_settings.user_shift)
		if user_settings.user_shift < 1 || user_settings.user_shift > default_mass_size {
			fmt.Println("\nЗначение смещения должно быть в пределах от 0 до", default_mass_size)
		}
	}

	return user_settings
}
