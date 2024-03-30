package main

import (
	"fmt"

	"frontendmasters.com/go/io/data"
)

func calculateTax(price float32) (float32, float32) {
	return price * 0.09, price * 0.02
}

// pointer
func birthday(pointerAge *int) {
	if *pointerAge > 140 {
		painc("Too old to be true")
	}
	fmt.Printf("The pointer is %v and the value is %v\n", pointerAge, *pointerAge)
	*pointerAge++
}

func painc(s string) {
	panic("unimplemented")
}

func main() {
	// can write _ instead of name
	_, cityTax := calculateTax(100)
	fmt.Println(cityTax) // and don't need to  call _ (underscore)

	defer fmt.Println("Bye!")
	defer fmt.Println("Good ")

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

// only one type ==
// other operator => > < >= <= !=

// struct
