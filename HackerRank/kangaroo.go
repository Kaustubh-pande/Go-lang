package main

import "fmt"

func kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) string {

	var kangaroo1 int32 = x1
	var kangaroo2 int32 = x2
	var result string
	if (x2 > x1) && (v2 > v1) {
		result = "NO"
	} else {
		for i := 0; i <= 10000; i++ {
			kangaroo1 += v1
			kangaroo2 += v2
			if kangaroo1 == kangaroo2 {
				result = "YES"
				fmt.Println(i + 1)
				break
			} else {
				result = "NO"
			}
		}
	}
	return result
}
func main() {
	fmt.Println(kangaroo(0, 2, 5, 3))
}
