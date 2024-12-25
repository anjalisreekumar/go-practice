package main

import "fmt"

func sumOfDigits(num int) int {

	var sum int

	for num > 0 {

		digit := num % 10
		sum += digit
		num /= 10

	}
	return sum
}

func main() {
	var num int
	fmt.Println("Enter your number: ")
	fmt.Scanf("%d",&num)
	sum := sumOfDigits(num)
	fmt.Printf("Sum is %d ", sum)
}
