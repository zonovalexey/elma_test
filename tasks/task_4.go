package tasks

import (
	"fmt"
	"sort"
	"tasks/generators"
	"time"
)

func solution4(A []int) int {
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

func TaskSolution4() {
	user_settings := user_hello()
	start := time.Now()
	mass := user_settings.Generator()
	duration := time.Since(start)
	if user_settings.Mass_on == 1 {
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

func user_hello4() generators.Task_4 {
	var user_settings generators.Task_4
	var mass_on string
	user_settings.Mass_size = 0
	user_settings.Mass_on = 0
	user_settings.El_size = 0

	for user_settings.Mass_size < 1 || user_settings.Mass_size > default_mass_size {
		fmt.Printf("\nВведите размер массива (1 - %v)? ", default_mass_size)
		fmt.Scanf("%d\n", &user_settings.Mass_size)
		if user_settings.Mass_size < 1 || user_settings.Mass_size > default_mass_size {
			fmt.Println("\nРазмер массива должен быть в пределах от 1 до", default_mass_size)
		}
	}

	for user_settings.El_size < user_settings.Mass_size || user_settings.El_size > default_el_size {
		fmt.Printf("\nВведите максимальный размер элемента массива (%v - %v)? ", user_settings.Mass_size, default_el_size)
		fmt.Scanf("%d\n", &user_settings.El_size)
		if user_settings.El_size < user_settings.Mass_size || user_settings.El_size > default_el_size {
			fmt.Println("\nРазмер элемента массива должен быть в пределах от", user_settings.Mass_size, "до", default_el_size)
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
	return user_settings
}
