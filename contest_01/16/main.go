package main

import "fmt"

func FormData(n int, data []int) []float64 {
	izmeren := make([]float64, n)

	izmeren[0] = float64(data[0])
	izmeren[n-1] = float64(data[n-1])

	for i := 1; i < n-1; i++ {
		izmeren[i] = (float64(data[i-1] + data[i] + data[i+1])) / 3.0
	}
	return izmeren
}

func main() {
	var n int
	fmt.Scan(&n)

	data := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&data[i])
	}

	smoothedData := FormData(n, data)

	for i := 0; i < n; i++ {
		fmt.Printf("%.15f ", smoothedData[i])
	}
	fmt.Println()
}
