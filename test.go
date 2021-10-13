package main

import (
	"fmt"
	"time"

	"tasks/test_2"
)

func select_work() int {
	work := 0
	for work != 1 && work != 2 {
		fmt.Printf("\nВыберите задачу (1 - Циклическая ротация, 2 - Чудные вхождения в массив): ")
		fmt.Scanf("%d\n", &work)
		if work != 1 && work != 2 {
			fmt.Println("\nВведите 1 или 2")
		}
	}
	return work
}

func main() {
	selected_work := select_work()
	switch selected_work {
	case 1:
		{
			fmt.Println("Не сделана ")
		}
	case 2:
		{
			user_settings := test_2.User_hello()
			if user_settings.User_params == 1 {
				test_2.A = user_settings.Mass_size
				test_2.N = user_settings.Number_size
			}
			start := time.Now()
			mass := test_2.Unsort_mass(test_2.Gen_mass())
			duration := time.Since(start)
			if user_settings.Mass_on == 1 {
				fmt.Printf("Исходный массив:\n%v\n", mass)
			}
			fmt.Println("\nВремя создания массива размером", len(mass), "элементов:", duration)
			start = time.Now()
			lonely := test_2.Find_lonely(mass)
			duration = time.Since(start)
			fmt.Printf("Элемент %v с индексом %v не имеет пары\n", mass[lonely], lonely)
			fmt.Println("Время поиска элемента:", duration)
		}
	}
}
