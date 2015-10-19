// Given an image represented by an NxN matrix, where each pixel in the image is 4
//  bytes, write a method to rotate the image by 90 degrees. Can you do this in place?

package q06

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
	dat, err := ioutil.ReadFile("q06_testdata.json")
	if err != nil {
		panic(err)
	}

	if err = json.Unmarshal(dat, &tests); err != nil {
		panic(err)
	}
}

func TestMatrixRotate(t *testing.T) {
	for _, testCase := range tests {
		var matrix Matrix
		err := json.Unmarshal([]byte(testCase.Input), &matrix)
		if err != nil {
			panic(err)
		}

		matrix.Rotate()

		result, err := json.Marshal(matrix)
		if err != nil {
			panic(err)
		}

		output := string(result)

		if output != testCase.Expected {
			t.Errorf("matrix.Rotate() failed with input %s; expected %s got %s", testCase.Input, testCase.Expected, output)
		}
	}
}
