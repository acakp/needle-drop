package main

import (
	"fmt"
	"math"
	"math/rand/v2"
)

func main() {
	// space between lines
	linesLen := 100.0
	// lenght of a needle
	needleLen := 70.0

	type Coordinates struct {
		X float64
		Y float64
	}

	// the size of the field on which we will throw a needle
	// (e.g. 100x100)
	var field Coordinates
	field.X = 1000
	field.Y = 1000

	// the number of times when the needle crossed lines
	crossings := 0

	iterations := 1000000
	for i := 0; i < iterations; i++ {

		// throw a needle, determine coordinates of the middle of the needle
		var needleCoords Coordinates
		needleCoords.X = float64(rand.Int32N(int32(field.X))) + rand.Float64()
		needleCoords.Y = float64(rand.Int32N(int32(field.Y))) + rand.Float64()
		fmt.Println("needleCoords.Y: ", needleCoords.Y)

		// the angle under which the needle fell
		// gen angle in radians 0 - pi/2 (0 to 90 degrees)
		needleAngle := rand.Float64() * (math.Pi / 2)
		fmt.Println("needleAngle: ", needleAngle)

		// determine coordinates of the lines
		linesCount := int(field.X/linesLen) + 1
		linesCoords := make([]float64, linesCount)
		linesCoords[0] = 0
		for i := 1; i < linesCount; i++ {
			linesCoords[i] = linesCoords[i-1] + linesLen
		}

		// calculate lenght of pependicular of the line from middle of the needle
		// to end of the needle
		needlePerp := math.Cos(needleAngle) * (needleLen / 2)
		fmt.Println("needlePerp: ", needlePerp)

		isCr := isCrossing(linesCoords, needlePerp, needleCoords.Y)
		if isCr {
			crossings++
		}
		fmt.Println(isCr)
		fmt.Println("")
	}

	fmt.Println("")
	fmt.Println("number of iterations: ", iterations)
	fmt.Println("times the needle crossed lines: ", crossings)
	crRate := (float64(crossings)) / float64(iterations)
	fmt.Println("line crossing rate: ", crRate)

	pi := (2 * needleLen) / (crRate * linesLen)
	// the closer to pi, the more accurate the result.
	fmt.Println("pi = ", pi)
}

func isCrossing(lines []float64, needlePerp, needleCoordsY float64) bool {
	for i := 1; i < len(lines); i++ {
		fmt.Println("lines[i]:", lines[i], "needleCoordsY: ", needleCoordsY+needlePerp)
		if needleCoordsY < lines[i] {
			fmt.Println("needleCoordsY < lines[i]")
			if needleCoordsY+needlePerp > lines[i] || needleCoordsY-needlePerp < lines[i-1] {
				return true
			} else {
				return false
			}
		}
	}
	return false
}
