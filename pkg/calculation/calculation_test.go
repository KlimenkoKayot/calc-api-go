// TODO
// Нужно сделать 100% покрытие тестами файла calculation_test.go

package calculation

import "testing"

type Test struct {
	name       string
	expression string
	expected   float64
	err        error
}

func TestCalc(t *testing.T) {
	cases := []Test{
		{
			name:       "Good test #1",
			expression: "(2+2)*2",
			expected:   8,
		},
		{
			name:       "Good test #2",
			expression: "2+2*2",
			expected:   6,
		},
		{
			name:       "Good test #3",
			expression: "2+2*2+3*(5+1)*2/2/5/1+3/4",
			expected:   10.35,
		},
		{
			name:       "Good test #4",
			expression: "(   (2  )  )",
			expected:   2,
		},
		{
			name:       "Good test #5",
			expression: "((2-2)+1)/2",
			expected:   0.5,
		},
		{
			name:       "Error test #1",
			expression: "(()",
			err:        ErrInvalidExpression,
		},
		{
			name:       "Error test #2",
			expression: "g",
			err:        ErrInvalidSymbolExpression,
		},
		{
			name:       "Error test #3",
			expression: "2/0",
			err:        ErrDevisionByZero,
		},
		{
			name:       "Error test #4",
			expression: "2/0",
			err:        ErrDevisionByZero,
		},
		{
			name:       "Error test #5",
			expression: "",
			err:        ErrEmptyExpression,
		},
		{
			name:       "Error test #6",
			expression: "+)",
			err:        ErrBadReversedPoland,
		},
		{
			name:       "Error test #7",
			expression: "-+-2+",
			err:        ErrInvalidExpression,
		},
		{
			name:       "Error test #7",
			expression: "2?3",
			err:        ErrInvalidSymbolExpression,
		},
		{
			name:       "Error test #8",
			expression: "2+2(2)(2)(20)",
			err:        ErrInvalidExpression,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := Calc(tc.expression)
			if err != tc.err {
				t.Fatalf("Calc(\"%s\"), error: %v, expected: %v", tc.expression, err, tc.err)
			}
			if got != tc.expected {
				t.Fatalf("Calc(\"%s\"), got: %v, expected: %v", tc.expression, got, tc.expected)
			}
		})
	}
}
