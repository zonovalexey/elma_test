package generators

import (
	"math/rand"
	"tasks/settings"
	"time"
)

type Task_1 settings.Task_1
type Task_2 settings.Task_2
type Task_3 settings.Task_3
type Task_4 settings.Task_4

func (t Task_1) Generator() []int {
	default_el_size := 2000
	mass := make([]int, t.Mass_size)
	generator := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < t.Mass_size; i++ {
		mass[i] = generator.Intn(default_el_size) + 1 - 1000
	}
	return mass
}

func (t Task_2) Generator() []int {
	generator := rand.New(rand.NewSource(time.Now().UnixNano()))

	mass_size := t.Mass_size / 2
	if t.User_params == 0 {
		mass_size = generator.Intn(t.Mass_size/2-10) + 10
	}

	temp_mass := make([]int, t.Mass_size+1)
	for i := 0; i <= t.Mass_size; i++ {
		temp_mass[i] = i
	}

	mass := make([]int, t.Mass_size+1)
	for i := 0; i < mass_size; i++ {
		a := generator.Intn(len(temp_mass))
		ind1 := temp_mass[a]
		temp_mass = append(temp_mass[:a], temp_mass[a+1:]...)
		a = generator.Intn(len(temp_mass))
		mass[ind1] = generator.Intn(t.El_size) + 1
		mass[temp_mass[a]] = mass[ind1]
		temp_mass = append(temp_mass[:a], temp_mass[a+1:]...)

	}
	mass[temp_mass[generator.Intn(len(temp_mass))]] = generator.Intn(t.El_size) + 1
	return mass
}

func (t Task_3) Generator() []int {
	generator := rand.New(rand.NewSource(time.Now().UnixNano()))

	temp_mass := make([]int, t.Mass_size+1)
	for i := 0; i <= t.Mass_size; i++ {
		temp_mass[i] = i
	}

	mass := make([]int, t.Mass_size+1)
	for i := 0; i < t.Mass_size+1; i++ {
		a := generator.Intn(len(temp_mass))
		mass[temp_mass[a]] = i
		temp_mass = append(temp_mass[:a], temp_mass[a+1:]...)
	}

	s := generator.Intn(int(len(mass)))
	return append(mass[:s], mass[s+1:]...)
}

func (t Task_4) Generator() []int {
	first_el_index := -1

	mass := make([]int, t.Mass_size+1)

	generator := rand.New(rand.NewSource(time.Now().UnixNano()))

	first_el := generator.Intn(t.El_size) + 1
	if first_el+t.Mass_size > t.El_size {
		first_el = t.El_size - t.Mass_size + 1
	}

	temp_mass := make([]int, t.Mass_size+1)
	for i := 0; i <= t.Mass_size; i++ {
		temp_mass[i] = i
	}

	for i := 0; i <= t.Mass_size; i++ {
		a := generator.Intn(len(temp_mass))
		if first_el_index == -1 {
			first_el_index = temp_mass[a]
		}
		mass[temp_mass[a]] = first_el + i
		temp_mass = append(temp_mass[:a], temp_mass[a+1:]...)
	}

	del_indicator := generator.Intn(2)
	if del_indicator == 0 {
		return append(mass[:first_el_index], mass[first_el_index+1:]...)
	} else {
		s := generator.Intn(int(len(mass)))
		return append(mass[:s], mass[s+1:]...)
	}
}
