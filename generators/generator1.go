package generators

import (
	"math/rand"
	"time"
)

type Task_1 struct {
	User_shift, Mass_size, Default_mass_size int
	Mass                                     []int
}

func (t *Task_1) SetSettings() *Task_1 {
	t.Default_mass_size = 100
	t.Mass_size = 0
	t.User_shift = -1
	return t
}

func (t Task_1) Generator() []int {
	default_el_size := 2000
	mass := make([]int, t.Mass_size)
	generator := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < t.Mass_size; i++ {
		mass[i] = generator.Intn(default_el_size) + 1 - 1000
	}
	return mass
}

// type Generator interface {
// 	Generator() []int
// }

// type Task interface {
// 	Solution() int
// 	SetSettings(Generator)
// }

// type Task123 struct {
// 	gen Generator
// }

// func (t Task123) Solution() int {
// 	return 1
// }

// func (t Task123) SetSettings(g Generator) {
// 	g.Generator()
// }
