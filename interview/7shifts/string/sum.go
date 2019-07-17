package string

import (
	"fmt"
	"strconv"
	"strings"
)

// Add - calculate the sum of the numbers in the given string.
//	@params:
//		numbers: The string that contains both numbers and custom delimiters.
//	@return:
//		int: The sum of the numbers in the given string.
//		error: An error which indicates something goes wrong.
func (s *calculator) Add(numbers string) (int, error) {
	delimiters, leftStr := s.parseDelimiters(numbers)
	return s.recursiveAdd(leftStr, delimiters, 0)
}

// recursiveAdd - add all the numbers in the given string recursively.
//	@params:
//		numbers: The string that contains both numbers and custom delimiters.
// 		delimiters: All the delimiters used in the given string.
//	@return:
//		int: The sum of the numbers in the given string.
//		error: An error which indicates something goes wrong.
func (s *calculator) recursiveAdd(numbers string, delimiters map[string]bool, sum int) (int, error) {
	numStr := make([]byte, 0, 10) // Declare a byte slice which init ial capacity is 10.

	// Basic line
	if len(numbers) == 0 {
		return sum, nil
	}

	if '0' <= numbers[0] && numbers[0] <= '9' {
		// numbers found
		index := 0

		for ; index < len(numbers) && '0' <= numbers[index] && numbers[index] <= '9'; index++ {
			numStr = append(numStr, numbers[index])
		}
		num, err := strconv.Atoi(string(numStr))
		if err != nil {
			return 0, fmt.Errorf("error converting string %s to integer, err: %s", string(numStr), err.Error())
		}
		if num <= 1000 {
			// Only sum the number which is not large than 1000
			sum += num
		}
		return s.recursiveAdd(numbers[index:], delimiters, sum) // Handle the rest of the input string.

	} else {
		// A non-number character found
		delimiter := make([]rune, 0, 10) // Declare a rune (unicode char) slice which initial capacity is 10.

		for index, char := range numbers {
			switch {
			case char == '-' && '0' <= numbers[index+1] && numbers[index+1] <= '9':
				// `-` character found
				numIndex := index + 1
				for ; numIndex < len(numbers) && '0' <= numbers[numIndex] && numbers[numIndex] <= '9'; numIndex++ {
					numStr = append(numStr, numbers[numIndex])
				}
				return 0, fmt.Errorf("negatives not allowed, negative number: [-%s]", string(numStr))

			case '0' <= char && char <= '9':
				// Number character found
				if len(delimiter) != 0 && !delimiters[string(delimiter)] {
					return 0, fmt.Errorf("invalid delimiter [%s] found", string(delimiter))
				} else {
					return s.recursiveAdd(numbers[index:], delimiters, sum) // Handle the rest of the input string.
				}

			default:
				// Other characters
				delimiter = append(delimiter, char)
				if delimiters[string(delimiter)] {
					// A delimiter found, clean up the delimiter slice.
					delimiter = delimiter[:0]
				}
			}

		}
	}

	return sum, nil
}

// parseDelimiters - parse delimiters from the given string
//	@params:
//		numbers: the string that contains both numbers and custom delimiters
//	@return:
//		map[string]bool: A string-bool pair which represents all the delimiters
//		string: a part of the input string without customer delimiters
//	TODO: handle edge cases.
func (s *calculator) parseDelimiters(numbers string) (map[string]bool, string) {
	// The following delimiters are default delimiters.
	allDelimiters := map[string]bool{
		"\n": true,
		",":  true,
	}
	index := 0 // The first index of the input string after filtering out custom delimiters.

	if strings.HasPrefix(numbers, "//") {
		delimiter := make([]rune, 0, 10)

		// Parse custom delimiters from the rest of the string.
		for left, char := range numbers[2:] {
			switch char {
			case rune(','), rune('\n'):
				if len(delimiter) != 0 {
					allDelimiters[string(delimiter)] = true // Add the customer delimiter to the map.
					delimiter = delimiter[:0]               // Clean up the slice.
				}

				if char == rune('\n') {
					index = left + 2 + 1 // Skip the part from `\\' to '\n`.
					break                // No more custom delimiter.
				}

			default:
				delimiter = append(delimiter, char)
			}
		}
	}

	return allDelimiters, numbers[index:]
}
