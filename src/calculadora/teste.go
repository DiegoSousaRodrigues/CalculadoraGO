package main

import (
	"fmt"
	"os"
)

var n1 float64
var n2 float64
var op int

func getN1() float64 {
	fmt.Println("Digite um numero: ")
	fmt.Scan(&n1)
	return n1
}

func getN2() float64 {
	fmt.Println("Digite outro numero: ")
	fmt.Scan(&n2)
	return n2
}

func somar() float64 {
	return float64(getN1() + getN2())
}

func sub() float64 {
	return float64(getN1() - getN2())
}

func mult() float64 {
	return float64(getN1() * getN2())
}

func div() float64 {
	n1 = getN1()

	n2 = getN2()

	if n2 == 0 {
		os.Exit(0)
	}
	return n1 / n2
}

func main() {

	for {
		fmt.Println("Digite sua opção: \n1-Adição\n2-Subtração\n3-Multiplicação\n4-Divisao")
		fmt.Scan(&op)

		if op == 1 {
			fmt.Println("Resultado :", somar())
		}

		if op == 2 {
			fmt.Println("Resultado :", sub())
		}

		if op == 3 {
			fmt.Println("Resultado :", mult())
		}

		if op == 4 {
			fmt.Println("Resultado :", div())
		}

		if !(op < 1 || op > 4) {
			break
		}
	}
}
