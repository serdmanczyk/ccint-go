package q07

import (
	"encoding/json"
	"io/ioutil"
	"reflect"
	"runtime"
	"testing"
)

type TestCase struct {
	Input    string `json:"input"`
	Expected string `json:"expected"`
}

var tests []TestCase

func init() {
	dat, err := ioutil.ReadFile("q07_testdata.json")
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

func TestZeroInPlace(t *testing.T) {
	RunTestSuite(t, ZeroInPlace)
}

func TestZeroLessWork(t *testing.T) {
	RunTestSuite(t, ZeroLessWork)
}

func RunTestSuite(t *testing.T, testFunc func(*Matrix)) {
	funcName := nameOf(testFunc)

	for _, testCase := range tests {
		var matrix Matrix

		err := json.Unmarshal([]byte(testCase.Input), &matrix)
		if err != nil {
			panic(err)
		}

		testFunc(&matrix)

		bites, err := json.Marshal(matrix)
		if err != nil {
			panic(err)
		}

		output := string(bites)
		if output != testCase.Expected {
			t.Errorf("%s(*Matrix) failed with input %s; expected %s got %s", funcName, testCase.Input, testCase.Expected, output)
		}
	}
}
