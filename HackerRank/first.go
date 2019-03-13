package main

func miniMaxSum(arr []int32) int32 {
	sort.int32(arr)
	var (
		min int32 = 0
		max int32 = 0
	)
	xmin := arr[0 : len(arr)-1]
	xnax := arr[1:len(arr)]
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
	miniMaxSum(x)

}
