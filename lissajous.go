// Lissajous.go contains types and methods to implement the lissajous
// figures returned to the server.
package main

import (
	"image"
	"image/color/palette"
	"image/gif"
	"io"
	"log"
	"math"
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

// var palette = []color.Color{color.White, color.Black}
var p = palette.Plan9

const (
	whiteIndex = 0 //first color in the palette
	blackIndex = 1 //second color in the palette
)

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
				log.Println("Unable to parse float from reuest string")
			}
			ljp.res = f
		case k == "phaseshift":
			f, err := strconv.ParseFloat(v[0], 64)
			if err != nil {
				log.Println("Unable to parse float from request string")
			}
			ljp.phaseshift = f
		case k == "yFreq":
			f, err := strconv.ParseFloat(v[0], 64)
			if err != nil {
				log.Println("Unable to parse float from request string")
			}
			ljp.yFreq = f
		default:
			log.Println("updateParams: request for unknown parmater:", k)
		}
	}

}

func (ljp *lissajousParameters) write(out io.Writer) {
	const (
		size    = 100 //image canvas covers [-size...+size]
		nframes = 128 //number of animation frames
		delay   = 8   //delay between frames in 10ms units
	)
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 //phase difference of x and y oscillators which grows
	//by ljp.phaseshift each loop
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, p)
		for t := 0.0; t < ljp.cycles*2*math.Pi; t += ljp.res {
			x := math.Sin(t)
			y := math.Sin(t*ljp.yFreq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8(i)+32)
		}
		phase += ljp.phaseshift
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) //Note: ignoring encoding errors
}
