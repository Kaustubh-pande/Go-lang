package main

import (
	"fmt"
	"strings"
)

func staircase(n int) {

	for i := 1; i <= n; i++ {
		// if i%n == 0 {
		// 	fmt.Printf("#")
		// } else {
		// 	fmt.Printf(" ")
		// }
		// fmt.Println()
		for j := n; j >= i; j-- {
			if i == j {
				fmt.Print(strings.Repeat("#", i))
			} else {
				fmt.Print(" ") //strings.Repeat("*", i))
			}
		}
		fmt.Println()
	}

}
func main() {
	var n int = 6
	staircase(n)
}
