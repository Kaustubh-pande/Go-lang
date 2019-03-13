package main

import (
	"fmt"
)

// Complete the plusMinus function below.
func plusMinus(arr []int32) {
	var (
		postive  float32 = 0
		negative float32 = 0
		zeros    float32 = 0
		length   float32 = 0
	)
	length = float32(len(arr))
	fmt.Println("length", length)
	for _, value := range arr {
		if value > 0 {
			postive += 1
		} else if value < 0 {
			negative += 1
		} else {
			zeros += 1
		}
	}
	fmt.Println("Positive ", postive)
	fmt.Println("negative ", negative)
	fmt.Println("zeros ", zeros)
	var postive1 = float32(postive / length)
	var negative1 = float32(negative / length)
	var zeros1 = float32(zeros / length)
	fmt.Println(postive1)
	fmt.Println(negative1)
	fmt.Println(zeros1)
}

//func main() {
//	var arr = []int32{-4, 3, -9, 0, 4, 1}
//	plusMinus(arr)
//}
