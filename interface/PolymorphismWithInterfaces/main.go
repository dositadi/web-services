package main

import "fmt"

// Polymorphism using Interfaces

const PI float64 = 3.142

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Shape interface {
	Area() float64
	Perimeter() float64
	Scale(value float64)
}

type Calculate struct {
	shape Shape
}

func (c Calculate) ReturnArea() float64 {
	return c.shape.Area()
}

func (c Calculate) ReturnPerimeter() float64 {
	return c.shape.Perimeter()
}

func (c Calculate) Scale(value float64) {
	c.shape.Scale(value)
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Height + r.Width)
}

func (r *Rectangle) Scale(value float64) {
	r.Height = r.Height * value
	r.Width = r.Width * value
	fmt.Println("New Height: ", r.Height, "\nNew Width: ", r.Width)
}

func (c Circle) Area() float64 {
	return PI * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * PI * c.Radius
}

func (c *Circle) Scale(value float64) {
	c.Radius = c.Radius * value
	fmt.Println("New Radius: ", c.Radius)
}

func PrintArea(s Shape) {
	fmt.Println("The Area is: ", s.Area())
}

func PrintPerimeter(s Shape) {
	fmt.Println("The Perimeter is: ", s.Perimeter())
}

func main() {
	c := &Circle{
		Radius: 12.43,
	}

	r := &Rectangle{
		Width:  6.43,
		Height: 10.84,
	}

	calc1 := Calculate{shape: c}
	fmt.Println("Calculating for Circle")
	PrintArea(calc1.shape)
	PrintPerimeter(calc1.shape)

	fmt.Println()
	fmt.Println("Calculating for Rectangle")

	calc2 := Calculate{shape: r}
	PrintArea(calc2.shape)
	PrintPerimeter(calc2.shape)

	fmt.Println()
	fmt.Println("Scaling Circle")

	calc1.Scale(2)

	fmt.Println()
	fmt.Println("Scaling Rectangle")

	calc2.Scale(2)

	fmt.Println()
	fmt.Println("Re-Calculating for Circle")

	fmt.Println()
	fmt.Println("Re-Calculating for Rectangle")

	var shapes []Shape = []Shape{calc1.shape, calc2.shape}

	i := 0

	for i < len(shapes) {
		PrintArea(shapes[i])
		PrintPerimeter(shapes[i])
		i++
	}

}
