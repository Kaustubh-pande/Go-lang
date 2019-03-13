package main

import "fmt"

func breakingRecords(scores []int32) []int32 {
	//n := scores[0]
	scores = scores[1:]
	fmt.Println(scores)
	maxscore := scores[0]
	minscore := scores[0]
	fmt.Println("max", maxscore)
	fmt.Println("min", minscore)
	var best int32
	var worst int32
	var result []int32
	for _, v := range scores {
		if v > maxscore {
			maxscore = v
			best++
		} else if v < minscore {
			minscore = v
			worst++
		}
		fmt.Println("max", maxscore)
		fmt.Println("min", minscore)
	}
	result = append(result, best, worst)
	fmt.Println(best, worst)
	return result
}
func main() {
	var arr = []int32{9, 3, 4, 21, 36, 10, 28, 35, 5, 24, 42}
	breakingRecords(arr)
}
