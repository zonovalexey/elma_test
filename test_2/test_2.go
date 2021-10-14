package test_2

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type settings struct {
	user_params, mass_on, mass_size, el_size int
}

const default_mass_size = 1000000
const default_el_size = 1000000000

func gen_mass(user_settings settings) []int {
	generator := rand.New(rand.NewSource(time.Now().UnixNano()))
	mass_size := user_settings.mass_size / 2
	if user_settings.user_params == 0 {
		mass_size = generator.Intn(user_settings.mass_size / 2)
	}
	mass := make([]int, 0, user_settings.mass_size)
	for i := 0; i < mass_size; i++ {
		num := generator.Intn(user_settings.el_size) + 1
		mass = append(mass, num)
		mass = append(mass, num)
	}
	mass = append(mass, generator.Intn(user_settings.el_size)+1)
	return mass
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

func solution(mass []int) int {
	lonely := len(mass) - 1
	temp_mass := make([]int, len(mass))
	copy(temp_mass, mass)
	sort.Ints(temp_mass)
	for i := 0; i < len(temp_mass)-2; i += 2 {
		if temp_mass[i] != temp_mass[i+1] {
			lonely = i
			break
		}
	}
	for i, v := range mass {
		if v == temp_mass[lonely] {
			return i
		}
	}
	return -1
}

func TaskSolution() {
	user_settings := user_hello()
	if user_settings.user_params == 0 {
		user_settings.mass_size = default_mass_size
		user_settings.el_size = default_el_size
	}
	start := time.Now()
	mass := unsort_mass(gen_mass(user_settings))
	duration := time.Since(start)
	if user_settings.mass_on == 1 {
		fmt.Println("Исходный массив:\n", mass)
	}
	fmt.Println("Время создания массива размером", len(mass), "элементов:", duration)
	start = time.Now()
	lonely := solution(mass)
	duration = time.Since(start)
	fmt.Printf("Элемент %v с индексом %v не имеет пары\n", mass[lonely], lonely)
	fmt.Println("Время поиска элемента:", duration)
}

func user_hello() settings {
	var user_settings settings
	var user_params, mass_on string
	user_settings.user_params = 0
	user_settings.mass_on = 0
	user_settings.mass_size = 0
	user_settings.el_size = 0

	for user_params != "y" && user_params != "Y" && user_params != "n" && user_params != "N" {
		fmt.Printf("\nЗадать параметры вручную (y/n)? ")
		fmt.Scanf("%s", &user_params)
		if user_params != "y" && user_params != "Y" && user_params != "n" && user_params != "N" {
			fmt.Println("\nВаш ответ должен быть y или n")
		} else {
			if user_params == "y" || user_params == "Y" {
				user_settings.user_params = 1
			}
			if user_params == "n" || user_params == "N" {
				user_settings.user_params = 0
			}
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

	if user_settings.user_params == 1 {
		for user_settings.mass_size < 3 || user_settings.mass_size > default_mass_size {
			fmt.Printf("\nВведите размер массива (3 - %v)? ", default_mass_size)
			fmt.Scanf("%d\n", &user_settings.mass_size)
			if user_settings.mass_size < 3 || user_settings.mass_size > default_mass_size {
				fmt.Println("\nРазмер массива должен быть в пределах от 3 до", default_mass_size)
			}
		}

		for user_settings.el_size < 1 || user_settings.el_size > default_el_size {
			fmt.Printf("\nВведите максимальный размер элемента массива (1 - %v)? ", default_el_size)
			fmt.Scanf("%d\n", &user_settings.el_size)
			if user_settings.el_size < 3 || user_settings.el_size > default_el_size {
				fmt.Println("\nРазмер элемента массива должен быть в пределах от 1 до", default_el_size)
			}
		}
	}
	return user_settings
}
