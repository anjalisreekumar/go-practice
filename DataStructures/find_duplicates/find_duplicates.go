package main

import "fmt"

func findDuplicates(arr []int) []int {
var duplicates []int
seen := make(map[int]bool)
	for i := 0; i <= len(arr)-1; i++ {
		fmt.Printf("value of i is %d \n", i)
		fmt.Printf("length of arr is %d and value of arr[i] is %v \n", len(arr), arr[i])
		for j := i + 1 ; j < len(arr) ; j++ {
			fmt.Printf("length of arr is %d and value of arr[j] is %v \n", len(arr), arr[j])
			if arr[i] == arr[j] {
				fmt.Printf("Found identical element %v \n", arr[i])
				if !seen[arr[i]] {
				duplicates = append(duplicates,arr[i])
				fmt.Printf("Duplicates array is %v\n",duplicates)
				seen[arr[i]] = true
			}
			}
			
		}

	
	
}
return duplicates
}

func main() {
	var size int
	var arr []int
	fmt.Printf("Enter size of your array: \n")
	fmt.Scanf("%d", &size)
	fmt.Printf("size is %d \n", size)

	fmt.Printf("Enter elements of your array: \n")
	arr = make([]int, size)
	for i := 0; i <= size - 1; i++ {
		fmt.Scanf("%d", &arr[i])
	}
	duplicates := findDuplicates(arr)
	fmt.Printf("Final Duplicates: %v\n", duplicates)
}
