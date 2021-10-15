package generators

import (
	"math/rand"
	"time"
)

type Task_4 struct {
	Mass_on, Mass_size, El_size, Default_mass_size, Default_el_size int
	Mass                                                            []int
}

func (t *Task_4) SetSettings() *Task_4 {
	t.Default_mass_size = 100000
	t.Default_el_size = 1000000000
	t.Mass_on = 0
	t.Mass_size = 0
	t.El_size = 0
	return t
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
