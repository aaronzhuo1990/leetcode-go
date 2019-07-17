package string

// Calculator is the calculator for performing math for the numbers in a string.
type Calculator interface {
	Add(numbers string) (int, error)
}

// calculator is the implementation of the `Calculator` interface.
type calculator struct {
}

// NewCalculator creates an instance of `Calculator`.
func NewCalculator() Calculator {
	return &calculator{}
}
