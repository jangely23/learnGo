package main

import (
	"fmt"
)

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	fmt.Println(x[:5])  //imprime hasta el los primeros 5 valores
	fmt.Println(x[5:])  //muestra los ultimos 5 valores
	fmt.Println(x[2:7]) //imprime los valores de en medio
}
