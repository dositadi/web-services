package main

import (
	"fmt"
)

type Duration int

func (d Duration) Hours() float64 {
	return float64(d) / 60.0
}

func (d Duration) IsLongerThan(other Duration) bool {
	return d > other
}

func (d *Duration) SetDuration(minutes int) {
	*d = Duration(minutes)
}

type Temperature float64

func (t Temperature) Fahreinheit() Temperature {
	return (t * 9 / 5) + 32
}

func (t Temperature) IsFreezing() bool {
	return t <= 0
}

func main() {
	duration := Duration(10)
	hours := duration.Hours()
	isLongerThan := duration.IsLongerThan(Duration(5))

	fmt.Println("Hours: ", hours, "\nIs Longer: ", isLongerThan)

	duration.SetDuration(20)

	fmt.Println("Current Duration: ", duration)

	temperature := Temperature(25.5)
	fahreinheit := temperature.Fahreinheit()
	IsFreezing := temperature.IsFreezing()

	fmt.Println("Fahreinheit: ", fahreinheit, "\nIs Freezing: ", IsFreezing)
}
