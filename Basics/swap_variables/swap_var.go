package main

import "fmt"

func main() {
	var a,b uint
fmt.Printf("Enter your first number a: ")
fmt.Scanf("%d",&a)
fmt.Printf("Enter your second number b: ")
fmt.Scanf("%d",&b)

a,b = swapVariables(a,b)
fmt.Printf("Swapped values for a is %d and b is %d \n",a,b)
}

func swapVariables(a uint, b uint)( uint, uint) {	
	//a = 20 , b = 30
a = b + a  // a = 30+20 = 50
b = a - b //b = 50 - 30 = 20
a = a - b //a = 50 - 20 = 30
return a,b
}