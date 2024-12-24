package main

import "fmt" 

func main() {
var a , b int
fmt.Print("Enter your first digit: ")
fmt.Scan(&a)
fmt.Print("Enter your second digit: ")
fmt.Scan(&b)
fmt.Printf("Sum of %d and %d is %d", a,b,sumOfTwo(a,b))
}

func sumOfTwo(a int, b int) int {
	return a+b
}