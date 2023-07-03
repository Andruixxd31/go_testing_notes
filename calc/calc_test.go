// +build unit
package calc_test

import (
	"github.com/andruixxd31/go-testing-bible/calc"
	"testing"
)

type TestCase struct {
	name     string
	value    int
	expected bool
	actual   bool
}

func TestCalculateIsArmstrong(t *testing.T) {
	t.Run("Test for all 3 digits of armstrong", func(t *testing.T) {
		tests := []TestCase{
            TestCase{name: "Test for value: 370", value: 370, expected: true},
			TestCase{name: "Test for value: 371", value: 371, expected: true},
			TestCase{name: "Test for value: 407", value: 407, expected: true},
			TestCase{name: "Test for value: 153", value: 153, expected: true},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				test.actual = calc.CalculateIsArmstrong(test.value)
				if test.actual != test.expected {
					t.Fail()
				}
			})
		}

	})
}

func TestNegativeCalculateIsArmstrong(t *testing.T) {
	t.Run("Should fail for 350", func(t *testing.T) {
		testCase := TestCase{
			value:    350,
			expected: false,
		}

		testCase.actual = calc.CalculateIsArmstrong(testCase.value)
		if testCase.actual != testCase.expected {
			t.Fail()
		}
	})

	t.Run("Should fail for 300", func(t *testing.T) {
		testCase := TestCase{
			value:    300,
			expected: false,
		}

		testCase.actual = calc.CalculateIsArmstrong(testCase.value)
		if testCase.actual != testCase.expected {
			t.Fail()
		}
	})
}
