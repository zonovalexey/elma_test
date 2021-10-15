package tasks

import (
	"fmt"
	"sort"
	"tasks/generators"
	"time"
)

var user_settings4 generators.Task_4

func solution4(t generators.Task_4) int {
	A := t.Mass
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
	user_settings4.SetSettings()
	user_hello4(&user_settings4)
	start := time.Now()
	user_settings4.Mass = user_settings4.Generator()
	duration := time.Since(start)
	if user_settings4.Mass_on == 1 {
		fmt.Println("Исходный массив:", user_settings4.Mass)
	}
	fmt.Println("Время создания массива размером", len(user_settings4.Mass), "элементов:", duration)
	start = time.Now()
	res := solution4(user_settings4)
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

func user_hello4(t *generators.Task_4) *generators.Task_4 {
	var mass_on string

	for t.Mass_size < 1 || t.Mass_size > t.Default_mass_size {
		fmt.Printf("\nВведите размер массива (1 - %v)? ", t.Default_mass_size)
		fmt.Scanf("%d\n", &t.Mass_size)
		if t.Mass_size < 1 || t.Mass_size > t.Default_mass_size {
			fmt.Println("\nРазмер массива должен быть в пределах от 1 до", t.Default_mass_size)
		}
	}

	for t.El_size < t.Mass_size || t.El_size > t.Default_el_size {
		fmt.Printf("\nВведите максимальный размер элемента массива (%v - %v)? ", t.Mass_size, t.Default_el_size)
		fmt.Scanf("%d\n", &t.El_size)
		if t.El_size < t.Mass_size || t.El_size > t.Default_el_size {
			fmt.Println("\nРазмер элемента массива должен быть в пределах от", t.Mass_size, "до", t.Default_el_size)
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
