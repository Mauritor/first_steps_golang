package main

import (
	"fmt"
	"strconv"
)

func main() {
	edad := 35
	num := "35"

	edadStr := strconv.Itoa(edad)
	numInt, _ := strconv.Atoi(num)

	fmt.Println("Mi edad es", edadStr)
	fmt.Println(numInt + 5)
}
