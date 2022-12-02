package main

import (
	"fmt"
	"math/rand"
)

type creature interface {
	String() string
	Move() string
	Eat() string
}

type animal struct {
	name    string
	actions []string
	foods   []string
}

func (a animal) String() string {
	return a.name
}

func (a animal) Move() string {
	idx := rand.Intn(len(a.actions))
	return a.actions[idx]
}

func (a animal) Eat() string {
	idx := rand.Intn(len(a.foods))
	return a.foods[idx]
}

func chooseCreature(creature creature) string {
	i := rand.Intn(2) == 1
	if i {
		return fmt.Sprintf("%s is eating %s.", creature, creature.Eat())
	} else {
		return fmt.Sprintf("%s %s.", creature, creature.Move())
	}
}

func main() {
	hours := 24
	days := 3

	dog := animal{
		name:    "Shiba",
		actions: []string{"bark", "woof", "tail wave"},
		foods:   []string{"dog food", "bone", "meat"},
	}

	cat := animal{
		"Sphinix",
		[]string{"meow", "purr", "acting weird", "sprint"},
		[]string{"cat food", "pills", "dried shrimp"},
	}

	lion := animal{
		"Lion King",
		[]string{"roar", "haunt", "play"},
		[]string{"pray meat", "beast meat"},
	}

	capybara := animal{
		"Capybara",
		[]string{"chill", "soak in water", "slow walk"},
		[]string{"idk", "idk2", "grass maybe"},
	}

	animals := []animal{dog, cat, lion, capybara}

	for d := 0; d < days; d++ {
		for h := 0; h < hours; h++ {
			a := animals[rand.Intn(len(animals))]
			fmt.Print("day ", d, " hour ", h, ": ")
			if h < 6 || h > 18 {
				fmt.Println(a, "is sleeping.")
			} else {
				fmt.Println(chooseCreature(a))
			}
		}
	}
}
