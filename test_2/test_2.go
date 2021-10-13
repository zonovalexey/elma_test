package test_2

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type Settings struct {
	User_params, Mass_on, Mass_size, Number_size int
}

var A = 1000000
var N = 1000000000

func Gen_mass() []int {
	mass := make([]int, 0, A)
	generator := rand.New(rand.NewSource(time.Now().UnixNano()))
	mass_size := generator.Intn(A / 2)
	for i := 0; i <= mass_size; i++ {
		num := generator.Intn(N) + 1
		mass = append(mass, num)
		mass = append(mass, num)
	}
	mass = append(mass, generator.Intn(N)+1)
	return mass
}

func Unsort_mass(mass []int) []int {
	unsorted_mass := make([]int, 0, len(mass))
	generator := rand.New(rand.NewSource(time.Now().UnixNano()))
	length := len(mass)
	for i := 0; i < length; i++ {
		ind := generator.Intn(int(len(mass)))
		unsorted_mass = append(unsorted_mass, mass[ind])
		mass[ind] = mass[len(mass)-1]
		mass = mass[:len(mass)-1]
	}
	return unsorted_mass
}

func Find_lonely(mass []int) int {
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

func User_hello() Settings {
	var user_settings Settings
	var user_params, mass_on string
	user_settings.User_params = 0
	user_settings.Mass_on = 0
	user_settings.Mass_size = 0
	user_settings.Number_size = 0
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
		for user_settings.Mass_size < 3 || user_settings.Mass_size > 1000000 {
			fmt.Printf("\nВведите размер массива (3 - 1000000)? ")
			fmt.Scanf("%d\n", &user_settings.Mass_size)
			if user_settings.Mass_size < 3 || user_settings.Mass_size > A {
				fmt.Println("\nРазмер массива должен быть в пределах от 3 до 1000000")
			}
		}
		for user_settings.Number_size < 1 || user_settings.Number_size > 1000000000 {
			fmt.Printf("\nВведите максимальный размер элемента массива (1 - 1000000000)? ")
			fmt.Scanf("%d\n", &user_settings.Number_size)
			if user_settings.Number_size < 3 || user_settings.Number_size > A {
				fmt.Println("\nРазмер элемента массива должен быть в пределах от 1 до 1000000000")
			}
		}
	}
	return user_settings
}
