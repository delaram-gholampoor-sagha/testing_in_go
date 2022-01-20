package calculator_test

import (
	"testing"

	"github.com/Delaram-Gholampoor-Sagha/testing_in_go/calculator"
)

type TestCase struct {
	value    int
	expected bool
	actual   bool
}

func TestCalculateIsArmstrong(t *testing.T) {
	testCase := TestCase{
		value:    371,
		expected: true,
	}

	testCase.actual = calculator.CalculateIsArmstrong(testCase.value)
	if testCase.actual != testCase.expected {
		t.Fail()
	}
}

func TestNegativeCalculateIsArmstrong(t *testing.T) {
	testCase := TestCase{
		value:    372,
		expected: false,
	}

	testCase.actual = calculator.CalculateIsArmstrong(testCase.value)
	if testCase.actual != testCase.expected {
		t.Fail()
	}
}
