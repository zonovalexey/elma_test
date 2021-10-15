package tasks

import (
	"fmt"
	"sort"
	"tasks/generators"
	"time"
)

func solution3(A []int) int {
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

func TaskSolution3() {
	user_settings := user_hello3()
	start := time.SetSettings()
	mass := user_settings.Generator()
	duration := time.Since(start)
	if user_settings.Mass_on == 1 {
		fmt.Println("Исходный массив:\n", mass)
	}
	fmt.Println("Время создания массива размером", len(mass), "элементов:", duration)
	start = time.Now()
	res := solution3(mass)
	duration = time.Since(start)
	if res != -1 {
		fmt.Println("Отсутствующий элемент массива:", res)
	} else {
		fmt.Println("Элемент не найден")
	}
	fmt.Println("Время поиска элемента:", duration)
}

func user_hello3() generators.Task_3 {
	var user_settings generators.Task_3
	var mass_on string
	user_settings.Mass_size = 0
	user_settings.Mass_on = 0

	for user_settings.Mass_size < 1 || user_settings.Mass_size > user_settings.Default_mass_size {
		fmt.Printf("\nВведите размер массива (1 - %v)? ", user_settings.Default_mass_size)
		fmt.Scanf("%d\n", &user_settings.Mass_size)
		if user_settings.Mass_size < 1 || user_settings.Mass_size > user_settings.Default_mass_size {
			fmt.Println("\nРазмер массива должен быть в пределах от 1 до", user_settings.Default_mass_size)
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
