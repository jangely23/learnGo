/* Estructura de una función: func (r receptor) identificador(parámetros) (retorno(s)) { Código } */

package main

import (
	"fmt"
)

func main() {
	total := suma(1, 2, 3, 4, 5, 6, 7, 8, 9) // llama la funcion y pasa multiples argumentos
	fmt.Printf("\n----------------------\nTipo de dato de total es: %T\n", total)
	fmt.Println("valor total de la suma es:", total)
}

func suma(x ...int) int { //funcion que recibe multiples parametros y retorna un entero

	fmt.Printf("%T\n\n", x)
	fmt.Printf("%v\n\n", x)
	result := 0

	for i, v := range x {
		result += v
		fmt.Printf("el indice %v suma a %v y el total es: %v\n", i, v, result)
	}

	fmt.Println("la resultado fue de:", result)

	return result
}
