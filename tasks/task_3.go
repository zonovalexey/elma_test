package tasks

import (
	"fmt"
	"sort"
	"tasks/generators"
	"time"
)

var user_settings3 generators.Task_3

func solution3(t generators.Task_3) int {
	A := t.Mass
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
	user_settings3.SetSettings()
	user_hello3(&user_settings3)
	start := time.Now()
	user_settings3.Mass = user_settings3.Generator()
	duration := time.Since(start)
	if user_settings3.Mass_on == 1 {
		fmt.Println("Исходный массив:\n", user_settings3.Mass)
	}
	fmt.Println("Время создания массива размером", len(user_settings3.Mass), "элементов:", duration)
	start = time.Now()
	res := solution3(user_settings3)
	duration = time.Since(start)
	if res != -1 {
		fmt.Println("Отсутствующий элемент массива:", res)
	} else {
		fmt.Println("Элемент не найден")
	}
	fmt.Println("Время поиска элемента:", duration)
}

func user_hello3(t *generators.Task_3) *generators.Task_3 {
	var mass_on string

	for t.Mass_size < 1 || t.Mass_size > t.Default_mass_size {
		fmt.Printf("\nВведите размер массива (1 - %v)? ", t.Default_mass_size)
		fmt.Scanf("%d\n", &t.Mass_size)
		if t.Mass_size < 1 || t.Mass_size > t.Default_mass_size {
			fmt.Println("\nРазмер массива должен быть в пределах от 1 до", t.Default_mass_size)
		}
	}

	for mass_on != "y" && mass_on != "Y" && mass_on != "n" && mass_on != "N" {
		fmt.Printf("\nОтображать исходный массив (y/n)? ")
		fmt.Scanf("%s", &mass_on)
		if mass_on != "y" && mass_on != "Y" && mass_on != "n" && mass_on != "N" {
			fmt.Println("\nВаш ответ должен быть y или n")
		} else {
			if mass_on == "y" || mass_on == "Y" {
				t.Mass_on = 1
			}
			if mass_on == "n" || mass_on == "N" {
				t.Mass_on = 0
			}
		}
	}
	return t
}
