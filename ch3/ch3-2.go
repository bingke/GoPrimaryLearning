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
				fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
					ax, ay, bx, by, cx, cy, dx, dy)
			} else {
				fmt.Errorf("corner() produces a non-numeric value")
			}

		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (bool, float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := eggbox(x, y)

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
