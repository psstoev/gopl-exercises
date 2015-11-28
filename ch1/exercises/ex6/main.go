package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"time"
)

const (
	blackIndex = 0 // first color in palette
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	if len(os.Args) > 1 && os.Args[1] == "web" {
		handler := func(w http.ResponseWriter, r *http.Request) {
			lissajous(w)
		}
		http.HandleFunc("/", handler)
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}
	lissajous(os.Stdout)
}

func rainbowPalette() []color.Color {
	// Stolen from http://www.poirrier.be/~jean-etienne/info/clibs/gd-rainbow.php
	
	var i uint8
	var palette []color.Color

	palette = make([]color.Color, 255, 255)
	
	for i = 0; i < 255; i++ {
		palette[i] = color.Black;
	}

	for i = 0; i < 42; i++ {
		// Line 1: red = 255 ; green = 0 -> 255 ; blue = 0
		palette[i] = color.RGBA{255, i * 6, 0, 0xff}
		// Line 2: red = 255 -> 0 ; green = 255 ; blue = 0
		palette[i+42] = color.RGBA{(255 - i*6), 255, 0, 0xff}
		// Line 3: red = 0 ; green = 255 ; blue = 0 -> 255
		palette[i+84] = color.RGBA{0, 255, i * 6, 0xff}
		// Line 4: red = 0 ; green = 255 -> 0 ; blue = 255
		palette[i+126] = color.RGBA{0, (255 - i*6), 255, 0xff}
		// Line 5: red = 0 -> 255 ; green = 0 ; blue = 255
		palette[i+168] = color.RGBA{i * 6, 0, 255, 0xff}
		// Line 6: red = 255 ; green = 0 ; blue = 255 -> 0
		palette[i+210] = color.RGBA{255, 0, (255 - i*6), 0xff}
	}
	
	palette[0] = color.Black

	return palette
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	palette := rainbowPalette()
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				uint8(int(t*10)%255))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
