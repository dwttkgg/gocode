package main

import "fmt"

func maopaot(arr *[13]int) {
	c := *arr

	for i := 0; i < len(c); i++ {
		for j := 0; j < len(c)-i-1; j++ {
			if c[j] > c[j+1] {
				c[j+1] = c[j+1] + c[j]
				c[j] = c[j+1] - c[j]
				c[j+1] = c[j+1] - c[j]
			}
		}
	}
	fmt.Println(c)

}
func main() {
	var arr = [13]int{-4, 5, -2, 3, 4, 8, 9, -2, -11, 33, 22, -22, -11}
	maopaot(&arr)

}
