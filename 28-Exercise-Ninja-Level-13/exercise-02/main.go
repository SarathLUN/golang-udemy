package main

import (
	"fmt"
	"github.com/SarathLUN/golang-udemy/28-Exercise-Ninja-Level-13/exercise-02/quote"
	"github.com/SarathLUN/golang-udemy/28-Exercise-Ninja-Level-13/exercise-02/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}
