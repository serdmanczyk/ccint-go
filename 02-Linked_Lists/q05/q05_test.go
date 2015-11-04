package q05

import (
	// "encoding/json"
	"github.com/serdmanczyk/ccint-go/datastructures/linkedlist"
	// "io/ioutil"
	"testing"
)

type TestCase struct {
	listone, listtwo, expected string
}

func TestSumLinkedLists(t *testing.T) {
	for _, test := range []TestCase{
		{"[1,2,3,4]", "[5,6,7,8]", "[6,8,0,3,1]"},
		{"[1,2,3,4,5]", "[5,6,7,8]", "[6,8,0,3,6]"},
		{"[1,2,3,4]", "[5,6,7,8,6]", "[6,8,0,3,7]"},
		{"[9,9,9,9]", "[9,9,9,9]", "[8,9,9,9,1]"},
	} {
		l1 := linkedlist.FromJSON(test.listone)
		l2 := linkedlist.FromJSON(test.listtwo)
		sum := SumLinkedLists(l1, l2)
		if sum.JSON() != test.expected {
			t.Errorf("Addition of %s and %s failed\n\t got %s expected%s", l1.JSON(), l2.JSON(), sum.JSON(), test.expected)
		}
	}
}

func TestSumLinkedListsReverse(t *testing.T) {
	for _, test := range []TestCase{
		{"[1,2,3,4]", "[5,6,7,8]", "[6,9,1,2]"},
		{"[1,2,3,4,5]", "[5,6,7,8]", "[1,8,0,2,3]"},
		{"[1,2,3,4]", "[5,6,7,8,6]", "[5,8,0,2,0]"},
		{"[9,9,9,9]", "[9,9,9,9]", "[1,9,9,9,8]"},
	} {
		l1 := linkedlist.FromJSON(test.listone)
		l2 := linkedlist.FromJSON(test.listtwo)
		sum := SumLinkedListsReverse(l1, l2)
		if sum.JSON() != test.expected {
			t.Errorf("Addition of %s and %s failed\n\t got %s expected%s", l1.JSON(), l2.JSON(), sum.JSON(), test.expected)
		}
	}
}
