package main

import (
	"fmt"
	"woofwoof"
)

// canine
type canine struct {
	name string
	age  int
}

func main() {
	fido := canine{"Hank", woofwoof.Years(7)}

	fmt.Println(fido)
}
