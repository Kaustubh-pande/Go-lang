package main

import (
	"fmt"
	"strconv"
	"strings"
)

func timeConversion(s string) string {
	/*
	 * Write your code here.
	 */
	st := strings.Split(s, ":")
	hh, mm, ss := st[0], st[1], st[2]

	if strings.Contains(ss, "PM") {

		sso := strings.Split(ss, "PM")
		ss = sso[0]
		i1, err := strconv.Atoi(hh)

		if err == nil {
			if i1 != 12 {
				i1 = i1 + 12
			}
		}
		hh = strconv.Itoa(i1)

	} else if strings.Contains(ss, "AM") {
		sso := strings.Split(ss, "AM")
		ss = sso[0]
		if hh == "12" {
			hh = "00"
		}

	}
	newst := string(hh + ":" + mm + ":" + ss)
	return newst
}
func main() {
	fmt.Println(timeConversion("12:45:54PM"))
}

//str2 := "5678"
/** converting the str2 variable into an int using ParseInt method
fmt.Println("name is a type of: ", reflect.TypeOf(hh))
	fmt.Println("name is a type of: ", reflect.TypeOf(mm))
	fmt.Println("name is a type of: ", reflect.TypeOf(ss))


*/
