/* Las funciones son tipos, ciudadanos de primer orden, se pueden asignar a una variable */

package main

import "fmt"

func main() {

	prueba := func() { //la expresion de prueba es una funcion.
		fmt.Println("estoy ejecutandome desde la variable")
	}

	prueba() // ejecuta la funcion que se le asgino.

	prueba2 := func(z int) { //la expresion de prueba es una funcion con parametros.
		fmt.Println("tengo el parameto año:", z)
	}

	prueba2(2000)
}
