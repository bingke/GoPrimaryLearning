// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	zmax, zmin := zminZmax()
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			a, ax, ay := corner(i+1, j)
			b, bx, by := corner(i, j)
			c, cx, cy := corner(i, j+1)
			d, dx, dy := corner(i+1, j+1)
			if a && b && c && d {
				fmt.Printf("<polygon style = 'stroke: %s;' points = '%g, %g, %g, %g, %g, %g, %g, %g'/>\n",
					color(i, j, zmax, zmin), ax, ay, bx, by, cx, cy, dx, dy)
			} else {
				continue
			}

		}
	}
	fmt.Println("</svg>")
}

func zminZmax() (float64, float64) {
	min := math.NaN()
	max := math.NaN()

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			for xoff := 0; xoff < 2; xoff++ {
				for yoff := 0; yoff < 2; yoff++ {
					x := xyrange * (float64(i+xoff)/cells - 0.5)
					y := xyrange * (float64(j+yoff)/cells - 0.5)
					z := f(x, y)
					if math.IsNaN(min) || z < min {
						min = z
					}
					if math.IsNaN(max) || z > max {
						max = z
					}
				}
			}

		}
	}
	return max, min
}

func color(i, j int, zmax, zmin float64) string {
	min := math.NaN()
	max := math.NaN()

	for xoff := 0; xoff < 2; xoff++ {
		for yoff := 0; yoff < 2; yoff++ {
			x := xyrange * (float64(i+xoff)/cells - 0.5)
			y := xyrange * (float64(j+yoff)/cells - 0.5)
			z := f(x, y)
			if math.IsNaN(min) || z < min {
				min = z
			}
			if math.IsNaN(max) || z > max {
				max = z
			}
		}
	}

	color := ""

	if math.Abs(max) > math.Abs(min) {
		red := math.Exp(math.Abs(max)) / math.Exp(math.Abs(zmax)) * 255
		if red > 255 {
			red = 255
		}
		color = fmt.Sprintf("#%02x0000", int(red))
	} else {
		blue := math.Exp(math.Abs(min)) / math.Exp(math.Abs(zmin)) * 255
		if blue > 255 {
			blue = 255
		}
		color = fmt.Sprintf("#0000%02x", int(blue))
	}

	return color

}

func corner(i, j int) (bool, float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale

	//Check whether the return value is a normal floating point number
	if math.IsNaN(sx) || math.IsInf(sx, 0) || math.IsNaN(sy) || math.IsInf(sy, 0) {
		return false, 0.0, 0.0
	} else {
		return true, sx, sy
	}

}

//draw the shape of egg box
func eggbox(x, y float64) float64 {
	r := 0.2 * (math.Cos(x) + math.Cos(y))
	return r
}

//draw the shape of Snowdrift
func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}
