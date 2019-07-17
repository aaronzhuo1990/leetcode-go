package main

import (
	"fmt"
	"log"

	"github.com/azhuox/leetcode-go/interview/7shifts/string"
)

func main() {
	strCalculator := string.NewCalculator()
	inputStr := "1,2,3,4"

	result, err := strCalculator.Add(inputStr)
	if err != nil {
		log.Fatalf("Error summing all the numbers in the string %s, err: %s", inputStr, err.Error())
	}

	fmt.Printf("The sum of all the numbers in the string [%s] is %d\n", inputStr, result)

}
