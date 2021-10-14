package test_3

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

const default_mass_size = 100000

type settings struct {
	mass_on, mass_size int
}

func solution(A []int) int {
	temp_mass := make([]int, len(A))
	copy(temp_mass, A)
	sort.Ints(temp_mass)
	for i := 0; i < len(temp_mass)-1; i++ {
		if temp_mass[i]+1 != temp_mass[i+1] {
			return temp_mass[i] + 1
		}
	}
	if temp_mass[0] == 0 {
		return len(temp_mass)
	}
	if temp_mass[len(temp_mass)-1] == len(temp_mass) {
		return 0
	}
	return -1

}

func gen_mass(user_settings settings) []int {
	mass := make([]int, 0, user_settings.mass_size+1)
	for i := 0; i < user_settings.mass_size+1; i++ {
		mass = append(mass, i)
	}
	generator := rand.New(rand.NewSource(time.Now().UnixNano()))
	s := generator.Intn(int(len(mass)))
	return append(mass[:s], mass[s+1:]...)
}

func unsort_mass(mass []int) []int {
	unsorted_mass := make([]int, 0, len(mass))
	generator := rand.New(rand.NewSource(time.Now().UnixNano()))
	length := len(mass)
	for i := 0; i < length; i++ {
		ind := generator.Intn(int(len(mass)))
		unsorted_mass = append(unsorted_mass, mass[ind])
		mass[ind] = mass[len(mass)-1]
		mass = mass[:len(mass)-1]
	}
	return unsorted_mass
}

func TaskSolution() {
	user_settings := user_hello()
	start := time.Now()
	mass := unsort_mass(gen_mass(user_settings))
	duration := time.Since(start)
	if user_settings.mass_on == 1 {
		fmt.Println("Исходный массив:\n", mass)
	}
	fmt.Println("Время создания массива размером", len(mass), "элементов:", duration)
	start = time.Now()
	res := solution(mass)
	duration = time.Since(start)
	if res != -1 {
		fmt.Println("Отсутствующий элемент массива:", res)
	} else {
		fmt.Println("Элемент не найден")
	}
	fmt.Println("Время поиска элемента:", duration)
}

func user_hello() settings {
	var user_settings settings
	var mass_on string
	user_settings.mass_size = 0
	user_settings.mass_on = 0

	for user_settings.mass_size < 1 || user_settings.mass_size > default_mass_size {
		fmt.Printf("\nВведите размер массива (1 - %v)? ", default_mass_size)
		fmt.Scanf("%d\n", &user_settings.mass_size)
		if user_settings.mass_size < 1 || user_settings.mass_size > default_mass_size {
			fmt.Println("\nРазмер массива должен быть в пределах от 1 до", default_mass_size)
		}
	}

	for mass_on != "y" && mass_on != "Y" && mass_on != "n" && mass_on != "N" {
		fmt.Printf("\nОтображать исходный массив (y/n)? ")
		fmt.Scanf("%s", &mass_on)
		if mass_on != "y" && mass_on != "Y" && mass_on != "n" && mass_on != "N" {
			fmt.Println("\nВаш ответ должен быть y или n")
		} else {
			if mass_on == "y" || mass_on == "Y" {
				user_settings.mass_on = 1
			}
			if mass_on == "n" || mass_on == "N" {
				user_settings.mass_on = 0
			}
		}
	}
	return user_settings
}
