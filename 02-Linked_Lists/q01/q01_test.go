package q01

import (
	"encoding/json"
	"github.com/serdmanczyk/ccint-go/datastructures/linkedlist"
	"io/ioutil"
	"reflect"
	"runtime"
	"testing"
)

type TestCase struct {
	Input          string `json:input`
	Expected       string `json:expected`
	ExpectedLength int    `json:"expectedlength"`
}

var tests []TestCase

func init() {
	dat, err := ioutil.ReadFile("q01_testdata.json")
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

func TestRemoveDuplicates(t *testing.T) {
	runTestSuite(t, RemoveDuplicates)
}

func TestRemoveDuplicatesNoStruct(t *testing.T) {
	runTestSuite(t, RemoveDuplicatesNoStruct)
}

func runTestSuite(t *testing.T, testFunc func(l *linkedlist.LinkedList)) {
	funcName := nameOf(testFunc)

	for _, testcase := range tests {
		llist := linkedlist.FromJSON(testcase.Input)
		testFunc(llist)
		output := llist.JSON()
		if output != testcase.Expected {
			t.Errorf("%s(*LinkedList) failed\n\tinput:%s\n\texpected:%s\n\tgot %s", funcName, testcase.Input, testcase.Expected, output)
		}

		if llist.Length != testcase.ExpectedLength {
			t.Errorf("%s(*LinkedList) failed\n\tinput:%s\n\texpected length:%d\n\tgot length: %d", funcName, testcase.Input, testcase.ExpectedLength, llist.Length)
		}
	}
}
