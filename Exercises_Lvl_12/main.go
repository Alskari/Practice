package main

import (
	"fmt"
	//"woofwoof"
	//"github.com/GoesToEleven/go-programming/code_samples/008-ninja-level-twelve/01/dog"
	"github.com/Alskari/Practice/Exercises_Lvl_12/woofwoof"
	"github.com/GoesToEleven/go-programming/code_samples/010-ninja-level-thirteen/01/finished/dog"
)

// canine
type canine struct {
	name string
	age  int
}

func main() {
	fido := canine{"Hank", woofwoof.Years(7)}

	fmt.Println(fido)
	fmt.Println(dog.YearsTwo(8))
}
