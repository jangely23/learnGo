package main

import (
	"fmt"
)

func main() {

	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(x)

	y := append(x[:3], x[5:]...) //borra el elemento 45 y 46 del slice
	fmt.Println(y)

}
