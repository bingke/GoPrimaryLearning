// Lissajous generates GIF animations of random Lissajous figures.
package main

import (
	"image"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

//color.RGBA{0xRR, 0xGG, 0xBB, 0xff}
//Invalid value 0xRR, 0XGG in Original code. Not a number literal.color.RGBA type has uint8. so has valid values 0 <= C <= A.
//var palette = []color.Color{color.White, color.RGBA{0, 255, 68, 255}}

const (
	whiteIndex = 0
	greenIndex = 1
)

func main() {

	rand.Seed(time.Now().UTC().UnixNano())
	lissajous(os.Stdout)

}

func lissajous15(out io.Writer) {
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
			img.SetColorIndex(size+int(size*x+0.5), size+int(size*y+0.5), greenIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)

}
