package main

import (
	"fmt"
	"math"
)

func main() {
	a := 0.5
	var dub float64 = 20
	var top float64 = 32
	b := a * 365
	c := b / top
	v := b / dub
	fmt.Println(b, math.Ceil(c), math.Ceil(v))
}
