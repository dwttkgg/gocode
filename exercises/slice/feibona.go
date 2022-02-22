package main

import "fmt"

func feibon(n int) []float64 {
	slice := make([]float64, n)
	slice[0] = 1
	slice[1] = 1
	for i := 2; i < n; i++ {
		slice[i] = slice[i-1] + slice[i-2]
	}
	return slice
}
func main() {

	fmt.Println(feibon(10))
}
