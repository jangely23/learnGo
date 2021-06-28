/*Crea un valor y asigna una a una variable un valor de cualquier tipo,
luego vamo a imprimir la direccion de memoria donde esta almacenado ese valor.
*/

package main

import "fmt"

func main() {
	valor := 125

	fmt.Println(&valor)
}
