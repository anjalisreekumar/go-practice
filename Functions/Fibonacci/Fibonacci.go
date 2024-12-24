package main

import "fmt"

func printFibonacciSeries(num int) {
	a, b := 0, 1

	if num > 0 {
		fmt.Printf("%d ", a)
	}
	if num > 1 {
		fmt.Printf("%d ", b)
	}

	for i := 2; i <= num; i++ {
		sum := a + b
		fmt.Printf("%d ", sum)
		a = b
		b = sum
	}

}

func main() {
	var num int
	fmt.Printf("Enter your limit: ")
	fmt.Scanf("%d",&num)
	fmt.Printf("Fibonacci Series upto %d numbers are: ", num)
	printFibonacciSeries(num)

}
