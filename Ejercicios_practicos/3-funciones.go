/* usar la funcion defer */

package main

import "fmt"

func main() {
	defer foo()
	fmt.Println("Hola, yo soy del main y me encuentro en la liena 2 el codigo")
	fmt.Println("y yo soy la ultima linea del main")
}

func foo() {
	defer func() {
		fmt.Println("Hola, soy la funcion con defer dentro foo")
	}()
	fmt.Println("Hola, soy la funcion foo")
}
