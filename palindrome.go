package main

import "fmt"

func CheckIfPalindrome(s string) bool {
	j := len(s) - 1
	for i := 0; i < len(s)/2; i++ {
		k := j - i
		if s[i] != s[k] && k > j {
			return false
		}
	}
	//fmt.Println(s[0])

	return true
}

func main() {
	s := "abcba"

	result := CheckIfPalindrome(s)

	fmt.Println(result)
}
