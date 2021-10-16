package main

import (
	"fmt"
	"tasks/tasks"
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

	// var Tasks []generators.Task

	// for i, v := range Tasks {
	// 	go fu
	// 	v.SetSettings(g)
	// 	v.Solution()
	// }
	tasks.TaskSolution2()
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
				tasks.TaskSolution1()
			}
		case 2:
			{
				tasks.TaskSolution2()
			}
		case 3:
			{
				tasks.TaskSolution3()
			}
		case 4:
			{
				tasks.TaskSolution4()
			}
		}
	}
}
