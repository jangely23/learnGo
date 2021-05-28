/* Todo en Go es pasado por valor.
- se deben usa cuando tenemos alguna variable o valor que involucra una cantidad grande de datos
- cuando se quiere cambiar directamente un valor que esta almacenado en una direccion de memoria
*/
package main

import "fmt"

func main() {

	x := 0
	fmt.Println("sin puntero desde main antes: ", x)  //muestra el valor
	fmt.Println("con puntero desde main antes: ", &x) //mustera la direccion en memoria
	fmt.Println("")

	foo(&x) //paso una direccion de memoria como argumento

	fmt.Println("")
	fmt.Println("sin puntero desde main despues: ", x)
	fmt.Println("con puntero desde main despues: ", &x)

}

/*se pasa el puntero para poder modificar el valor de x en el espacio en memoria, de lo contrario no se modifica x,
lo que se hacia es pasar un valor, es decir pasar el valor de x*/

func foo(i *int) {

	fmt.Println("sin puntero desde foo() antes de: ", i)  //muestra el valor que es una direccion en memoria
	fmt.Println("con puntero desde foo() antes de: ", *i) //accede al valor de esa direccion
	fmt.Println("")

	*i = 42 //modifica el valor de esa direccion "de referencia"

	fmt.Println("sin puntero desde foo() despues de: ", i)
	fmt.Println("con puntero desde foo() despues de: ", *i)
}
