package main

import (
	"fmt"
)

// Complete the diagonalDifference function below.
func diagonalDifference(arr [][]int32) int32 {
	var (
		LR int32 = 0
		RL int32 = 0
		n  int32 = 0
	)
	n = len(arr)
	fmt.Println(n)
	/*
		for i, value := range arr {
			for j, value := range arr {
				if i == j {
					LR += value[i][j]
				}

			}
		}
		return LR
	*/
}
func main() {
	var arr = [][]int32{{0, 0}, {1, 2}, {2, 4}, {3, 6}, {4, 8}}
	fmt.Println(len(arr))
	//fmt.Println(diagonalDifference(arr))
}
