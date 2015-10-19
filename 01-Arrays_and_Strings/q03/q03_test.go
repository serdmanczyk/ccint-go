// Given two strings, write a method to decide if one is a
// permutation of the other.

package q03

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

type TestCase struct {
	Strone   string `json:"strone"`
	Strtwo   string `json:"strtwo"`
	Expected bool   `json:"expected"`
}

var tests []TestCase

func init() {
	dat, err := ioutil.ReadFile("q03_testdata.json")
	if err != nil {
		panic(err)
	}

	if err = json.Unmarshal(dat, &tests); err != nil {
		panic(err)
	}
}

func TestIsPermutation(t *testing.T) {
	for _, testcase := range tests {
		one, two, expected := testcase.Strone, testcase.Strtwo, testcase.Expected
		if res := IsPermutation(one, two); res != expected {
			t.Errorf("IsPermutation(%s, %s) failed; expected %t got %t", one, two, expected, res)
		}
	}
}
