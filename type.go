package main

import "fmt"

type NewInt int
type IntAlias = int

func main() {
	var a NewInt
	fmt.Printf("a type: %T\n", a)

	var b IntAlias
	fmt.Printf("b type: %T\n", b)
}
