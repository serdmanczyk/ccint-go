package q05

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
	dat, err := ioutil.ReadFile("q05_testdata.json")
	if err != nil {
		panic(err)
	}

	if err = json.Unmarshal(dat, &tests); err != nil {
		panic(err)
	}
}

func TestCompress(t *testing.T) {
	for _, testCase := range tests {
		output := Compress(testCase.Input)
		if output != testCase.Expected {
			t.Errorf("Compress(%s) failed; expected %v got %v", testCase.Input, []rune(testCase.Expected), []rune(output))
		}
	}
}
