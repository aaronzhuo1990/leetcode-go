package main

import (
	"fmt"
)

func main() {
	result := partition("aabc")
	fmt.Printf("result %v\n", result) // Should be [["aa", "b"], ["a", "a", "b"]]
}

func partition(s string) [][]string {
	result := make([][]string, 0)
	current := make([]string, 0, len(s))
	searchPalindromes(s, 0, current, &result)
	return result
}

func searchPalindromes(str string, strIndex int, curPalindromes []string, result *[][]string) {
	if strIndex == len(str) {
		// This means curPalindromes is the partition we want when the search reaches the end of the string.
		*result = append(*result, curPalindromes)
		return
	}

	for index := strIndex; index < len(str); index++ {
		if ifPalindrome(str[strIndex : index+1]) {
			searchPalindromes(str, index+1, append(curPalindromes, str[strIndex:index+1]), result)
		}
	}
}

// ifPalindrome detects whether the given string is a palindrome.
func ifPalindrome(str string) bool {
	if len(str) <= 1 {
		return true
	}
	begin, end := 0, len(str)-1
	for begin < end {
		if str[begin] != str[end] {
			return false
		}
		begin++
		end--
	}
	return true
}
