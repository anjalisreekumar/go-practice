package main

import "fmt"

func arrayReverse(arr []int) []int {

	for i,j := 0,len(arr) - 1 ; i < j ; i,j = i+1 , j-1 {
		arr[i],arr[j] = arr[j],arr[i]
	}
	return arr
}

func main(){
	
	var num int
	fmt.Printf("Enter the number of elements in the array: ")
	fmt.Scanf("%d", &num)
	 arr := make([]int,num)
	 fmt.Printf("Enter %d elements:\n", num)
	 for i := 0 ; i<num ; i++ {
	 fmt.Scan(&arr[i])
}

	fmt.Printf("Reversed array is %v ", arrayReverse(arr))
}