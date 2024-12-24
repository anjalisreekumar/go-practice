package main

import "fmt"

func main() {
	var num int
fmt.Printf("Enter your number:")
fmt.Scanf("%d",&num)
fmt.Printf("Your number %d is %s \n", num,checkIfOddOrEven(num))
}

func checkIfOddOrEven(number int) string {
	if number%2 == 0 {
		return "even"
	} else {
		return "odd"
	}
}