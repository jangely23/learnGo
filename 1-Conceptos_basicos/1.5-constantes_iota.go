/* iota : es un identificador que al ser declarado permite crear una sucesion de constantes con valor que corresponde al de un entero incremental pero sin tipo. */

package main

import "fmt"

const (
	a = iota
	b
	c
	d
)

func main() {
	fmt.Printf("variable de tipo %T de valor %v\n", a, a)
	fmt.Printf("variable de tipo %T de valor %v\n", b, b)
	fmt.Printf("variable de tipo %T de valor %v\n", c, c)
	fmt.Printf("variable de tipo %T de valor %v\n", d, d)

}
