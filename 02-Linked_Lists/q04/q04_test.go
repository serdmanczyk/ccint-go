// Implement an algorithm to find the kth to last element of a singly linked list.

package q04

import (
	"encoding/json"
	"github.com/serdmanczyk/ccint-go/datastructures/linkedlist"
	"io/ioutil"
	"testing"
)

type TestCase struct {
	Input string `json:"input"`
	PartitionValDest int `json:"partitionvaldest"`
	Expected string `json:"expected"`
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

func TestPartitionLinkedList(t *testing.T) {
	for _, test := range tests {
		l := linkedlist.FromJSON(test.Input)
		val := l.NodeAtPos(test.PartitionValDest).Value
		PartitionLinkedList(l, val)
		output := l.JSON()
		if output != test.Expected {
			t.Errorf("Partition(*LinkedList) failed partitioning around value %d\n\tgot %s, expected %s", val, output, test.Expected)
		}
	}
}
