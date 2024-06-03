package main

import "fmt"

type IProduct interface {
	getName() string
	getPrice() float32
}

// base

type Product struct {
	name  string
	price float32
}

func (p Product) getName() string {
	return p.name
}

func (p Product) getPrice() float32 {
	return p.price
}

// END base

// wrapper / decorator

type ProductBlackFriady50 struct {
	product IProduct
}

func (p ProductBlackFriady50) getName() string {
	return fmt.Sprintf("%s %s", p.product.getName(), "50% sale")
}

func (p ProductBlackFriady50) getPrice() float32 {
	return p.product.getPrice() * 0.5
}

// END wrapper / dcorator

// aux function to emulate an use scenario

func purchase(product IProduct) {
	fmt.Printf("Paid %0.2f \n", product.getPrice())
}

// END aux function

func main() {

	p := Product{
		name:  "Yu-Gi-Oh tin",
		price: 27.99,
	}

	pDiscount := ProductBlackFriady50{
		product: p,
	}

	fmt.Printf("normal %s: %f \n", p.getName(), p.getPrice())
	fmt.Printf("black friday %s: %f \n", pDiscount.getName(), pDiscount.getPrice())

	purchase(pDiscount)

}
