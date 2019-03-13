package main

import "fmt"

func gradingStudents(grades []int32) []int32 {
	/*
	 * Write your code here.
	 */
	n := grades[0]
	grades = grades[1:]

	var result []int32
	for _, v := range grades {
		fmt.Println(v)
		rem := v % 5
		fmt.Println("rem", rem)
		actual := 5 - rem
		fmt.Println("act", actual)
		if v < 38 {
			result = append(result, v)
		} else if v >= 38 && v <= 100 {
			if actual >= 3 {
				result = append(result, v)
			} else if rem == 0 {
				result = append(result, v)
			} else {
				result = append(result, (v + actual))

			}
		}

	}
	fmt.Println(n)
	fmt.Println(result)
	return result
}
func main() {
	var arr = []int32{4, 73, 46, 38, 80}
	gradingStudents(arr)
}
