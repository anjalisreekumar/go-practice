package main
import "fmt"

func reverseString(str string) []rune {
	 chars := []rune(str)
	
	 for i,j := 0,len(str) -1 ; i < j ; i, j = i+1 ,j-1 {
		chars[i],chars[j] = chars[j] ,chars[i]
	 }

	 return chars
}

func checkIfPalindrome(str string) bool {
	if str == string(reverseString(str)) {
		return true
	} else {
		return false
	}
}

func main(){
	var str string
	fmt.Printf("Enter your string: ")
	fmt.Scanf("%s",&str)
	if checkIfPalindrome(str) {
		fmt.Printf("Your string %s is Palindrome", str)
	} else {
		fmt.Printf("Your string %s is not Palindrome", str)
	}
}