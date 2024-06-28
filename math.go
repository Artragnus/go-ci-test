package main

import "fmt"

func main() {
	fmt.Println(Soma(10, 10))
}

func Soma(a int, b int) int {
	return a + b
}

func Divisao(a int, b int) int {
	return a / b
}

func Multiplicao(a int, b int) int {
	return a * b
}

func Subtracao(a int, b int) int {
	return a - b
}

func Media(a int, b int) int {
	return (a + b) / 2
}

func Maior(a int, b int) int {
	if a > b {
		return a
	}
	return b
}