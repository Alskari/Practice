package main

import (
	"fmt"

	"github.com/Alskari/Practice/Exercises_Lvl_13/quote"
	"github.com/Alskari/Practice/Exercises_Lvl_13/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}
