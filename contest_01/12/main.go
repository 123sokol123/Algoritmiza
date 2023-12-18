package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)
	var chetchik int
	for chetchik = 0; x != 1; chetchik++ {
		if x%2 == 0 {
			x = x / 2
		} else {
			x = 3*x + 1
		}
	}
	fmt.Println(chetchik)

}
