package main

import (
	"fmt"
	"sort"
)

func birthdayCakeCandles(ar []int32) int32 {
	var (
		n     int32
		max   int32
		total = 0
	)
	n = int32(len(ar))
	sort.Slice(ar, func(i, j int) bool {
		return ar[i] < ar[j]
	})
	max = ar[n-1]
	for _, value := range ar {
		if max == value {
			total++
		}
	}
	//fmt.println(n)
	return int32(total)
}

func main() {
	var arr = []int32{3, 2, 1, 3}
	fmt.Println(birthdayCakeCandles(arr))

}
