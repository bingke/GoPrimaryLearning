// Lissajous generates GIF animations of random Lissajous figures.
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

var palette = []color.Color{color.White, color.RGBA{0, 255, 25, 255}, color.RGBA{0, 5, 5, 255}, color.RGBA{0, 255, 6, 7}, color.RGBA{0, 12, 68, 100}}

func main1_6() {

	rand.Seed(time.Now().UTC().UnixNano())
	lissajous(os.Stdout)

}

func lissajous(out io.Writer) {
	const (
		cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)

	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}

	phase := 0.0

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(freq*t + phase)
			img.SetColorIndex(size+int(size*x+0.5), size+int(size*y+0.5), byte(rand.Intn(4))) //use rand to choose different color from palette
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)

}
