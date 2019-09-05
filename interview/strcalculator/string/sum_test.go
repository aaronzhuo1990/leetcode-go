package string

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumCalculator_Add(t *testing.T) {
	ast := assert.New(t)
	strCalculator := NewCalculator()

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
			expectErr:      false,
		},
		{
			description:    "Test the `,` delimiter.",
			numbersStr:     "1,2,5",
			expectedResult: 8,
			expectErr:      false,
		},
		{
			description:    "Test the new line delimiter.",
			numbersStr:     "1,\n2,4",
			expectedResult: 7,
			expectErr:      false,
		},
		{
			description:    "Test a custom delimiter.",
			numbersStr:     "//$\n1$2$3",
			expectedResult: 6,
			expectErr:      false,
		},
		{
			description:    "Test a negative number.",
			numbersStr:     "//$\n1$-213$3",
			expectedResult: 0,
			expectErr:      true,
		},
		{
			description:    "Test a large number.",
			numbersStr:     "//$\n1$2,10001",
			expectedResult: 3,
			expectErr:      false,
		},
		{
			description:    "Test multiple customer delimiters.",
			numbersStr:     "//$,@,#\n1$2@3#4",
			expectedResult: 10,
			expectErr:      false,
		},
		{
			description:    "Test invalid delimiters.",
			numbersStr:     "//$,@,#\n1$**2@3#4",
			expectedResult: 0,
			expectErr:      true,
		},
		{
			description:    "Test unicode delimiters.",
			numbersStr:     "//$,@,你好\n1$你好2@3@4", // '你好' means 'hello' in Chinese
			expectedResult: 10,
			expectErr:      false,
		},
	}

	for _, c := range testCases {
		result, err := strCalculator.Add(c.numbersStr)
		t.Logf("Test case: %s\n", c.description)
		if err != nil {
			if !c.expectErr {
				t.Errorf("unexpected error %s.", err.Error())
			} else {
				t.Logf("expected error: %s", err.Error())
			}
		} else {
			inputStr := strings.Replace(c.numbersStr, "\n", "\\n", -1)
			t.Logf("input string: \"%s\" \texpected result %d, actual result %d\n", inputStr, c.expectedResult, result)
			ast.Equal(c.expectedResult, result)
		}
		t.Logf("##################################################\n")
	}
}

// TODO: Add unit tests for `recursiveAdd` and `parseDelimiters` private methods.
