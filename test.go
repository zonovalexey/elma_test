package main

import (
	"fmt"
	"tasks/test_1"
	"tasks/test_2"
	"tasks/test_3"
	"tasks/test_4"
)

func select_work() int {
	work := -1
	for !(work >= 0 && work <= 4) {
		fmt.Println("\nВыберите задачу:")
		fmt.Println("1 - Циклическая ротация")
		fmt.Println("2 - Чудные вхождения в массив")
		fmt.Println("3 - Поиск отсутствующего элемента")
		fmt.Println("4 - Проверка последовательности")
		fmt.Println("0 - выход")
		fmt.Scanf("%d\n", &work)
		if !(work >= 0 && work <= 4) {
			fmt.Println("\nВведите 0, 1, 2, 3 или 4")
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
		case 3:
			{
				test_3.TaskSolution()
			}
		case 4:
			{
				test_4.TaskSolution()
			}
		}
	}
}
