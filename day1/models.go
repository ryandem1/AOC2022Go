package day1

type Food struct {
	calories int
}

type Elf struct {
	food []Food // Food being carried by an Efl
}

func (elf *Elf) totalCaloriesHeld() (caloriesHeld int) {
	for _, food := range elf.food {
		caloriesHeld += food.calories
	}
	return caloriesHeld
}
