package main

import (
	"fmt"
)

func calculateArea() {
	var length, width float64
	fmt.Print("Enter the length of the rectangle: ")
	fmt.Scanln(&length)
	fmt.Print("Enter the width of the rectangle: ")
	fmt.Scanln(&width)

	fmt.Println("The area of the rectangle is: ", length*width)
}
