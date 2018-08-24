// Lissajous.go contains types and methods to implement the lissajous
// figures returned to the server.
package main

import (
	"log"
	"net/http"
	"strconv"
)

// LissajousParamaters are the configuration parameters affiliated with
// a lissajous figure
type lissajousParameters struct {
	cycles     float64
	res        float64
	phaseshift float64
	yFreq      float64
}

func (ljp *lissajousParameters) updateParams(r *http.Request) {
	// log.Printf("The form contains %d query string values.\n", len(r.Form))
	for k, v := range r.Form {
		switch {
		case k == "cycles":
			f, err := strconv.ParseFloat(v[0], 64)
			if err != nil {
				log.Println("Unable to parse float from request string")
			}
			ljp.cycles = f
		case k == "res":
			f, err := strconv.ParseFloat(v[0], 64)
			if err != nil {
				log.Println("Unable to parse float from output string")
			}
			ljp.res = f
		case k == "phaseshift":
			f, err := strconv.ParseFloat(v[0], 64)
			if err != nil {
				log.Println("Unable to parse float from output string")
			}
			ljp.phaseshift = f
		case k == "yFreq":
			f, err := strconv.ParseFloat(v[0], 64)
			if err != nil {
				log.Println("Unable to parse float from output string")
			}
			ljp.yFreq = f

		}
	}

}
