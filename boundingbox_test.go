package goemetry

import (
	"os"
	"runtime"
	"strings"
	"testing"

	svg "github.com/ajstarks/svgo"
	"github.com/stretchr/testify/assert"
)

func TestIsAboveishSuperSimpleBaseCase(t *testing.T) {

	one := BoundingBox{
		BottomLeft: Point{
			X: 100,
			Y: 100,
		},
		Height: 10,
		Width:  10,
	}

	two := BoundingBox{
		BottomLeft: Point{
			X: 100,
			Y: 80,
		},
		Height: 10,
		Width:  10,
	}

	m := make(map[string]BoundingBox)

	m["one"] = one
	m["two"] = two

	drawSvg(m)

	assertAboveish(t, one, two)
}
func TestIsNotAboveishWithLargeVerticalGap(t *testing.T) {

	one := BoundingBox{
		BottomLeft: Point{
			X: 100,
			Y: 500,
		},
		Height: 10,
		Width:  10,
	}

	two := BoundingBox{
		BottomLeft: Point{
			X: 100,
			Y: 90,
		},
		Height: 10,
		Width:  10,
	}

	m := make(map[string]BoundingBox)

	m["one"] = one
	m["two"] = two

	drawSvg(m)

	assertNotAboveish(t, one, two)
}

func TestLol(t *testing.T) {

	one := BoundingBox{
		BottomLeft: Point{
			X: 100,
			Y: 500,
		},
		Height: 10,
		Width:  10,
	}

	two := BoundingBox{
		BottomLeft: Point{
			X: 100,
			Y: 90,
		},
		Height: 10,
		Width:  10,
	}

	m := make(map[string]BoundingBox)

	m["one"] = one
	m["two"] = two

	drawSvg(m)
}

func assertAboveish(t *testing.T, above BoundingBox, below BoundingBox) {
	assert.True(t, above.IsAboveish(below), "%+v IS expected to be below %+v", above, below)
}

func assertNotAboveish(t *testing.T, above BoundingBox, below BoundingBox) {
	assert.False(t, above.IsAboveish(below), "%+v is NOT expected to be below %+v", above, below)
}

func getCaller() string {
	// we get the callers as uintptrs - but we just need 1
	fpcs := make([]uintptr, 1)

	// skip 3 levels to get to the caller of whoever called Caller()
	n := runtime.Callers(3, fpcs)
	if n == 0 {
		panic("oops, no caller")
	}

	// get the info of the actual function that's in the pointer
	fun := runtime.FuncForPC(fpcs[0] - 1)
	if fun == nil {
		panic("oops, no caller")
	}

	// return its name
	parts := strings.Split(fun.Name(), "/")
	return parts[len(parts)-1]
}

func drawSvg(boxes map[string]BoundingBox) {

	if len(boxes) == 0 {
		return
	}

	file, err := os.Create(getCaller() + ".svg")
	if err != nil {
		panic(err)
	}

	width := 0
	height := 0

	for _, v := range boxes {

		vx := v.BottomLeft.X + int(v.Width)
		vy := v.BottomLeft.Y + int(v.Height)

		if width < vx {
			width = vx
		}
		if height < vy {
			height = vy
		}
	}

	width = int(1.5 * float64(width))
	height = int(1.5 * float64(height))

	canvas := svg.New(file)
	canvas.Start(width, height)

	for k, v := range boxes {
		canvas.Rect(v.BottomLeft.X, v.BottomLeft.Y, int(v.Width), int(v.Height))
		canvas.Text(v.BottomLeft.X, v.BottomLeft.Y, k)
	}

	canvas.End()
}
