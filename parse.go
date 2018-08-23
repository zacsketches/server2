package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

// parseCycles returns the number of cycles in the query string and an error value
func parseCycles(r *http.Request) (int, error) {
	s := r.Form.Get("cycles") //returns "" without query string 'cycles'
	if s == "" {
		errorString := fmt.Sprintf("no 'cycles' in query string and s==%q", s)
		return invalidCycles, errors.New(errorString)
	}
	c, err := strconv.Atoi(s)
	if err != nil {
		return invalidCycles, err
	}
	return c, nil
}
