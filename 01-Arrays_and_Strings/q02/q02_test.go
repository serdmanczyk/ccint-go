// Implement a function void reverse(char* str) in C or C++ which reverses a null-
//  terminated string.

package q02

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

type TestCase struct {
	Input    string `json:input`
	Expected string `json:expected`
}

var tests []TestCase

func init() {
	dat, err := ioutil.ReadFile("q02_testdata.json")
	if err != nil {
		panic(err)
	}

	if err = json.Unmarshal(dat, &tests); err != nil {
		panic(err)
	}
}

func TestReverse(t *testing.T) {
	for _, testcase := range tests {
		word, expected := testcase.Input, testcase.Expected
		if res := Reverse(testcase.Input); res != testcase.Expected {
			t.Errorf("Reverse(%s) failed; expected %s got %s", word, expected, res)
		}
	}
}
