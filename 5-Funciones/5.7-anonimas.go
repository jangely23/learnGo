/*  estas funciones no tienen identificador, tambien se conocen como auto ejecutables */

package main

import "fmt"

func main() {
	func() {
		fmt.Println("La funcion se ejecuta sin parametros.")
	}() // ejecuta la funcion

	valor := 2000

	func(x int) { //recibe parametros
		fmt.Println("La funcion se ejecuta con X:", x)
	}(valor) // ejecuta la funcion con el argumento valor
}
