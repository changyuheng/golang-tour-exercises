// +build ignore

package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) (z float64) {
	z = x
	for {
		t := z
		z = z - (z*z-x)/(z*2)
		if math.Abs(z-t) < 1e-9 {
			break
		}
	}
	return
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}
