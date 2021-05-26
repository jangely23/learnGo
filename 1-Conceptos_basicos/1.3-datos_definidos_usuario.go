/* Para crear un nuevo tipo de variable se utiliza la palabra clave **type** seguido del nombre de la variable e indicando el tipo de dato subyacente por ejemplo: **type dinero int**

Una vez creado ya se puede utilizar dentro del programa como un tipo de dato, es importante aclarar que por mas de que el tipo subyacente sea un int este nunca sera igual al tipo dinero. */

package main

import (
	"fmt"
)

var a int

type dinero int

var b dinero

func main() {
	a = 42
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	b = 1000
	fmt.Println(b)
	fmt.Printf("%T\n", b)
}
