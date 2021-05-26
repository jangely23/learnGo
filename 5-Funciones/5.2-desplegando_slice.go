/* Este proceso permite pasar cada uno de los valores de un slice para usarlos de forma individual, esto se hace usando el operador ... cuando se pasa el argumento a una función, importante: el operador ... debe usarse en la función como ultimo parametro, por tanto como ultimo argumento al llamar la función. */

package main

import (
	"fmt"
)

func main() {
	x1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9} //slice

	total := suma(x1...) // llama la funcion y pasa un slice desplegado individualmente

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
