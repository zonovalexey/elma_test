package generators

import (
	"math/rand"
	"time"
)

type Task_2 struct {
	User_params, Mass_on, Mass_size, El_size, Default_mass_size, Default_el_size int
	Mass                                                                         []int
}

func (t *Task_2) SetSettings() *Task_2 {
	t.Default_mass_size = 1000000
	t.Default_el_size = 1000000000
	t.User_params = 0
	t.Mass_on = 0
	t.Mass_size = 0
	t.El_size = 0
	return t
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
