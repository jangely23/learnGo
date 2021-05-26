/* Crear una funcion que retorne un int y otra que retorne un int y un string, mostrar resultados del llamado */

package main

import "fmt"

func main() {
	x := foo()
	y, z := bar() //retorna 2 valores

	fmt.Println(x)
	fmt.Println(y, z)
}

func foo() int {
	return 1000
}

func bar() (int, string) {
	return 1, "funciones"
}
