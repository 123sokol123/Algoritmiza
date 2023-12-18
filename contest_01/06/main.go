package main

import (
	"fmt"
)

const pit int = 5000
const cos int = 1000
const pot = 500
const dve int = 200
const sot int = 100

func main() {
	var a, b, c, d, e, f int
	fmt.Scan(&a)
	b = a / pit
	a = a - b*pit
	c = a / cos
	a = a - c*cos
	d = a / pot
	a = a - d*pot
	e = a / dve
	a = a - e*dve
	f = a / sot

	fmt.Println(b, c, d, e, f)
}