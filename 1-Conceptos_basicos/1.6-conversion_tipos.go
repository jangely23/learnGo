/* Go tiene su propio lenguaje, es por ello que no se habla de objetos, si no, de crear tipos de datos y valores de tipos, igualmente se habla de conversión y aserción que permite cambiar el tipo de dato de una variable. */

package main

import (
	"fmt"
)

var a int

type dinero int //crea un tipo de dato dinero

var b dinero //crea una variable con el nuevo tipo de dato

func main() {
	a = 42
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	b = 1000
	fmt.Println(b)
	fmt.Printf("%T\n", b)

	a = int(b) //convierte a int un tipo dinero
	fmt.Println(a)
	fmt.Printf("%T\n", a)
}
