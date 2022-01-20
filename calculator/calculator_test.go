package calculator_test

import (
	"testing"

	"github.com/Delaram-Gholampoor-Sagha/testing_in_go/calculator"
)

type TestCase struct {
	name     string
	value    int
	expected bool
	actual   bool
}

// a few helpful commands :
// go test --cover
// go test -coverprofile=coverage.out
// go tool cover -html=coverage.out

func TestCalculateIsArmstrong(t *testing.T) {

	t.Run("test for all 3 digit armstrog numbers", func(t *testing.T) {
		tests := []TestCase{
			TestCase{name: "Testing value for: 153", value: 153, expected: true},
			TestCase{name: "Testing value for: 370", value: 370, expected: true},
			TestCase{name: "Testing value for: 371", value: 371, expected: true},
			TestCase{name: "Testing value for: 407", value: 407, expected: true},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				actual := calculator.CalculateIsArmstrong(test.value)
				if test.expected != actual {
					t.Fail()
				}
			})
		}

	})

}

func TestNegativeCalculateIsArmstrong(t *testing.T) {

	t.Run("should fail for 350", func(t *testing.T) {
		testCase := TestCase{
			value:    350,
			expected: false,
		}

		testCase.actual = calculator.CalculateIsArmstrong(testCase.value)
		if testCase.actual != testCase.expected {
			t.Fail()
		}

	})
	t.Run("should fail for 350", func(t *testing.T) {
		testCase := TestCase{
			value:    300,
			expected: false,
		}

		testCase.actual = calculator.CalculateIsArmstrong(testCase.value)
		if testCase.actual != testCase.expected {
			t.Fail()
		}

	})

}

func benchmarkCalculateIsArmstring(input int , b *testing.B) {
	for i := 0; i < b.N; i++ {
           calculator.CalculateIsArmstrong(input)
	}
}


// go test -run=Benchmark -bench=.

func BenchmarkCalculateIsArmstring370(b *testing.B) {
	 benchmarkCalculateIsArmstring(370 , b)
}
func BenchmarkCalculateIsArmstring371(b *testing.B) {
	 benchmarkCalculateIsArmstring(371 , b)
}
func BenchmarkCalculateIsArmstring0(b *testing.B) {
	benchmarkCalculateIsArmstring(0 , b)
}
