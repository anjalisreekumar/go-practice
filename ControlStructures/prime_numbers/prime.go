package main

import "fmt"

func checkIfPrime(num int) bool {
if num < 2 {
	return false
}
for i := 2 ; i*i <= num; i++ {
if num%i == 0 {
	return false
} 
}
return true
}	


func printAllPrime() {

	for num := 1 ; num <= 50 ; num ++ {
	if checkIfPrime(num) {
		fmt.Printf("%d ",num)
	}
	}

}

func main() {
	fmt.Printf("Prime number between 1 and 50 are: \n")
	printAllPrime()
}