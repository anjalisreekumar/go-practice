package main

import (
	"errors"
	"fmt"
	"os"
)

func divideNumbers(num1 , num2 float32) (float32,error) {
	if num2 == 0 {
		return 0, errors.New("number not divisible by 0")
	}
	val := num1/num2

	return val , nil

}

func main() {
	var num1 , num2 float32
fmt.Println("Enter your first and second number :")
fmt.Scanf("%f",&num1)
fmt.Scanf("%f",&num2)

val, err := divideNumbers(num1,num2)

if err != nil {
	fmt.Printf("%s", err.Error())
	os.Exit(-1)
} 
fmt.Println("Result of division is " , val)

}