package yamltohtmlgo_test

import (
	"fmt"
	"os"
	"testing"

	yamltohtmlgo "github.com/Delaram-Gholampoor-Sagha/testing_in_go/yamltohtml"
)

type TestCase struct {
	desc     string
	path     string
	expected string
}

func TestMain(m *testing.M) {
	fmt.Println("Hello world")
	ret := m.Run()
	fmt.Println("Tests have executed")
	os.Exit(ret)
}

func TestYamlToHTML(t *testing.T) {
	testCases := []TestCase{
		TestCase{
			desc:     "Test Case 1",
			path:     "testdata/test_01.yaml",
			expected: "<html><head><title>My Awesome Page</title></head><body>This is my fantastic content</body></html>",
		},
		TestCase{
			desc:     "Test Case 2",
			path:     "testdata/test_02.yaml",
			expected: "<html><head><title>My Second Page</title></head><body>This is my fantastic content</body></html>",
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			result, err := yamltohtmlgo.YamlToHTML(test.path)
			if err != nil {
				t.Fail()
			}
			t.Log(result)
			if result != test.expected {
				t.Fail()
			}
		})
	}

}
