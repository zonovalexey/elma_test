package tasks

import (
	"fmt"
	"sort"
	"tasks/generators"
	"time"
)

var user_settings2 generators.Task_2

func solution2(t generators.Task_2) int {
	lonely := len(t.Mass) - 1
	temp_mass := make([]int, len(t.Mass))
	copy(temp_mass, t.Mass)
	sort.Ints(temp_mass)
	for i := 0; i < len(temp_mass)-2; i += 2 {
		if temp_mass[i] != temp_mass[i+1] {
			lonely = i
			break
		}
	}
	for i, v := range t.Mass {
		if v == temp_mass[lonely] {
			return i
		}
	}
	return -1
}

func solution2a(t generators.Task_2) int {
	lonely := 0

	for i, v := range t.Mass {
		lonely = 0
		for i2, v2 := range t.Mass {
			if v == v2 && i != i2 {
				lonely = 1
				break
			}
		}
		if lonely == 0 {
			return i
		}
	}
	return -1
}

func solution2b(t generators.Task_2) int {
	var m = make(map[int]int, len(t.Mass))

	for _, v := range t.Mass {
		m[v]++
	}

	for i, v := range m {
		if v%2 == 1 {
			return i
		}
	}
	return -1
}

func TaskSolution2() {
	user_settings2.SetSettings()
	user_hello2(&user_settings2)
	if user_settings2.User_params == 0 {
		user_settings2.Mass_size = user_settings2.Default_mass_size
		user_settings2.El_size = user_settings2.Default_el_size
	}
	start := time.Now()

	user_settings2.Mass = user_settings2.Generator()
	duration := time.Since(start)
	if user_settings2.Mass_on == 1 {
		fmt.Println("Исходный массив:\n", user_settings2.Mass)
	}
	fmt.Println("Время создания массива размером", len(user_settings2.Mass), "элементов:", duration)
	start = time.Now()
	lonely := solution2(user_settings2)
	duration = time.Since(start)
	fmt.Printf("Элемент %v с индексом %v не имеет пары\n", user_settings2.Mass[lonely], lonely)
	fmt.Println("Время поиска элемента:", duration)
	start = time.Now()
	lonely = solution2a(user_settings2)
	duration = time.Since(start)
	fmt.Printf("Элемент %v с индексом %v не имеет пары\n", user_settings2.Mass[lonely], lonely)
	fmt.Println("Время поиска элемента:", duration)
	start = time.Now()
	lonely2 := solution2b(user_settings2)
	duration = time.Since(start)
	fmt.Printf("Элемент %v не имеет пары\n", lonely2)
	fmt.Println("Время поиска элемента:", duration)
}

func user_hello2(t *generators.Task_2) *generators.Task_2 {
	var user_params, mass_on string

	user_params = "y"
	mass_on = "y"
	t.El_size = 1000000000
	t.Mass_on = 0
	t.Mass_size = 100000
	t.User_params = 1

	for user_params != "y" && user_params != "Y" && user_params != "n" && user_params != "N" {
		fmt.Printf("\nЗадать параметры вручную (y/n)? ")
		fmt.Scanf("%s", &user_params)
		if user_params != "y" && user_params != "Y" && user_params != "n" && user_params != "N" {
			fmt.Println("\nВаш ответ должен быть y или n")
		} else {
			if user_params == "y" || user_params == "Y" {
				t.User_params = 1
			}
			if user_params == "n" || user_params == "N" {
				t.User_params = 0
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
				t.Mass_on = 1
			}
			if mass_on == "n" || mass_on == "N" {
				t.Mass_on = 0
			}
		}
	}

	if t.User_params == 1 {
		for t.Mass_size < 3 || t.Mass_size > t.Default_mass_size {
			fmt.Printf("\nВведите размер массива (3 - %v)? ", t.Default_mass_size)
			fmt.Scanf("%d\n", &t.Mass_size)
			if t.Mass_size < 3 || t.Mass_size > t.Default_mass_size {
				fmt.Println("\nРазмер массива должен быть в пределах от 3 до", t.Default_mass_size)
			}
		}

		for t.El_size < 1 || t.El_size > t.Default_el_size {
			fmt.Printf("\nВведите максимальный размер элемента массива (1 - %v)? ", t.Default_el_size)
			fmt.Scanf("%d\n", &t.El_size)
			if t.El_size < 3 || t.El_size > t.Default_el_size {
				fmt.Println("\nРазмер элемента массива должен быть в пределах от 1 до", t.Default_el_size)
			}
		}
	}
	return t
}
