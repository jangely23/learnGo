/* crear funcion que tome parametros variables tipo int
parar un slice como argumento y hacer el despliegue del int
retornar la suma de el slice.

crear otra funcion que reciba un parametro tipo slice int
retorne la suma de los valores pasados
*/

package main

import "fmt"

func main() {
	numero := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sumaUno := foo(numero...) //despliega el slice
	sumaDos := bar(numero)    //pasa el slice

	fmt.Println(sumaUno)
	fmt.Println(sumaDos)
}

func foo(n ...int) int {
	var total int
	for _, valor := range n {
		total += valor
	}
	return total
}

func bar(j []int) int {
	var total int
	for _, valor := range j {
		total += valor
	}
	return total
}
