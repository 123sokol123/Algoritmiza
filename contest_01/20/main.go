package main

import (
	"fmt"
	"sort"
	"strings"
)

func sortString(s string) string {
	str := strings.Split(s, "")
	sort.Strings(str)
	return strings.Join(str, "")
}

func main() {
	var a, b string
	fmt.Scan(&a, &b)

	if sortString(a) == sortString(b) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
