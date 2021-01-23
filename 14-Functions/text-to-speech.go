package main

import (
	"log"
	"os/exec"
)

func main() {
	s := "Make the computer speak"
	cmd := exec.Command("espeak", s)
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
