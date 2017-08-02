
package numberutil

import "testing"

type testCases struct{
	value1 float64
	value2 float64
	result float64
}

var addtests = []testCases{
	{ 1, 2, 3.0 },
	{ 1, 0, 1.0 },
	{ -1, 1, 0.0 },
	{ 0, 0, 0.0},
	{ 0.111111111111111111, 1, 1.111111111111111111 },
	{ -1, 3.9406564584, 2.9406564584 },
}

var subtracttests = []testCases{
	{ 1, 2, -1.0 },
	{ 987, 0, 987.0 },
	{ 0, 1, -1.0 },
	{ 0, 0, 0.0},
	{ -1, 1, -2.0 },
	{ 1.11, 1, 0.1100000000000001 },
	{ -1, -3.9406564584, 2.9406564584 },
}


func TestAdd(t *testing.T){

	for _,	testcase := range addtests {
		v, err := Add(testcase.value1, testcase.value2)
		if v != testcase.result || err != nil {
			t.Error(
				"For", testcase.value1,
				"and", testcase.value2,
				"expected", testcase.result,
				"got",v,
			)
		}
	}
}

func TestSubtract(t *testing.T){

	for _,	testcase := range subtracttests {
		v, err := Subtract(testcase.value1, testcase.value2)
		if v != testcase.result || err != nil {
			t.Error(
				"For", testcase.value1,
				"and", testcase.value2,
				"expected", testcase.result,
				"got",v,
			)
		}
	}
}
