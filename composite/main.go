package main

import "fmt"

type IComponent interface {
	calculateCost() float32
}

type Trip struct {
	components []IComponent
	name       string
}

func (t Trip) calculateCost() float32 {

	var cost float32

	fmt.Printf("-- %s \n", t.name)
	for _, component := range t.components {
		cost = cost + component.calculateCost()
	}

	fmt.Printf("---- -> %0.2f \n\n", cost)
	return cost

}

func (t *Trip) addComponent(newComponents ...IComponent) {
	t.components = append(t.components, newComponents...)
}

type Details struct {
	cost float32
	name string
}

func (d Details) calculateCost() float32 {
	fmt.Printf("--- %s -> %0.2f \n", d.name, d.cost)
	return d.cost
}

func main() {

	mediterranean := Trip{
		name: "Mediterranean trip",
	}
	egypt := Trip{
		name: "Egypt",
	}
	egypt.addComponent(Details{
		name: "food",
		cost: 200,
	}, Details{
		name: "cruise",
		cost: 1000,
	}, Details{
		name: "souvenirs",
		cost: 25,
	})

	greece := Trip{
		name: "Greece",
	}
	greece.addComponent(Details{
		name: "food",
		cost: 250,
	}, Details{
		name: "cruise",
		cost: 750,
	}, Details{
		name: "souvenirs",
		cost: 70,
	},
		Details{
			name: "museums",
			cost: 40,
		})

	mediterranean.addComponent(egypt, greece)

	fmt.Printf("The price for mediterranean trip is %0.2f \n", mediterranean.calculateCost())

}
