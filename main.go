package main

import (
	"fmt"

	"frontendmasters.com/go/io/data"
)

func calculateTax(price float32) (float32, float32) {
	return price * 0.09, price * 0.02
}

func birthday(age *int) {
	*age++
}

func main() {
	stateTax, cityTax := calculateTax(100)
	fmt.Println(stateTax, cityTax)

	age := 22
	birthday(&age)

	fmt.Println(&age)

	fmt.Println(data.MaxSpeed)
	printData()
}

// call anywhere inside module
// copy method does not work
// it does not oop
// static typed ( compiled ) language
