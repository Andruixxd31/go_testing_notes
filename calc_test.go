package calc_test

import (
    "testing"
    "github.com/andruixxd31/go-testing-bible"
)

type TestCase struct{
    value int
    expected bool
    actual bool
}

func TestCalculateIsArmstrong(t *testing.T){
    testCase := TestCase{
        value: 371,
        expected: true,
    }

    testCase.actual = calc.CalculateIsArmstrong(testCase.value) 
    if testCase.actual != testCase.expected{
        t.Fail()
    }
}
