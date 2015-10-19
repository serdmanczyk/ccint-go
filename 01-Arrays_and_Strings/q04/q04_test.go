// Write a method to replace all spaces in a string with'%%20'. You may assume that
//  the string has sufficient space at the end of the string to hold the additional
//  characters, and that you are given the 'true' length of the string. (Note: if
//  imple- menting in Java, please use a character array so that you can perform this
//  opera- tion in place.) EXAMPLE Input: 'Mr John Smith Output: 'Mr%%20Dohn%%20Smith'

package q04

import (
	"encoding/json"
	"io/ioutil"
	"strings"
	"testing"
)

// https://github.com/btford/philosobot
type TestCase struct {
	Input    string `json:input`
	Expected string `json:expected`
}

var tests []TestCase

func init() {
	dat, err := ioutil.ReadFile("q04_testdata.json")
	if err != nil {
		panic(err)
	}

	if err = json.Unmarshal(dat, &tests); err != nil {
		panic(err)
	}
}

// Turn test string into an array of runes that' the input function can use
// Returns string cast into an array of runes that's large enough to
// replace the spaces with '%20' alongside the length of the input data
func Runify(s string) ([]rune, int) {
	spaces := strings.Count(s, " ")
	arr := make([]rune, len(s)+spaces*2)
	copy(arr, []rune(s))
	return arr, len(s)
}

func TestEncodeSpaces(t *testing.T) {
	for _, testCase := range tests {
		input, strlen := Runify(testCase.Input)
		EncodeSpaces(input, strlen)
		result := string(input)
		if result != testCase.Expected {
			t.Errorf("EncodeSpaces(%s) failed; expected %s got %s", testCase.Input, testCase.Expected, result)
		}
	}
}
