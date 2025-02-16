package main

import (
	"fmt"
	"math"
	"math/rand/v2"
)

//
// so many falses for some reason, there is a bug
//
// думаю нужно будет сделать так, чтобы там не
// поле создавалось, в которое кидается игла
// (думаю бак связан с этим), а просто использовалась
// чисто длина между линиями

func main() {
	// space between lines
	linesLen := 10.0
	// lenght of a needle
	needleLen := 4.0

	type Coordinates struct {
		X float64
		Y float64
	}

	// the size of the field on which we will throw a needle
	// (e.g. 100x100)
	var field Coordinates
	field.X = 100
	field.Y = 100

	for i := 0; i < 100; i++ {

		// throw a needle, determine coordinates of the middle of the needle
		var needleCoords Coordinates
		needleCoords.X = float64(rand.Int32N(int32(field.X))) + rand.Float64()
		needleCoords.Y = float64(rand.Int32N(int32(field.Y))) + rand.Float64()
		fmt.Println("needleCoords.Y: ", needleCoords.Y)

		// the angle under which the needle fell
		// it will be a cos between needle and perpendicular to the line
		needleAngle := math.Cos(rand.Float64()*2 - 1)
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
		needlePerp := needleAngle * (needleLen / 2)
		fmt.Println("needlePerp: ", needlePerp)

		fmt.Println(isCrossing(linesCoords, needlePerp, needleCoords.Y))
		fmt.Println("")
	}
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
