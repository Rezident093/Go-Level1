package fibbonachi

import (
	"errors"
	"testing"
)

type Xyz struct {
	name          string
	input         int
	useOptimized  bool
	output        int
	expectedError error
}

func TestFibbonachi(t *testing.T) {
	testCases := []Xyz{
		{
			name:   "sim test 10",
			input:  10,
			output: 55,
		},
		{
			name:         "sim test 11",
			input:        11,
			useOptimized: true,
			output:       89,
		},
		{
			name:   "case with 0",
			input:  0,
			output: 0,
		},
		{
			name:          "case with negative number",
			input:         -1,
			expectedError: ErrNegativeNumbers,
		},
	}

	for _, cse := range testCases {
		cse := cse
		t.Run(cse.name, func(t *testing.T) {
			result, err := Fibbonachi(cse.input, cse.useOptimized)
			if !errors.Is(err, cse.expectedError) {
				t.Errorf("function didn't return error")
			}
			if cse.expectedError == nil && result != cse.output {
				t.Errorf("reterns wrong answer: expected %d got %d", cse.output, result)
			}

		})
	}
}
