package main 


import "fmt" 

func main() {
var radius float32
	fmt.Printf("Enter radius of circle:")
	fmt.Scan(&radius)
	fmt.Printf("Area of circle with radius %.2f is %.2f\n", radius, calculateArea(radius))
}


func calculateArea(radius float32) float32 {
	const pi = 3.14
	area := pi * radius*radius

	return area
}