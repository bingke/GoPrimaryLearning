//Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main
import(
	"image"
	"os"
	"math/cmplx"
	"image/color"
	"image/png"
	"math/rand"
    "time"
)

func main(){
	const(
		xmin, xmax, ymin, ymax = -2, 2, -2, 2
		width, height = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < width; py++ {
		y := float64(py) / height * (ymax - ymin) + ymin
		for px :=0; px < height; px++ {
			x := float64(px) / width * (xmax - xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, mandelbrot(z))
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
			return color.Gray{255 - contrast*n}
		}
	}
	rand.Seed(time.Now().Unix())
	c := rand.Int()

	return colorRandom(c)
}

func colorRandom(n int) color.Color {

	paletted := [16]color.Color{
		color.RGBA{66, 30, 15, 255},    // # brown 3
		color.RGBA{25, 7, 26, 255},     // # dark violett
		color.RGBA{9, 1, 47, 255},      //# darkest blue
		color.RGBA{4, 4, 73, 255},      //# blue 5
		color.RGBA{0, 7, 100, 255},     //# blue 4
		color.RGBA{12, 44, 138, 255},   //# blue 3
		color.RGBA{24, 82, 177, 255},   //# blue 2
		color.RGBA{57, 125, 209, 255},  //# blue 1
		color.RGBA{134, 181, 229, 255}, // # blue 0
		color.RGBA{211, 236, 248, 255}, // # lightest blue
		color.RGBA{241, 233, 191, 255}, // # lightest yellow
		color.RGBA{248, 201, 95, 255},  // # light yellow
		color.RGBA{255, 170, 0, 255},   // # dirty yellow
		color.RGBA{204, 128, 0, 255},   // # brown 0
		color.RGBA{43, 4, 7, 255},
		color.RGBA{12, 52, 9, 255},
	}
	return paletted[n%16]
}
