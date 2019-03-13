package main

import (
	"fmt"
	"math"
)

// Complete the diagonalDifference function below.
//
///*
func diagonalDifference1(arr [][]int32, n1 int, n2 int) int32 {
	var (
		LR int32 = 0
		RL int32 = 0
	)
	//var n int32 = 0
	//n = len(arr)
	//fmt.Println(n1)

	for i := 0; i < n1; i++ {
		for j := 0; j < n2; j++ {
			if i == j {
				LR += arr[i][j]
			}
			if i == n1-j-1 {
				RL += arr[i][j]
			}

		}
	}
	return int32(math.Abs(float64(LR - RL)))
	/*
		for i, value := range arr {
			for j, value := range arr {
				if i == j {
					LR += value[i][j]
				}
				if i == n1-j-1 {
					RL += value[i][j]
				}

			}
		}
	*/

} //*/
func main() {
	var arr2 = [][]int32{{1, 0, 5}, {1, 2, 4}, {2, 4, 3}}
	var n12 int = 0
	var n13 int = 0
	for _, h := range arr2 {
		n13 = len(h)
	}
	n12 = len(arr2)
	//fmt.Println(n12)
	//fmt.Println(n13)
	fmt.Println(diagonalDifference1(arr2, n12, n13))
}
