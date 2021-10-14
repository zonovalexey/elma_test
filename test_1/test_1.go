package test_1

import (
	"fmt"
	"math/rand"
	"time"
)

var n = 100
var k = 100
var a = 2000

const nn = 100

type settings struct {
	user_shift, mass_size int
}

func solution(A []int, K int) []int {
	if K > len(A) {
		K = len(A) % K
	}
	if K < len(A) && K > 0 {
		mass1 := A[:(len(A) - K)]
		mass2 := A[(len(A) - K):]
		A = append(mass2, mass1...)
	}
	return A
}

func gen_mass() []int {
	mass := make([]int, 0, n)
	generator := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < n; i++ {
		mass = append(mass, generator.Intn(a)+1-1000)
	}
	return mass
}

func TaskSolution() {
	user_settings := user_hello()
	k = user_settings.user_shift
	n = user_settings.mass_size
	start := time.Now()
	mass := gen_mass()
	duration := time.Since(start)
	fmt.Println("Исходный массив:", mass)
	fmt.Println("\nВремя создания массива размером", len(mass), "элементов:", duration)
	start = time.Now()
	mass = solution(mass, k)
	duration = time.Since(start)
	fmt.Println("\nРезультат:", mass)
	fmt.Println("\nВремя смещения массива:", duration)
}

func user_hello() settings {
	var user_settings settings
	user_settings.user_shift = -1
	user_settings.mass_size = 0

	for user_settings.mass_size < 1 || user_settings.mass_size > nn {
		fmt.Printf("\nВведите размер массива (1 - %v)? ", nn)
		fmt.Scanf("%d\n", &user_settings.mass_size)
		if user_settings.mass_size < 1 || user_settings.mass_size > nn {
			fmt.Println("\nРазмер массива должен быть в пределах от 1 до", nn)
		}
	}

	for user_settings.user_shift < 0 || user_settings.user_shift > nn {
		fmt.Printf("\nВведите смещение массива (1 - %v)? ", nn)
		fmt.Scanf("%d\n", &user_settings.user_shift)
		if user_settings.user_shift < 1 || user_settings.user_shift > nn {
			fmt.Println("\nЗначение смещения должно быть в пределах от 0 до", nn)
		}
	}

	return user_settings
}
