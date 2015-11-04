// Implement an algorithm to find the kth to last element of a singly linked list.

package q02

import (
	"encoding/json"
	"github.com/serdmanczyk/ccint-go/datastructures/linkedlist"
	"io/ioutil"
	"reflect"
	"runtime"
	"testing"
)

type TestCase struct {
	Input    string `json:"input"`
	K        int    `json:"k"`
	KthValue int    `json:"kthvalue"`
}

type q02Function func(*linkedlist.LinkedList, int) *linkedlist.Node

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

func nameOf(f interface{}) string {
	v := reflect.ValueOf(f)
	if v.Kind() == reflect.Func {
		if rf := runtime.FuncForPC(v.Pointer()); rf != nil {
			return rf.Name()
		}
	}
	return v.String()
}

func TestKthToLastUnknownLength(t *testing.T) {
	edgeCases(t, KthToLastUnknownLength)
	tableTests(t, KthToLastUnknownLength)
}

func TestKthToLastUnknownLengthAlternative(t *testing.T) {
	edgeCases(t, KthToLastUnknownLengthAlternative)
	tableTests(t, KthToLastUnknownLengthAlternative)
}

func TestKthToLastKnownLength(t *testing.T) {
	edgeCases(t, KthToLastKnownLength)
	tableTests(t, KthToLastKnownLength)
}

func edgeCases(t *testing.T, testFunc q02Function) {
	l := linkedlist.New()

	n := testFunc(l, 1)
	if n != nil {
		t.Errorf("Failure with out of bounds K, zero length")
	}

	l.Add(1)

	n = testFunc(l, 2)
	if n != nil {
		t.Errorf("Failure with out of bounds K, non-zero length")
	}

	n = testFunc(l, 0)
	if n != nil {
		t.Errorf("Failure with invalid K of 0")
	}

	n = testFunc(l, -1)
	if n != nil {
		t.Errorf("Failure with negative K")
	}
}

func tableTests(t *testing.T, testFunc q02Function) {
	funcName := nameOf(testFunc)

	for _, test := range tests {
		l := linkedlist.FromJSON(test.Input)
		KthNode := testFunc(l, test.K)
		if test.KthValue != KthNode.Value {
			t.Errorf("%s(%s) failed\n\tgot node with value %d, expected %d", funcName)
		}
	}
}
