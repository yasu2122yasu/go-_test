package main

import (
	"fmt"

	"test01/hello"
)

func main() {
	a := hello.GetHello("World!")
	fmt.Println(a)
}
