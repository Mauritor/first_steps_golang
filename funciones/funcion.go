package main

import (
	"fmt"
)

func main() {
	saludar()
	res := sumar(3, 5)
	fmt.Println(res)
}
func saludar() {
	saludo := "hello"
	fmt.Println(saludo)
}

func sumar(a, b int) int {
	return a + b
}
