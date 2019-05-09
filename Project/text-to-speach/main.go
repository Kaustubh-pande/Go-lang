package main

import (
	"log"
	"os/exec"
)

func main() {
	s := "Make the computer speak, hello! how are you?"
	cmd := exec.Command("espeak", s)
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
	a := "hello azher what are you doing!"
	cmd = exec.Command("espeak", a)
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

	b := "Lets go to play carrom"
	cmd = exec.Command("espeak", b)

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
	c := "1234"
	cmd = exec.Command("espeak", c)

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
