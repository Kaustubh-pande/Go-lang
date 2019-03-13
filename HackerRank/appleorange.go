package main

import "fmt"

func countApplesAndOranges(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) (int32, int32) {
	//m := len(apples)
	//n := len(oranges)
	var applecount int32
	var orangecount int32
	for _, av := range apples {
		if (a+av) >= s && (a+av) <= t {
			applecount++
		}
	}
	for _, ov := range oranges {
		if (b+ov) >= s && (b+ov) <= t {
			orangecount++
		}
	}

	fmt.Println(applecount, orangecount)
	return applecount, orangecount
}
func main() {
	var a = []int32{-2, 2, 1}
	var o = []int32{5, -6}
	countApplesAndOranges(7, 11, 5, 15, a, o)
}
