package main

import "fmt"

type Product struct {
	Name     string
	Price    float64
	Quantity int
}

// METHODS THAT RECIEVES VALUE BY COPY
// For Displaying the informations about a product
func (p Product) Display() {
	fmt.Printf("Product Name: %v,\nProduct Price: %v,\nProduct Quantity: %v\n", p.Name, p.Price, p.Quantity)
}

func (p Product) CalculateTotalPrice(number int) float64 {
	return p.Price * float64(number)
}

// METHODS THAT RECIEVES POINTER ARGUMENTS
func (p *Product) IncreaseQty(value int) {
	p.Quantity = p.Quantity + value
}

func (p *Product) UpdatePrice(newPrice float64) {
	p.Price = newPrice
}

func main() {
	product := &Product{
		Name:     "Hollandia Yoghurt",
		Price:    1450,
		Quantity: 100,
	}

	product.Display()
	amount := product.CalculateTotalPrice(10)
	fmt.Println("Total Price: ", amount)

	product.IncreaseQty(20)
	product.UpdatePrice(1550)
	product.Display()
}
