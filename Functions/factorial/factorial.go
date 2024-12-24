package main

import "fmt"

func factorial(num int) {
	fact := 1

	for i := 1 ; i <=num ; i++ {
		fact = fact * i
	}
	fmt.Printf("Factorial of %d is %d ", num , fact)
}

func main() {
	var num int
	fmt.Printf("Enter your number: ")
	fmt.Scanf("%d", &num)
	factorial(num)
}