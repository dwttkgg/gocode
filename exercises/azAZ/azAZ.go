package main

import "fmt"

func main() {
	// for i := 'a'; i <= 'z'; i++ {
	// 	// var mn int
	// 	// mn = int(i)
	// 	fmt.Printf("%c", i)
	// }
	var a int = 0
	var sum int = 0
	var b int = 0
	for i := 1; i < 101; i++ {
		for j := 1; j <= i; j++ {

			if i%j == 0 {
				a++

			}

		}
		if a == 2 {
			fmt.Printf("%v \t", i)
			b++
			sum = sum + i
		}
		if b == 5 {
			fmt.Println()
			b = 0
		}
		a = 0
	}
	fmt.Println(sum)
}
