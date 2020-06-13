//A web server used to generate fractal images for clients.
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"net/http"
	"strconv"
	"log"
)

func main(){
	const(
		width, height = 1024, 1024
	)

	params := map[string]float64{
		"ximin" : -2,
		"xmax" : 2,
		"ymin" : -2,
		"ymax" :2,
		"zoom" :1,
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		for name := range params {
			s := r.FormValue(name)
			if s == "" {
				continue
			}
			f, err := strconv.ParseFloat(s, 64)
			if err != nil {
				http.Error(w, fmt.Sprintf("query param %s : %s", name, err), http.StatusBadRequest)
				return
			} 
			params[name] = f
		}
		if params["xmax"] <= params["xmin"] || params["ymax"] <= params["ymin"] {
			http.Error(w, fmt.Sprintf("min coordinate greater than max"), http.StatusBadRequest)
			return 
		}

		xmin := params["xmin"]
		xmax := params["xmax"]
		ymin := params["ymin"]
		ymax := params["ymax"]
		zoom := params["zoom"]

		lenX := xmax - xmin
		midX := xmin + lenX/2
		xmin = midX - lenX/2/zoom
		xmax = midX + lenX/2/zoom
		lenY := ymax - ymin
		midY := ymin + lenY/2
		ymin = midY - lenY/2/zoom
		ymax = midY + lenY/2/zoom

		img := image.NewRGBA(image.Rect(0, 0, width, height))
		for py := 0; py < height; py++ {
			y := float64(py)/height * (ymax - ymin) + ymin
			for px := 0; px < width; px++ {
				x := float64(px)/width * (xmax -xmin) + xmin
				z := complex(x, y)
				img.Set(px, py, mandelbrot(z))
			}
		}
		err := png.Encode(w ,img)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	log.Fatal(http.ListenAndServe(":5520", nil))
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const constrast = 15

	var v complex128

	for n := uint8(0); n < iterations; n++ {
		v = v * v + z

		if cmplx.Abs(v) > 2 {
			return colorGet(n)
		}

	}
	return color.Black
}


func colorGet(n uint8) color.Color {
	colorSheet := []color.Color{
		color.RGBA{238, 132, 170, 255},
		color.RGBA{132, 112, 255, 255},
		color.RGBA{0, 191, 255, 255},
		color.RGBA{154, 205, 50, 255},   //	YellowGreen
		color.RGBA{255, 215, 0, 255},    // Gold
		color.RGBA{250, 128, 114, 255},  // Salmon
		color.RGBA{218, 112, 214, 255},  // Orchid
		color.RGBA{205, 201, 201, 255},  // Snow3
		color.RGBA{72, 118, 255, 255},   //	RoyalBlue1
		color.RGBA{99, 184, 255, 255},   //	SteelBlue1
		color.RGBA{171, 130, 255, 255},  // MediumPurple1
		color.RGBA{255, 69, 0, 255},     //LightPink1
		color.RGBA{238, 132, 170, 255},  //	OrangeRed1
		color.RGBA{255, 127, 0, 255},  // DarkOrange1
	}
	len := uint8(len(colorSheet))
	return colorSheet[n%len]
}