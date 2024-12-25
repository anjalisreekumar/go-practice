package main

import "fmt"

func bubbleSort(arr []int) []int {
	len := len(arr)

	for i := 0; i < len-1; i++ {
		isSwapped := false
		for j := 0; j < len-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				isSwapped = true
			}
		}

		if !isSwapped {
			break
		}
	}
	return arr

}

func main(){
	var size int
	fmt.Printf("Enter the size of array: ")
	fmt.Scanf("%d", &size)
	fmt.Printf("Enter elements of array: ")

	arr := make([]int , size)
	for i := 0; i< size ; i++ {
		fmt.Scanf("%d",&arr[i])
	}
	fmt.Printf("Final sorted array is \n %v",bubbleSort(arr) )

}
