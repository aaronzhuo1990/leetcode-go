# string package

This package exposes the following interface:

- Calculator: This interface contains the method for performing math for the numbers in a string. Currently it has the following methods:
 - Add: Sum all the numbers in a string and return it.

## Testing the Package

You can execute the following commands to run the unit tests in this package.  

```
go get github.com/azhuox/leetcode-go
cd $GOPATH/src/github.com/azhuox/leetcode-go/interview/7shifts/
go test ./string/... -v
```


## Using the Package

The following example shows how to use this package:

```
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
```

You can also run the following the commands to run the above example:

```
go get github.com/azhuox/leetcode-go
cd $GOPATH/src/github.com/azhuox/leetcode-go/interview/7shifts/
go run ./example/main.go
```
