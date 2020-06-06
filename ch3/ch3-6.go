//Exercise 3.6: Upsampling can reduce the effect of each pixel on calculating color 
//values and average values. The simple method is to divide each pixel into four sub-pixels 
//and implement it.

package main

import(
	"image"
	"os"
	"math/cmplx"
	"image/color"
	"image/png"
)

func main(){
	const(
		xmin, xmax, ymin, ymax = -2, 2, -2, 2
		width, height = 1024, 1024
		espX =  (xmax - xmin) / width
		espY = (ymax - ymin) / height
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < width; py++ {
		y := float64(py) / height * (ymax - ymin) + ymin
		for px :=0; px < height; px++ {
			x := float64(px) / width * (xmax - xmin) + xmin
			subColor := make([]color.Color, 0)
			for i :=0; i < 2; i++ {
				for j :=0 ; j < 2; j++ {
					z := complex(x + float64(espX * i), y + float64(espY * i))
					subColor = append(subColor, mandelbrot(z))
				} 
			}
			img.Set(px, py, avgColor(subColor))
		}
	}
	png.Encode(os.Stdout, img)
}

func mandelbrot(z complex128) color.Color{
	const iterations = 200
	const contrast = 15

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

func avgColor(colors []color.Color) color.Color {
	len := len(colors)
	var r, g, b, a uint8
	for _, co := range colors {
		r_, g_, b_, a_ := co.RGBA()
		r += uint8(r_ / uint32(len))
		g += uint8(g_ / uint32(len))
		b += uint8(b_ / uint32(len))
		a += uint8(a_ / uint32(len))
	}
	return color.RGBA{r, g, b, a}
}