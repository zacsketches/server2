package main

import (
	"fmt"
	"net/http/httptest"
	"testing"
)

func TestUpdateParams(t *testing.T) {
	var testingParams = lissajousParameters{
		cycles:     5.0,
		res:        0.0001,
		phaseshift: 0.1,
		yFreq:      2.5,
	}
	type test struct {
		n   string
		val float64
	}
	tests := []test{
		test{n: "cycles", val: 10.0},
		test{n: "res", val: 0.001},
		test{n: "phaseshift", val: 0.2},
		test{n: "yFreq", val: 4.0},
	}
	for _, test := range tests {
		qs := fmt.Sprintf("http://localhost:8000/?%s=%v", test.n, test.val)
		r := httptest.NewRequest("GET", qs, nil)
		err := r.ParseForm()
		if err != nil {
			t.Error("Unable to parse form")
		}
		testingParams.updateParams(r)
		if test.n == "cycles" && testingParams.cycles != test.val {
			t.Errorf("updateParams failed to detect cycles change. \n\tHave: %f\tWant:%f\n", testingParams.cycles, test.val)
		}
		if test.n == "res" && testingParams.res != test.val {
			t.Errorf("updateParams failed to detect res change. \n\tHave: %f\tWant:%f\n", testingParams.res, test.val)
		}
		if test.n == "phaseshift" && testingParams.phaseshift != test.val {
			t.Errorf("updateParams failed to detect phaseshift change. \n\tHave: %f\tWant:%f\n", testingParams.phaseshift, test.val)
		}
		if test.n == "yFreq" && testingParams.yFreq != test.val {
			t.Errorf("updateParams failed to detect yFreq change. \n\tHave: %f\tWant:%f\n", testingParams.yFreq, test.val)

		}
	}

	//phaseshift

	//check yFreq

}
