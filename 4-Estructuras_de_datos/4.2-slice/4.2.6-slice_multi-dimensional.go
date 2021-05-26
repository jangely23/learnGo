package main

import (
	"fmt"
)

func main() {
	x1 := []string{"danelly", "analista de cuentas", "29"}
	x2 := []string{"jessica", "ingeniera", "31"}

	xx := [][]string{x1, x2} //crea un slice multi-dimencional
	fmt.Println(xx)

	for indice1, registro1 := range xx {
		fmt.Println("Registro:", indice1)

		for indice2, registro2 := range registro1 { //itera sobre el array que esta en el registro1
			fmt.Printf("\tIndice de posición: %v\tvalor: %v\n", indice2, registro2)
		}
	}
}
