/* Se usa para decirle al programa que ejecute esa funcion al final de la funcion en la que se declaro, ¿como de termina el sistema que llego al final? cuando realice un retorno (return), fuando finaliza todas las acciones y llega al final ( } ), o cuando una Go rutina esta en estado panic. */

package main

import (
	"fmt"
)

func main() {
	defer foo() //ejecuta al final esta funcion main
	bar()
}

func foo() {
	fmt.Println("Hola, mundo. Estoy en una funion.")
}

func bar() {
	fmt.Println("ahora estoy en la funion BAR.")
}
