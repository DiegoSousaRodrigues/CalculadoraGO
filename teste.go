package main

import (
	"fmt"
)

var n1 int
var n2 int

func getN1() int {
	fmt.Println("Informe o primeiro numero: ")
	fmt.Scanln(&n1)
	return n1
}

func getN2() int {
	fmt.Println("Informe o primeiro numero: ")
	fmt.Scanln(&n2)
	return n2
}

func adicao() {
	fmt.Println(getN1() + getN2())
}

func main() {
	adicao()
}
