package settings

type Task_1 struct {
	User_shift, Mass_size, Default_mass_size int
	Mass                                     []int
}

func (t Task_1) SetSettings() Task_1 {
	t.Default_mass_size = 100
	t.Mass_size = 0
	t.User_shift = -1
	return t
}

type Task_2 struct {
	User_params, Mass_on, Mass_size, El_size, Default_mass_size, Default_el_size int
	Mass                                                                         []int
}

func (t Task_2) SetSettings() Task_2 {
	t.Default_mass_size = 1000000
	t.Default_el_size = 1000000000
	t.User_params = 0
	t.Mass_on = 0
	t.Mass_size = 0
	t.El_size = 0
	return t
}

type Task_3 struct {
	Mass_on, Mass_size, Default_mass_size int
	Mass                                  []int
}

func (t Task_3) SetSettings() Task_3 {
	t.Default_mass_size = 100000
	t.Mass_on = 0
	t.Mass_size = 0
	return t
}

type Task_4 struct {
	Mass_on, Mass_size, El_size, Default_mass_size, Default_el_size int
	Mass                                                            []int
}

func (t Task_4) SetSettings() Task_4 {
	t.Default_mass_size = 100000
	t.Default_el_size = 1000000000
	t.Mass_on = 0
	t.Mass_size = 0
	t.El_size = 0
	return t
}
