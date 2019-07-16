package string

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumCalculator_Add(t *testing.T) {
	ast := assert.New(t)
	stringSumCalculator := NewSumCalculator()

	testCases := []struct {
		description    string
		numbersStr     string
		expectedResult int
		expectErr      bool
	}{
		{
			description:    "Test empty string.",
			numbersStr:     "",
			expectedResult: 0,
		},
		{
			description:    "Test the `,` delimiter.",
			numbersStr:     "1,2,5",
			expectedResult: 8,
		},
		{
			description:    "Test the new line` delimiter.",
			numbersStr:     "1,\n2,4",
			expectedResult: 7,
		},
		{
			description:    "Test a custom delimiter.",
			numbersStr:     "//$\n1$2$3",
			expectedResult: 6,
		},
		{
			description:    "Test a negative number.",
			numbersStr:     "//$\n1$-213$3",
			expectedResult: 0,
			expectErr:      true,
		},
		{
			description:    "Test a large number.",
			numbersStr:     "//$\n1$2,1001",
			expectedResult: 3,
		},
		{
			description:    "Test multiple customer delimiters.",
			numbersStr:     "//$,@,#\n1$2@3#4",
			expectedResult: 10,
		},
		{
			description:    "Test invalid delimiters.",
			numbersStr:     "//$,@,#\n1$**2@3#4",
			expectedResult: 0,
			expectErr:      true,
		},
	}

	for _, c := range testCases {
		result, err := stringSumCalculator.Add(c.numbersStr)
		t.Logf("Test case: %s\n", c.description)
		if err != nil {
			if !c.expectErr {
				t.Errorf("unexpected error %s.", err.Error())
			} else {
				t.Logf("expected error: %s", err.Error())
			}
		} else {
			t.Logf("expected result %d, actual result %d\n", c.expectedResult, result)
			ast.Equal(c.expectedResult, result)
			t.Logf("##################################################\n")
		}
	}
}

// TODO: Add unit tests for `recursiveAdd` and `parseDelimiters` private methods.
