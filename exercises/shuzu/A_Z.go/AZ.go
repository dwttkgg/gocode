package main

import (
	"fmt"
)

func main() {

	var a byte = 'A'
	var AZ [26]byte
	for i, _ := range AZ {
		AZ[i] = a + byte(i)
	}
	fmt.Printf("%c", AZ)

}
