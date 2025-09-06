package main

import "fmt"

func CanFormPalindrome(s string) bool {
	charCount := make(map[rune]int)

	for _, char := range s {
		charCount[char]++
	}

	oddCount := 0
	for _, count := range charCount {
		if count%2 != 0 {
			oddCount++
		}
	}

	return oddCount <= 1
}

func main() {
	var s = "code"
	fmt.Printf("%v", CanFormPalindrome(s))
}
