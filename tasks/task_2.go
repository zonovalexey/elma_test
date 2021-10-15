package tasks

import (
	"fmt"
	"sort"
	"tasks/generators"
	"time"
)

func solution2(mass []int) int {
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

func TaskSolution2() {
	user_settings := user_hello()
	if user_settings.User_params == 0 {
		user_settings.Mass_size = default_mass_size
		user_settings.El_size = default_el_size
	}
	start := time.Now()
	/*
		var user_settings generators.Task_2
		user_settings.Mass_size = 20
		user_settings.El_size = 100
		user_settings.Mass_on = 1
		user_settings.User_params = 1
	*/
	mass := user_settings.Generator()
	duration := time.Since(start)
	if user_settings.Mass_on == 1 {
		fmt.Println("Исходный массив:\n", mass)
	}
	fmt.Println("Время создания массива размером", len(mass), "элементов:", duration)
	start = time.Now()
	lonely := solution(mass)
	duration = time.Since(start)
	fmt.Printf("Элемент %v с индексом %v не имеет пары\n", mass[lonely], lonely)
	fmt.Println("Время поиска элемента:", duration)
}

func user_hello2() generators.Task_2 {
	var user_settings generators.Task_2
	user_settings.New()
	var user_params, mass_on string

	for user_params != "y" && user_params != "Y" && user_params != "n" && user_params != "N" {
		fmt.Printf("\nЗадать параметры вручную (y/n)? ")
		fmt.Scanf("%s", &user_params)
		if user_params != "y" && user_params != "Y" && user_params != "n" && user_params != "N" {
			fmt.Println("\nВаш ответ должен быть y или n")
		} else {
			if user_params == "y" || user_params == "Y" {
				user_settings.User_params = 1
			}
			if user_params == "n" || user_params == "N" {
				user_settings.User_params = 0
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
				user_settings.Mass_on = 1
			}
			if mass_on == "n" || mass_on == "N" {
				user_settings.Mass_on = 0
			}
		}
	}

	if user_settings.User_params == 1 {
		for user_settings.Mass_size < 3 || user_settings.Mass_size > default_mass_size {
			fmt.Printf("\nВведите размер массива (3 - %v)? ", default_mass_size)
			fmt.Scanf("%d\n", &user_settings.Mass_size)
			if user_settings.Mass_size < 3 || user_settings.Mass_size > default_mass_size {
				fmt.Println("\nРазмер массива должен быть в пределах от 3 до", default_mass_size)
			}
		}

		for user_settings.El_size < 1 || user_settings.El_size > default_el_size {
			fmt.Printf("\nВведите максимальный размер элемента массива (1 - %v)? ", default_el_size)
			fmt.Scanf("%d\n", &user_settings.El_size)
			if user_settings.El_size < 3 || user_settings.El_size > default_el_size {
				fmt.Println("\nРазмер элемента массива должен быть в пределах от 1 до", default_el_size)
			}
		}
	}
	return user_settings
}
