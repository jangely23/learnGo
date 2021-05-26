/*  son funciones que se llaman a si mismo, puede traer problemas de memoria */

package main

import "fmt"

func main() {

	n1 := factorial(4)
	fmt.Println(n1)
}

func factorial(n int) int { // factorial es la multiplicacion de todos los numeros regresivos de x valor
	if n == 0 {
		return 1
	}
	return n * factorial(n-1) // recursividad
}
