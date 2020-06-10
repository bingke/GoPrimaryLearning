//Use Newton's method to solve a complex equation, such as $z^4-1=0$. 
//The number of iterations from each starting point to four roots corresponds to the gray level of the shadow. 
//The point corresponding to the root of the equation is represented by color.

package main

import(
	"image/color"
	"image/png"
	"image"
	"math/cmplx"
	"math"
	"os"
)

type Func func(complex128) complex128

var colorPool = []color.RGBA{
	{30, 144, 255, 255},
	{255, 65, 0, 255},
	{138, 43, 226, 255},
	{25, 165,0, 255},
}

var chosenColors = map[complex128]color.RGBA{}

func main(){

	const(
		xmin, xmax, ymin, ymax = -2, 2, -2, 2
		width, height = 1024, 1024
	)
		
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < width; py++ {
		y := float64(py) / height * (ymax - ymin) + ymin
		for px :=0; px < height; px ++ {
			x := float64(px) / width * (xmax - xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, z4(z))
		}
	}
	png.Encode(os.Stdout, img)
}

func z4(z complex128) color.Color{
	f := func (z complex128) complex128{
		return z*z*z*z -1
	}
	fPrime := func(z complex128) complex128{
		return (z - 1 / (z*z*z)) / 4
	}
	return newton(z, f, fPrime) 
}

func newton(z complex128, f Func, fPrime Func ) color.Color{
	
	const (
		iterations = 40
		contrast = 20
	)

	for i := uint8(0); i < iterations; i++ {
		z -= fPrime(z)
		if cmplx.Abs(f(z)) < 1e-6 {
			root := complex(round(real(z), 4), round(imag(z),4))
			c, ok := chosenColors[root]
			if !ok {
				if len(colorPool) == 0 {
					panic("no color left")
				}
				c = colorPool[0]
				colorPool = colorPool[1: ]
				chosenColors[root] = c
			}
			y, cb, cr := color.RGBToYCbCr(c.R, c.G, c.B)
			scale := math.Log(float64(i)) / math.Log(iterations)
			y -= uint8(float64(y) * scale)
			return color.YCbCr{y, cb, cr}
		}
	}
	return color.Black
}

func round(f float64, digits int) float64{
	if math.Abs(f) < 0.5{
		return 0
	}
	pow := math.Pow10(digits)
	return math.Trunc(f*pow + math.Copysign(0.5, f)) / pow
}

