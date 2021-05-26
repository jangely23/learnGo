/* PrintF formatea de acuerdo a la especificación de formato, es especificador de formato se utilizan los verbos que pueden ser generales y se expresan con % ejemplo %v. */

package main

import (
	"fmt"
)

var a int
var b string = "programa"

func main() {

	fmt.Println(a)
	fmt.Println(b)
	fmt.Printf("El valor de la variable a es: %v \n", a) //%v formatea el valor de a variable y agregar manualmente el salto de linea
	fmt.Printf("El valor de la variable b es: %v \n", b)
	fmt.Printf("El tipo de a es: %T \n", a) //%T formatea para imprimir el tipo de dato
	fmt.Printf("El tipo de b es: %T \n", b)

}
