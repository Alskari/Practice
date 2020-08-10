package main

import (
	"fmt"
	//"woofwoof"
	//"github.com/GoesToEleven/go-programming/code_samples/008-ninja-level-twelve/01/dog"
	"github.com/Alskari/Practice/Exercises_Lvl_12/woofwoof"
)

// canine
type canine struct {
	name string
	age  int
}

func main() {
	fido := canine{"Hank", woofwoof.Years(7)}

	fmt.Println(fido)
	fmt.Println(woofwoof.YearsTwo(8))
}
