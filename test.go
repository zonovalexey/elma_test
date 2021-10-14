package main

import (
	"fmt"
	"tasks/test_1"
	"tasks/test_2"
)

func select_work() int {
	work := -1
	for !(work >= 0 && work <= 2) {
		fmt.Printf("\nВыберите задачу (1 - Циклическая ротация, 2 - Чудные вхождения в массив, 0 - выход): ")
		fmt.Scanf("%d\n", &work)
		if !(work >= 0 && work <= 2) {
			fmt.Println("\nВведите 0, 1 или 2")
		}
	}
	return work
}

func main() {
	selected_work := -1
	for selected_work != 0 {
		selected_work = select_work()
		switch selected_work {
		case 0:
			{
				return
			}
		case 1:
			{
				test_1.TaskSolution()
			}
		case 2:
			{
				test_2.TaskSolution()
			}
		}
	}
}
