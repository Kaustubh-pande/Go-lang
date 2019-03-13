package main

//sort int32
/*
sort.Slice(arr, func(i, j int) bool {
        return arr[i] < arr[j]
    })
*/

import (
	"fmt"
	"sort"
)

func miniMaxsum(arr []int32) (int32, int32) {
	//sort.Sort(arr) //Ints(arr)

	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	var (
		min int32 = 0
		max int32 = 0
	)
	xmin := arr[0 : len(arr)-1]
	xmax := arr[1:len(arr)]
	for _, value := range xmin {
		min += value
	}
	for _, value := range xmax {
		max += value
	}
	return min, max

}

func main() {
	/*
							fmt.Println("Enter a F value : ")
							var input_f float32
							fmt.Scanf("%f", &input_f)
							fmt.Println("Conveting into C ...")
							var C float32
							C = (input_f - 32) * 5 / 9
							fmt.Println(" C value is : ", C)


						fmt.Println("Enter feet value : ")
						var (
							feet   float32
							meater float32
						)
						fmt.Scanf("%f", &feet)
						meater = feet * 0.3048
						fmt.Println("Meater value is ", meater, "m")


					for i := 1; i <= 100; i++ {
						if i%3 == 0 {
							fmt.Printf("%d  ", i)
						}
					}

				for i := 1; i <= 100; i++ {
					if i%3 == 0 {
						fmt.Println("Fizz", i)
					}
					if i%5 == 0 {
						fmt.Println("buzz", i)
					}
					if (i%3 == 0) && (i%5 == 0) {
						fmt.Println("FizzBuzz", i)
					}
				}

			fmt.Println(make([]int, 3, 9))
			x := [6]string{"a", "b", "c", "d", "e", "f"}
			fmt.Println(x[2:4])

		// min no in array
		x := []int{
			48, 96, 86, 68,
			57, 82, 63, 70,
			37, 34, 83, 27,
			19, 97, 9, 17,
		}

		min := x[0]
		for _, value := range x {

			if value < min {
				min = value
			}
		}
		fmt.Println("min value is : ", min)
	*/
	x := []int32{1, 3, 5, 7, 9}
	fmt.Println(miniMaxsum(x))

}

/* //// runniug
func getTotal(arr []int64) int64 {
    var total int64
    for _, item := range arr {
        total += item
    }
    return total
    }
// Complete the miniMaxSum function below.
func miniMaxSum(arr []int64) {

    sort.Slice(arr, func(i, j int) bool {
        return arr[i] < arr[j]
    })
    var total int64 = getTotal(arr)
    var max int64 = arr[0]
    var min int64 = total
    for _, item := range arr {
        var minusItem int64 = total - item
        if minusItem > max {
            max = minusItem
        }
        if minusItem < min {
            min = minusItem
        }
    }
    fmt.Println(min, max)

}
*/
/*
package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
    "sort"

)

// Complete the miniMaxSum function below.(int,int)
func miniMaxSum(arr []int32){
    //sort.Ints(arr)
    sort.Slice(arr, func(i, j int) bool {
        return arr[i] < arr[j]
    })
    var (
        min int32 = 0
        max int32 = 0
    )
    xmin := arr[0 : len(arr)-1]
    xmax := arr[1:len(arr)]
    for _, value := range xmin {
        min += value
    }
    for _, value := range xmax {
        max += value
    }
    s := string(min)
    d := string(max)

    //println(s,d)
    //return min, max
    fmt.println(s,"",d)
    //return s, d
    //println("%d",min)
    //println("%d",max)

}



func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    arrTemp := strings.Split(readLine(reader), " ")

    var arr []int32

    for i := 0; i < 5; i++ {
        arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
        checkError(err)
        arrItem := int32(arrItemTemp)
        arr = append(arr, arrItem)
    }

    miniMaxSum(arr)
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}

*/
