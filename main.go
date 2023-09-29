package main

import (
	"fmt"
	"log"
)

func main() {
	OddEven(8888)
	OddEven(9999)

	var radius, Area float64
	log.Print("enter the radius of a CIRCLE:")
	fmt.Scanln(&radius)
	Area = areaOfCircle(radius)
	log.Println("The Area of a Circle whose radius is", radius, "is", Area, "cm^2.")
}

// Check if given number is odd or even
func OddEven(num int) {
	if num%2 == 0 {
		log.Println("Given number", num, " is even number")
	} else {
		log.Println("Given number", num, " is odd number")
	}
}

func areaOfCircle(radius float64) float64 {
	return (22 / 7.0) * (radius * radius)
}
