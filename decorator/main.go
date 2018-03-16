package main

import "fmt"

// Beverage defines coffee behavior
type Beverage interface {
	getDescription() string
	getCost() float32
}

// Mocha is a type of condiment for coffee.
type Mocha struct {
	cost        float32
	description string
	beverage    Beverage
}

func newMocha(b Beverage) Beverage {
	return Mocha{
		cost:        0.81,
		description: "Mocha",
		beverage:    b,
	}
}

func (m Mocha) getCost() float32 {
	return m.beverage.getCost() + m.cost
}

func (m Mocha) getDescription() string {
	return fmt.Sprintf("%s, %s", m.beverage.getDescription(), m.description)
}

// HouseBlend is a type of base coffee.
type HouseBlend struct {
	cost        float32
	description string
}

func newHouseBlend() Beverage {
	return HouseBlend{
		cost:        1.19,
		description: "HouseBlend",
	}
}

func (h HouseBlend) getCost() float32 {
	return h.cost
}

func (h HouseBlend) getDescription() string {
	return h.description
}

func main() {
	h := newHouseBlend()
	fmt.Println(h.getCost())

	h = newMocha(h)
	fmt.Println(h.getCost())
	fmt.Println(h.getDescription())
}
