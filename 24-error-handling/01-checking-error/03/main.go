package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.Create("name.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	r := strings.NewReader("Tony Stark")
	if _, err := io.Copy(f, r); err != nil {
		log.Fatal(err)
	}
}
