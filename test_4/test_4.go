package test_4

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

const default_mass_size = 100000
const default_el_size = 1000000000

type settings struct {
	mass_on, mass_size, el_size int
}

func solution(A []int) int {
	temp_mass := make([]int, len(A))
	copy(temp_mass, A)
	sort.Ints(temp_mass)
	for i := 0; i < len(temp_mass)-1; i++ {
		if temp_mass[i]+1 != temp_mass[i+1] {
			return 0
		}
	}
	return 1
}

func gen_mass(user_settings settings) []int {
	mass := make([]int, 0, user_settings.mass_size+1)
	generator := rand.New(rand.NewSource(time.Now().UnixNano()))
	first_el := generator.Intn(user_settings.el_size) + 1
	if first_el+user_settings.mass_size > user_settings.el_size {
		first_el = user_settings.el_size - user_settings.mass_size + 1
	}
	for i := 0; i <= user_settings.mass_size; i++ {
		mass = append(mass, first_el+i)
	}
	del_indicator := generator.Intn(2)
	if del_indicator == 0 {
		return mass[1:]
	} else {
		s := generator.Intn(int(len(mass)))
		return append(mass[:s], mass[s+1:]...)
	}
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
		fmt.Println("Исходный массив:", mass)
	}
	fmt.Println("Время создания массива размером", len(mass), "элементов:", duration)
	start = time.Now()
	res := solution(mass)
	duration = time.Since(start)
	switch res {
	case 0:
		{
			fmt.Println("Исходный массив не содержит последовательности")
		}
	case 1:
		{
			fmt.Println("Исходный массив содержит последовательность")
		}
	}
	fmt.Println("Время поиска элемента:", duration)
}

func user_hello() settings {
	var user_settings settings
	var mass_on string
	user_settings.mass_size = 0
	user_settings.mass_on = 0
	user_settings.el_size = 0

	for user_settings.mass_size < 1 || user_settings.mass_size > default_mass_size {
		fmt.Printf("\nВведите размер массива (1 - %v)? ", default_mass_size)
		fmt.Scanf("%d\n", &user_settings.mass_size)
		if user_settings.mass_size < 1 || user_settings.mass_size > default_mass_size {
			fmt.Println("\nРазмер массива должен быть в пределах от 1 до", default_mass_size)
		}
	}

	for user_settings.el_size < user_settings.mass_size || user_settings.el_size > default_el_size {
		fmt.Printf("\nВведите максимальный размер элемента массива (%v - %v)? ", user_settings.mass_size, default_el_size)
		fmt.Scanf("%d\n", &user_settings.el_size)
		if user_settings.el_size < user_settings.mass_size || user_settings.el_size > default_el_size {
			fmt.Println("\nРазмер элемента массива должен быть в пределах от", user_settings.mass_size, "до", default_el_size)
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
