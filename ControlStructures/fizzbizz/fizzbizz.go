package main

import "fmt"


func fizzOrBizz() {
	num := 1

	for num <=100 {
		if num%3 == 0 && num%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if num%3 == 0 {
			fmt.Println("Fizz")
		} else if num%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(num)
		}
		num ++
	}
	
}
func main(){
fizzOrBizz()
}

