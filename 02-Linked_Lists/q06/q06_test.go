package q06

import (
	"github.com/serdmanczyk/ccint-go/datastructures/linkedlist"
	"testing"
	"reflect"
	"runtime"
)

type TestCase struct {
	input string
	intersectionpos, expectedval int
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

func TestFindLoopUnknownLength(t *testing.T) {
	runTestSuite(FindLoopUnknownLength, t)
}

func TestFindLoopKnownLength(t *testing.T) {
	runTestSuite(FindLoopKnownLength, t)
}

func runTestSuite(testFunc func(*linkedlist.LinkedList) *linkedlist.Node, t*testing.T) {
	funcName := nameOf(testFunc)

	for _, test := range []TestCase{
		{"[1,2,3,4]", 2, 3},
		{"[1,2,3,4,5,6]", 4, 5},
		{"[1,1]", 1, 1},
		{"[1]", 0, 1},
		{"[1,2,3,4,5,6,7,8,9,10,11,12,13,14,15]", 10, 11},
	}{
		l := linkedlist.FromJSON(test.input)
		n := l.NodeAtPos(test.intersectionpos)
		l.Head.Next = n
		res := testFunc(l)
		if res.Value != test.expectedval {
			t.Errorf("%s(%s) failed finding loop with intersection node located at node %d, expected node val %d got node val %d",
				funcName, test.input, test.intersectionpos, test.expectedval, res.Value)
		}
	}
}