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
    t.Run("Should return true for 371", func(t *testing.T) {
        testCase := TestCase{
            value: 371,
            expected: true,
        }

        testCase.actual = calc.CalculateIsArmstrong(testCase.value) 
        if testCase.actual != testCase.expected{
            t.Fail()
        }
    })


    t.Run("Should return true for 370", func(t *testing.T) {
        testCase := TestCase{
            value: 370,
            expected: true,
        }

        testCase.actual = calc.CalculateIsArmstrong(testCase.value) 
        if testCase.actual != testCase.expected{
            t.Fail()
        }
    })
}


func TestNegativeCalculateIsArmstrong(t *testing.T){
    t.Run("Should fail for 350", func(t *testing.T) {
        testCase := TestCase{
            value: 350,
            expected: false,
        }

        testCase.actual = calc.CalculateIsArmstrong(testCase.value) 
        if testCase.actual != testCase.expected{
            t.Fail()
        }
    })


    t.Run("Should fail for 300", func(t *testing.T) {
        testCase := TestCase{
            value: 300,
            expected: false,
        }

        testCase.actual = calc.CalculateIsArmstrong(testCase.value) 
        if testCase.actual != testCase.expected{
            t.Fail()
        }
    })
}
