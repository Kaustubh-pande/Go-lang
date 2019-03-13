package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	csvfilename := flag.String("csv", "state and capital.csv", "A csv file in the format 'state,capital' ")
	flag.Parse()
	timelimit := flag.Int("limit", 30, " The time limit for the quiz in second")
	file, err := os.Open(*csvfilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open a file: %s ", *csvfilename))

	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Faild to parseing the file.")

	}
	problems := parseLines(lines)
	timer := time.NewTimer(time.Duration(*timelimit) * time.Second)
	correct := 0
problemloop:
	for i, q := range problems {
		fmt.Printf("Quiz #%d : %s => \n", i+1, q.state)
		anschan := make(chan string)
		go func() {
			var ans string
			fmt.Scanf("%s\n", &ans)
			anschan <- ans
		}()

		select {
		case <-timer.C:
			fmt.Printf("\nYour scored %d out of %d.", correct, len(problems))
			break problemloop
		case ans := <-anschan:

			if strings.ToLower(ans) == q.capital {
				correct++
				fmt.Println("correct!")
			}
		}

	}
	fmt.Printf("Your scored %d out of %d.\n", correct, len(problems))

}

type quiz struct {
	state   string
	capital string
}

func parseLines(lines [][]string) []quiz {
	ret := make([]quiz, len(lines))
	for i, line := range lines {
		ret[i] = quiz{
			state:   line[0], //struct assgin
			capital: strings.TrimSpace(strings.ToLower(line[1])),
		}
	}
	return ret
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
