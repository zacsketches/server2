package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http/httptest"
	"os"
	"regexp"
	"testing"
)

func TestUpdateParams(t *testing.T) {
	//create a buffer to catch log messages from the program
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(os.Stderr)
	}()

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
		test{n: "badString", val: 43.0},
	}
	for _, test := range tests {

		qs := fmt.Sprintf("http://localhost:8000/?%s=%v", test.n, test.val)
		r := httptest.NewRequest("GET", qs, nil)
		err := r.ParseForm()
		if err != nil {
			t.Error("Unable to parse form")
		}
		testingParams.updateParams(r)
		switch test.n {
		case "cycles":
			if testingParams.cycles != test.val {
				t.Errorf("updateParams failed to detect cycles change. \n\tHave: %f\tWant:%f\n", testingParams.cycles, test.val)
			}
		case "res":
			if testingParams.res != test.val {
				t.Errorf("updateParams failed to detect res change. \n\tHave: %f\tWant:%f\n", testingParams.res, test.val)
			}
		case "phaseshift":
			if testingParams.phaseshift != test.val {
				t.Errorf("updateParams failed to detect phaseshift change. \n\tHave: %f\tWant:%f\n", testingParams.phaseshift, test.val)
			}
		case "yFreq":
			if testingParams.yFreq != test.val {
				t.Errorf("updateParams failed to detect yFreq change. \n\tHave: %f\tWant:%f\n", testingParams.yFreq, test.val)
			}
		default:
			testString := "updateParams: request for unknown parmater: " + test.n
			if !regexMatch(buf.String(), testString) {
				t.Errorf("updateParams failed to log bad parmaater string:\n\tHave: %s\tWant: %s\n", buf.String(), testString)
			}
		}
	}
}

// regexMatch test to see if the logEntry matches the test string s without
// inluding the log timestamp format.
// Assumes a log timestamp in the format yyyy/mm/dd hh:mm:ss
func regexMatch(log string, test string) bool {
	// re := regexp.MustCompile(`\d+/\d+/\d+ \d+:\d+:\d+ (.*)`)
	re := regexp.MustCompile(`\d+ \d+:\d+:\d+ (.*)`)
	matches := re.FindStringSubmatch(log)
	if matches == nil {
		panic("Unable to match the log timstamp to the regexp")
	}
	if len(matches) < 2 {
		panic("There is no log content beyond the timestamp: " + log)
	}
	return matches[1] == test
}
