/* ayuda a limitar el scope de una variable */

package main

import "fmt"

func main() {
	a := incrementor() //recibe el closure de el return
	b := incrementor()

	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a()) //referencia el mismo lugar de memoria en a
	fmt.Println(b())
	fmt.Println(b()) //referencia el nuevo lugar de memoria en b

}

func incrementor() func() int {
	var x int // variable dentro del scope delcarada en 0
	return func() int {
		x++ // realiza closure de la variable x y se lleva una copia del entorno de x
		return x
	}
}
