package main

import "fmt"

func countApplesAndOranges(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) {
	m := len(apples)
	n := len(oranges)
	fmt.Println(m, n)
	return ""
}
func main() {
	var a = []int32{-2, 2, 1}
	var o = []int32{5, -6}
	countApplesAndOranges(7, 11, 5, 15, a, o)
}
