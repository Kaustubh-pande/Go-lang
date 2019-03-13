package main

import (
	"fmt"
	"reflect"
)

// func distinct(arr2 []int32) []int32 {
// 	var result []int32
// 	//var count1 int32
// 	var prev int32
// 	var last int32
// 	var next int32
// 	//var countlast int32 = 0
// 	n := len(arr2) - 1
// 	fmt.Println("n", n)
// 	last = arr2[n]
// 	fmt.Println("last", last)
// 	for i, v2 := range arr2 {
// 		//fmt.Println(i)
// 		//	fmt.Println("countlast", countlast)
// 		if n == i {
// 			if next == last {
// 				prev = v2
// 				result = append(result, prev)
// 			}
// 			fmt.Println("break", i)
// 			break
// 		} else {
// 			//curr = arr2[count1]
// 			for j:=1;j<n;j++{
// 				next = arr2[i+1]

// 			}
// 			//next = arr2[i+1]

// 			if v2 == next {

// 				continue
// 			} else {
// 				prev = v2
// 				result = append(result, prev)
// 				//count1 = 0
// 			}
// 			//	fmt.Println("countlast2", countlast)
// 		}

// 		//countlast++
// 	}
// 	return result
// }

func in_array(val interface{}, array interface{}) (exists bool, index int) {
	exists = false
	index = -1
	var result []int32
	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
				index = i
				exists = true
				return
			} else {
				result = append(result, array[i])
			}
		}
	}

	return false, result
}

func migratoryBirds(arr1 []int32) []int32 {
	var count int32
	var result []int32
	//arrayOfMaps := make([]map[int]int, 5)

	arr1 = distinct(arr1)
	fmt.Println("arr1", arr1)
	for _, num := range arr1 {
		count = 0
		//arrayOfMaps[num] = arrayOfMaps[num] + 1
		for _, v := range arr1 {
			if num == v {
				count++
			}
		}
		result = append(result, count)

	}
	return result
}
func main() {

	ints := []int{1, 4, 3, 2, 6}
	for i := 0; i < len(ints); i++ {
		fmt.Println(in_array(ints[i], ints))
	}
	var arr = []int32{1, 4, 4, 4, 5, 3, 1, 3, 4, 6}
	fmt.Println(migratoryBirds(arr))
}
