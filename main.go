package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"time"
)

func main() {
	OddEven(8888)
	OddEven(9999)

	var radius, Area float64
	log.Print("enter the radius of a CIRCLE:")
	fmt.Scanln(&radius)
	Area = areaOfCircle(radius)
	log.Println("The Area of a Circle whose radius is", radius, "is", Area, "cm^2.")

	findlargenumber()

	floor()

	rand.Seed(time.Now().UnixNano())
	randomNumber := generateRandomNumber(1, 100)
	fmt.Printf("Random number: %d\n", randomNumber)
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

func findlargenumber() {
	log.Println("Enter first number :")
	var first int
	fmt.Scanln(&first)
	log.Println("Enter second number :")
	var second int
	fmt.Scanln(&second)
	log.Println("Enter third number :")
	var third int
	fmt.Scanln(&third)
	/* check the boolean condition using if statement */
	if first >= second && first >= third {
		log.Println(first, "is Largest among three numbers.") /* if condition is true then print the following */
	}
	if second >= first && second >= third {
		log.Println(second, "is Largest among three numbers.")
	}
	if third >= first && third >= second {
		log.Println(third, "is Largest among three numbers")
	}
}

func floor() {
	var input float64
	var result float64
	log.Println("Enter a floating-point number: ")
	fmt.Scanf("%f", &input)

	result = math.Floor(input)

	log.Printf("The floor value is %.0f", result)
}

func generateRandomNumber(min, max int) int {
	return rand.Intn(max-min+1) + min
}
