package main

import "fmt"

func largestOfThree(a int,b int ,c int) int {
	if a >= b && a >= c {
		return a 
	} else if b >= a && b >= c {
		return b
	} else {
		return c
	}
}

func main() {
var a,b,c int

fmt.Printf("Enter three numbers (space seperated) :")
fmt.Scanf("%d %d %d", &a, &b, &c )

	fmt.Printf("Largest of three numbers %d %d and %d is %d \n",a , b , c , largestOfThree(a,b,c))
}