package main

import (
	"fmt"
	"strings"
)

func main() {
	path := "/a//b////c/d//././/.."

	fmt.Printf("result %s\n", simplifyPath(path)) // "/a/b/c"
}

func simplifyPath(path string) string {
	length := len(path)
	stack := make([]string, 0, length/2)
	validDir := make([]byte, 0, length)

	for i := 0; i < length; i++ {
		// Remove the data in validDir
		validDir = validDir[:0]

		for i < length && path[i] != '/' {
			validDir = append(validDir, path[i])
			i++
		}

		dirStr := string(validDir)
		switch dirStr {
		case ".", "":
			// Do nothing
		case "..":
			if len(stack) > 0 {
				// Pop
				stack = stack[:len(stack)-1]
			}
		default:
			// Push
			stack = append(stack, dirStr)
		}
	}

	return "/" + strings.Join(stack, "/")
}
