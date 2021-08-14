package main

import "fmt"

type Animal struct { //* parent objek
	Name     string
	Gender   string
	IsMammal bool
}

func (a *Animal) Speak() { //* parent behavior
	var text string

	text = fmt.Sprintf("I am %s. My gender is %s. I am not a mammal.", a.Name, a.Gender)
	if a.IsMammal {
		text = fmt.Sprintf("I am %s. My gender is %s. I am a mammal.", a.Name, a.Gender)
	}
	fmt.Println(text)
}

type Cat struct {
	Animal //* embeded struct
	Breed  string
}

func (a *Cat) Speak() { //* child behavior
	a.Animal.Speak()
	fmt.Printf("My breed is %s.\n", a.Breed)
}

type Whale struct {
	Animal
}

func (a *Whale) speak() {
	fmt.Println("WOOOOOOOOOOMMMM")
}

func polymorphism() {
	animalDiona := Animal{
		Name:     "Diona",
		Gender:   "female",
		IsMammal: true,
	}
	animalDiona.Speak()

	catKecing := Cat{
		Animal: Animal{
			Name:     "Kecing",
			Gender:   "female",
			IsMammal: true,
		},
		Breed: "Anggora",
	}
	catKecing.Speak()

	pausWhale := Whale{
		Animal: Animal{
			Name:     "Pauso",
			Gender:   "male",
			IsMammal: true,
		},
	}
	pausWhale.Speak()
}
