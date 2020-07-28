package main

import "fmt"

type Character interface {
	Attack()
}

/* Magician */

type Magician struct {
}

func (m *Magician) Attack() {

	fmt.Println("The magician throws a big fire ball")

}

/* Magician */

/* Archer */

type Archer struct{}

func (a *Archer) Attack() {

	fmt.Println("The archer throws his arrow")

}

/* Archer */

/* Warrior */

type Warrior struct{}

func (a *Warrior) Attack() {

	fmt.Println("The warrior uses his sword")

}

/* Warrior */

type Factory struct {
	ch Character
}

func (f *Factory) Attack() {
	f.ch.Attack()
}

func main() {

	typeCharacter := "warrior"

	character := &Factory{}

	switch typeCharacter {
	case "archer":
		character = &Factory{
			ch: &Archer{},
		}
	case "magician":
		character = &Factory{
			ch: &Magician{},
		}
	case "warrior":
		character = &Factory{
			ch: &Warrior{},
		}
	}

	character.Attack()

}
