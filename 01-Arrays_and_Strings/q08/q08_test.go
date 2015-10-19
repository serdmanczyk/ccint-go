package q08

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

type TestCase struct {
	WordOne    string `json:wordone`
	WordTwo    string `json:wordtwo`
	IsRotation bool   `json:isrotation`
}

var tests []TestCase

func init() {
	dat, err := ioutil.ReadFile("q08_testdata.json")
	if err != nil {
		panic(err)
	}

	if err = json.Unmarshal(dat, &tests); err != nil {
		panic(err)
	}
}

func TestCompress(t *testing.T) {
	for _, testCase := range tests {
		result := IsRotation(testCase.WordOne, testCase.WordTwo)
		if result != testCase.IsRotation {
			t.Errorf("IsRotation(%s, %s) failed; expected %b got %b", testCase.WordOne, testCase.WordTwo, testCase.IsRotation, result)
		}
	}
}
