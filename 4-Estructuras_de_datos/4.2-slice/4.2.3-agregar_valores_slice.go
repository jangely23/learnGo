package main

import (
	"fmt"
)

func main() {
	y := []int{56, 57, 58, 59, 60}
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(x)

	x = append(x, 52) //agrega 52
	fmt.Println(x)

	x = append(x, 53, 54, 55) // agrega en una declaracion varios valores
	fmt.Println(x)

	x = append(x, y...) //agrega a x todos los elementos de y
	fmt.Println(x)
}
