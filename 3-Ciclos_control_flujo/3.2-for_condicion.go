/* Cuando se especifica una condicion esta se ejecutara siempre que la condicion sea verdadera, siempre debe generar dentro del bloque de codigo el incremento de la variable que se evalua. */

package main

import "fmt"

func main() {
	i := 0

	for i < 4 { //condicion a evaluar
		fmt.Println("El valor es:", i)
		i++ //incremento
	}
}
