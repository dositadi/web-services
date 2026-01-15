package main

import (
	"fmt"
	"time"
)

// 1. Bank Account

type BankAccount struct {
	AccountHolder string
	Balance       float64
}

func (b *BankAccount) Deposit(amount float64) {
	if amount <= 0 {
		return
	} else {
		b.Balance += amount
	}
}

func (b *BankAccount) Withdraw(amount float64) {
	if amount <= 0 {
		fmt.Println("Amount must be above 0")
		return
	} else if b.Balance < amount {
		fmt.Println("Insufficient fund")
		return
	} else {
		b.Balance -= amount
	}
}

func (b BankAccount) GetBalance() float64 {
	return b.Balance
}

// 2. Logging

type LogMessage string

func (l *LogMessage) PrePendTime() {
	*l = LogMessage(time.Now().Format("2006-01-02 15:04:05")) + " " + *l
	fmt.Println("Time Prepended: ", *l)
}

func (l *LogMessage) UpperCase() {
	var temp LogMessage
	for _, i := range *l {
		if i >= 'a' && i <= 'z' {
			temp = temp + LogMessage(i-32)
		} else {
			temp = temp + LogMessage(rune(i))
		}
	}
	*l = temp
	fmt.Println("Log Message in UpperCase: ", *l)
	fmt.Println()
}

// 3. Rectangle Geometry

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Print() {
	fmt.Printf("Width: %v\nHeight: %v\n\n", r.Width, r.Height)
}

func (r Rectangle) Perimeter() float64 {
	return 2*r.Height + 2*r.Width
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (r *Rectangle) Scale(factor float64) {
	r.Height *= factor
	r.Width *= factor
}

func main() {
	// Bank Account
	bankAccount := BankAccount{
		AccountHolder: "Okipo",
		Balance:       20_000,
	}

	bal := bankAccount.GetBalance()
	fmt.Println("Balance: ", bal)
	bankAccount.Deposit(10_000)
	bal1 := bankAccount.GetBalance()
	fmt.Println("Current Balance: ", bal1)
	bankAccount.Deposit(150_000)
	bankAccount.Withdraw(5_000)
	bankAccount.Withdraw(20_000)
	bal2 := bankAccount.GetBalance()
	fmt.Println("Most Current Balance: ", bal2)

	fmt.Println()

	// Log Message
	logMessage := LogMessage("Little drops of water will eventually form a mighty ocean")
	fmt.Println(logMessage)
	logMessage.PrePendTime()
	logMessage.UpperCase()

	// Rectangle
	rectangle := Rectangle{
		Width:  6.3,
		Height: 12.6,
	}

	rectangle.Print()

	a := rectangle.Area()
	p := rectangle.Perimeter()
	fmt.Printf("Area: %v\nPerimeter: %v\n\n", a, p)

	rectangle.Scale(3)
	rectangle.Print()

	a1 := rectangle.Area()
	p1 := rectangle.Perimeter()
	fmt.Printf("Area: %v\nPerimeter: %v\n", a1, p1)
}
