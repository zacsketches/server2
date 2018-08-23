package main

import (
	"fmt"
	"net/http/httptest"
	"testing"
)

// Test of correctly formed input
func TestParseCycles(t *testing.T) {
	tests := map[string]int{
		"123":      123,
		"-123":     -123,
		"12345678": 12345678,
	}
	for k, v := range tests {
		qs := fmt.Sprintf("http://localhost:8000/?bob=allison&tony=hawk&cycles=%s", k)
		r := httptest.NewRequest("GET", qs, nil)
		err := r.ParseForm()
		if err != nil {
			t.Fatal("Error parsing form in TestParseCycles")
		}
		cycles, err := parseCycles(r)
		t.Logf("TestParseCycles: cycles: %v\terr: %v", cycles, err)
		if cycles != v || err != nil {
			t.Error("parseCycles: fail to detect correct Cycles on input: ", r.URL.String())
		}
	}
}

func TestParseCycles_noCyclesInQuery(t *testing.T) {
	tests := []string{
		"http://localhost:8000/?skater=true&tony=hawk",
		"http://localhost:8000/?speed=55&capability=ICantDriveFiftyFive",
		"http://localhost:8000/?jet=yes&type=F18",
	}

	for _, s := range tests {
		r := httptest.NewRequest("GET", s, nil)
		err := r.ParseForm()
		if err != nil {
			t.Fatal("Error parsing form in TestParseCycles_noCyclesInQuery")
		}
		cycles, err := parseCycles(r)
		t.Logf("TestParseCycles_noCyclesInQuery: cycles: %v\terr: %v", cycles, err)
		if cycles != invalidCycles && err == nil {
			t.Error("parseCycles: failed to throw error on input: ", r.URL.RawQuery)
		}
	}
}

func TestParseCycles_badInputString(t *testing.T) {
	tests := []string{
		"notAnInt",
		"123.4",
		"0.3",
		"char23char",
		"23char",
		"char23",
	}
	for _, s := range tests {
		qs := fmt.Sprintf("http://localhost:8000/?skater=true&cycles=%s&tony=hawk", s)
		r := httptest.NewRequest("GET", qs, nil)
		err := r.ParseForm()
		if err != nil {
			t.Fatal("Error parsing form in TestParseCycles_badInputString")
		}
		cycles, err := parseCycles(r)
		t.Logf("TestParseCycles_badInputString: cycles: %v\terr: %v", cycles, err)
		if cycles != invalidCycles && err == nil {
			t.Error("parseCycles: failed to throw error on input: ", r.URL.RawQuery)
		}
	}
}
