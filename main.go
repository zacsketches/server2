// Server2 echoes the http request on localhost port 8000.
package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
)

const invalidCycles = math.MaxInt64

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Favicon requested")
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	fmt.Fprintf(w, "The form contains %d query string values.\n", len(r.Form))

	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}

	//test for a specific key in the form
	cycles, err := parseCycles(r)
	if err != nil {
		log.Println("Unable to parse query string for cycles. ", err)
	}
	if cycles != invalidCycles {
		fmt.Fprintf(w, "Run %d cycles!\n", cycles)
	}
}
