// Implement an algorithm to find the kth to last element of a singly linked list.

package q03

import (
	"encoding/json"
	"github.com/serdmanczyk/ccint-go/datastructures/linkedlist"
	"io/ioutil"
	"testing"
)

type TestCase struct {
	Input    string `json:"input"`
	DelPos        int    `json:"delpos"`
	Expected string    `json:"expected"`
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

func TestDeleteNode(t *testing.T) {
	for _, test := range tests {
		l := linkedlist.FromJSON(test.Input)
		n := l.NodeAtPos(test.DelPos)
		DeleteNode(n)
		output := l.JSON()
		if output != test.Expected {
			t.Errorf("DeleteNode(*n) failed deleting node at %d of LinkedList length %d\n\tgot %s, expected %s", test.DelPos, l.Length, output, test.Expected)
		}
	}
}
