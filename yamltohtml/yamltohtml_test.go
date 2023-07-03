// +build unit
package yamltohtml_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/andruixxd31/go-testing-bible/yamltohtml"
)

type TestCase struct {
	desc     string
	path     string
	expected string
}

func TestMain(m *testing.M){
    fmt.Println("Running tests of yamltohtml")
    ret := m.Run()
    fmt.Println("Tests Executed")
    os.Exit(ret)
}

func TestYamlToHTML(t *testing.T) {
	testCases := []TestCase{
		{
			desc:     "Tests title is set properly",
			path:     "testdata/test_01.yml",
			expected: "<html><head><title>My Awesome Page</title></head><body>This is my awesome body</body></html>",
		},
		{
			desc:     "Tests body is set properly",
			path:     "testdata/test_02.yml",
			expected: "<html><head><title>My Second Page</title></head><body>This is my awesome second body</body></html>",
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			result, err := yamltohtml.YamlToHTML(test.path)
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

