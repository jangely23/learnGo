/*
Requiere se indique una variable de inicio, una condicion a evaluar para ejecutar el codigo y que se incremente la variable inicio en cada iteracion

Cláusula_for = [Inicialización ";" Condición ";" Incremento]
*/

package main

import "fmt"

func main() {
	for i := 0; i < 4; i++ {
		fmt.Println("El valor es:", i)
	}
}
