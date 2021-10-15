package generators

import (
	"math/rand"
	"time"
)

type Task_3 struct {
	Mass_on, Mass_size, Default_mass_size int
	Mass                                  []int
}

func (t *Task_3) SetSettings() *Task_3 {
	t.Default_mass_size = 100000
	t.Mass_on = 0
	t.Mass_size = 0
	return t
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
