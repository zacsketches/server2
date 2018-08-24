// Server2 echoes the http request on localhost port 8000.
package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
)

const invalidCycles = math.MaxInt64

var myParams = lissajousParameters{
	cycles:     5,
	res:        0.0001,
	phaseshift: 0.1,
	yFreq:      2.5,
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Favicon requested")
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	//parse the query string
	err := r.ParseForm()
	if err != nil {
		log.Print("ParseForm: ", err)
	}
	//test for the lissajous paramater keys in the query string
	myParams.updateParams(r)

	//test for a specific key in the form
	cycles, err := parseCycles(r)
	if err != nil {
		log.Println("Unable to parse query string for cycles. ", err)
	}
	if cycles != invalidCycles {
		fmt.Fprintf(w, "Run %d cycles!\n", cycles)
	}
}
