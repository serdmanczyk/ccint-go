package q01

import (
	"encoding/json"
	"io/ioutil"
	"reflect"
	"runtime"
	"testing"
)

type TestCase struct {
	Input    string `json:input`
	Expected bool   `json:expected`
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

func TestAllUniqueCharsAssumeASCII(t *testing.T) {
	runTestSuite(t, AllUniqueCharsAssumeASCII)
}

func TestAllUniqueCharsHashMap(t *testing.T) {
	runTestSuite(t, AllUniqueCharsHashMap)
}

func TestUniqueNoStructsOne(t *testing.T) {
	runTestSuite(t, AllUniqueCharsNoDataStructsOne)
}

func TestUniqueNoStructsTwo(t *testing.T) {
	runTestSuite(t, AllUniqueCharsNoDataStructsTwo)
}

func runTestSuite(t *testing.T, testFunc func(string) bool) {
	funcName := nameOf(testFunc)

	for i := 0; i < 10000; i++ {
		for _, testcase := range tests {
			word, expected := testcase.Input, testcase.Expected
			if res := testFunc(word); res != expected {
				t.Errorf("%s(%s) failed; expected %b got %b", funcName, word, expected, res)
			}
		}
	}
}
