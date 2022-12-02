// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"math"
	"net/http"
	"log"
	"strconv"
	"io"
)

var (
	width, height float64 // = 600.0, 320.0 // canvas size in pixels
	cells = 100.0 // number of grid cells
	xyrange = 30.0 // axis ranges (-xyrange..+xyrange)
	xyscale float64 // = width / 2 / xyrange // pixels per x or y unit
	zscale float64 // = height * 0.4 // pixels per z unit
	angle = math.Pi / 6 // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/svg+xml");

		width, _ = strconv.ParseFloat(r.URL.Query().Get("width"), 64)
		height, _ = strconv.ParseFloat(r.URL.Query().Get("height"), 64)
		xyscale = width / 2 / xyrange
		zscale = height * 0.4

		color := r.URL.Query().Get("color")

		createSvg(w, width, height, color)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func createSvg(w io.Writer, width, height float64, color string) {
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
	"style='stroke: grey; fill: white; stroke-width: 0.7' "+
	"width='%v' height='%v' fill='%v'>", width, height, color)

	for i := 0; i < int(cells); i++ {
		for j := 0; j < int(cells); j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g' style='fill:%s'/>\n",
			ax, ay, bx, by, cx, cy, dx, dy, color)
		}
	}
	fmt.Fprintf(w, "</svg>")
}

func corner(i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2.0 + (x-y)*cos30*xyscale
	sy := height/2.0 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	result := math.Sin(r) / r
	if math.IsNaN(result) { return 0 }
	return result
}

