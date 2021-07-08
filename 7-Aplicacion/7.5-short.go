/*
Este paquete permite ordenar Slice, tambien permit ordenar colecciones definidas por el usuario.

*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	x1 := []int{4, 7, 3, 42, 99, 18, 16, 56, 12}
	x2 := []string{"Jessia", "Q", "A", "Lauren"}

	fmt.Println(x1)
	fmt.Println(x2)

	sort.Ints(x1)    //ordena enteros de menor a mayor
	sort.Strings(x2) //ordena cadenas de menor a mayor

	fmt.Println(x1)
	fmt.Println(x2)

}
