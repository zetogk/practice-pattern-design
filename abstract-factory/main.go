package main

import (
	"fmt"
	"strconv"
)

//// Abstract clases

// Printer ...
type Printer interface {
	Print(string)
}

// Toner ...
type Toner interface {
	CheckLevel()
}

//// Concrete

//// HP

// HpPrinter ...
type HpPrinter struct {
	reference string
}

// Print ...
func (p HpPrinter) Print(text string) {

	fmt.Println("The text '" + text + "' is printed by the HP printer " + p.reference)

}

// HpToner ...
type HpToner struct {
	level int
}

// CheckLevel ...
func (t HpToner) CheckLevel() {

	fmt.Println("The current level is " + strconv.Itoa(t.level))

}

//// Epson

// EpsonPrinter ...
type EpsonPrinter struct {
	reference string
}

// Print ...
func (p EpsonPrinter) Print(text string) {

	fmt.Println("The text " + text + " is printed by the Epson printer " + p.reference)

}

// EpsonToner ...
type EpsonToner struct {
	level int
}

// CheckLevel ...
func (t EpsonToner) CheckLevel() {

	fmt.Println("The current level is " + strconv.Itoa(t.level))

}

//// Abstract factory

// FactoryPrinters ...
type FactoryPrinters interface {
	CreatePrinter() Printer
	CreateToner() Toner
}

//// Concrete factories

// HpFactory ...
type HpFactory struct {
}

// CreatePrinter ...
func (f HpFactory) CreatePrinter() Printer {
	return HpPrinter{reference: "hp001"}
}

// CreateToner ...
func (f HpFactory) CreateToner() Toner {
	return HpToner{
		level: 100,
	}
}

// EpsonFactory ...
type EpsonFactory struct {
}

// CreatePrinter ...
func (f EpsonFactory) CreatePrinter() Printer {
	return EpsonPrinter{reference: "ep001"}
}

// CreateToner ...
func (f EpsonFactory) CreateToner() Toner {
	return EpsonToner{
		level: 100,
	}
}

// Application ...
type Application struct {
	factory FactoryPrinters
}

// CreateFactory ...
func (a *Application) CreateFactory(brand string) {

	fmt.Println("input brand: ", brand)
	var factory FactoryPrinters

	switch brand {
	case "hp":
		fmt.Println("HP is the choosen brand")
		factory = &HpFactory{}
		break

	case "epson":
		fmt.Println("Epson is the choosen brand")
		factory = &EpsonFactory{}
		break

	default:
		fmt.Println("HP is the choosen brand by default")
		factory = &HpFactory{}
		break
	}

	a.factory = factory

}

func main() {

	app := Application{}
	app.CreateFactory("hp")

	printer := app.factory.CreatePrinter()
	printer.Print("Testing the app")

	toner := app.factory.CreateToner()
	toner.CheckLevel()

}
