package main

import (
	"github.com/ajstarks/svgo"
	"os"
)

func main() {
	width := 500
	height := 500
	canvas := svg.New(os.Stdout)
	canvas.Start(width, height)

	lg := []svg.Offcolor{
		{10, "black", 1},
		{20, "gray", 1},
		{100, "lightgray", 1},
	}
	canvas.Def()
	canvas.LinearGradient("lg", 0, 100, 0, 0, lg)
	canvas.DefEnd()

	x := []int{50, 100, 150, 250, 350}
	y := []int{50, 100, 50, 65, 40}

	canvas.Polyline(x, y,
		`fill="url(#lg)"`,
		`stroke-linejoin="round"`,
		`stroke-linecap="round"`,
		`stroke-width="2"`,
		`stroke="#28AF47"`,
	)

	canvas.End()
}
