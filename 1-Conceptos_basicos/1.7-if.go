/* La instruccion if evalua si una condicion es verdadera o falsa y ejecuta el codigo solo cuando sea verdadera
La instruccion else se ejecuta cuando la validacion de el if es falsa y requerimos hacer alguna accion en ese caso particular.
La instruccion else if cuando el if inicial es falso permite realizar otra validacion adicional con ese resultado. Esto quiere decir que podemos validar mas de 2 posibles resultados antes de realizar alguna ejecucion.*/

package main

import "fmt"

func main() {
	grade := 60

	if grade >= 65 {
		fmt.Println("Passing grade")
	} else if grade >= 50 {
		fmt.Println("Grade recovery")
	} else {
		fmt.Println("Failing grade")
	}
}
