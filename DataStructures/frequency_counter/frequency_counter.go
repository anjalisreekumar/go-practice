package main

import "fmt"

func findStringOccurance(slice []string) map[string]int {
	occuranceMap := make(map[string]int)

	for _,str := range slice {
		occuranceMap[str] ++
	}
	return occuranceMap
}

func main(){
	var size int
	var strArr []string
	fmt.Printf("Enter size of array :")
	fmt.Scanf("%d",&size)
	fmt.Printf("Enter your string array :")
	 // Read input strings into the slice
	 for i := 0; i < size; i++ {
        var input string
        fmt.Scanf("%s", &input)
        strArr = append(strArr, input)
    }
fmt.Printf("Occurance of each string in array is : \n")
for str, num := range findStringOccurance(strArr) {
    fmt.Println(str, num)
}

}