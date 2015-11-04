
package q07

import (
	"github.com/serdmanczyk/ccint-go/datastructures/linkedlist"
	"testing"
)

type TestCase struct {
	input string
	expected bool
}

func TestIsPalindrome(t *testing.T) {
	for _, test := range []TestCase{
		{"[1,2,3,2,1]", true},
		{"[1,2,3,2,1,2]", false},
		{"[1,2,2,1]", true},
		{"[1,2,3,4,5,4,3,2,1]", true},
		{"[1,2,3,4,5,5,4,3,2,1]", true},
		{"[1,2,3,4,5,5,6,4,3,2,1]", false},
	}{
		l := linkedlist.FromJSON(test.input)
		result := IsPalindrome(l)
		if result != test.expected {
			t.Errorf("IsPalindrome(%s) failed, expected %b got %b", test.input, test.expected, result)
		}
	}
}