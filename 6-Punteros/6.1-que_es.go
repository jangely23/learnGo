/*Todo lo que se guarda en el codigo queda en un espacio en memoria
Un puntero es una direccion en memoria donde hay guardado un valor de cualquier tipo.
& ampersan permite ver la direccion en memoria

existen tambien un conjunto de metodos o method sets que:
deterina cuales metodos adjuntar a un tipo.
cuanto el receptro de la funcion no es puntero, puedo usar valores normales o punteros para llamar el metodo.
cuando es un receptro puntero solo puedo usar valores tipo puntero.

*/

package main

import "fmt"

func main() {
	a := 42
	fmt.Println(a)
	fmt.Println(&a) //muestra la direccion en memoria de la variable a en hexadecimal

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a) //muestra tipo de dato de la direccion en memoria, el asterisco indica que es un puntero hacia un lugar en memoria

	var b *int = &a // a la variable b que es de tipo puntero se le asigna el valor de la memoria de a
	fmt.Println(b)

	fmt.Println(*b) //accede el valor almacenado en la direccion de memoria a traves de "de referencia"

	*b = 100 //moficica el valor de el campo en memoria usando "de referencia"

	fmt.Println(*b)
	fmt.Println(a) //vale 100 ya que fue moficada desde b a traves del puntero

}
