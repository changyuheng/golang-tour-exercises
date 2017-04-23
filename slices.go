// +build ignore

package main

import (
	"golang.org/x/tour/pic"
	"math/rand"
	"time"
)

func scaleFunc1(x, y int) uint8 {
	return uint8((x + y) / 2)
}

func scaleFunc2(x, y int) uint8 {
	return uint8(x * y)
}

func scaleFunc3(x, y int) uint8 {
	return uint8(x ^ y)
}

func scalePic(p [][]uint8, scaleFunc func(int, int) uint8) {
	for y, row := range p {
		for x := range row {
			row[x] = scaleFunc(x, y)
		}
	}
}

func Pic(dx, dy int) [][]uint8 {
	p := make([][]uint8, dy)
	for y := 0; y < dy; y++ {
		p[y] = make([]uint8, dx)
	}

	scaleFuncs := [3]func(int, int) uint8{}
	scaleFuncs[0] = scaleFunc1
	scaleFuncs[1] = scaleFunc2
	scaleFuncs[2] = scaleFunc3
	rand.Seed(int64(time.Now().Nanosecond())) // Doesn't work on tour engine
	scalePic(p, scaleFuncs[rand.Intn(3)])

	return p
}

func main() {
	pic.Show(Pic)
}
