package main

import "log"

func main() {
	OddEven(8888)
	OddEven(9999)   
}

//Check if given number is odd or even
func OddEven (num int) {
	if(num%2==0){
		log.Println("Given number",num," is even number")
	} else {
		log.Println("Given number",num," is odd number")
	}
}
