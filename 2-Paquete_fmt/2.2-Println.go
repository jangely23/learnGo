/* Esta función println agrega de forma automática los espacios entre variables concatenadas, es decir da formato al texto, y agrega un salto de linea a cada ejecucion a continuación un programa de ejemplo con su respectiva salida. */

package main

import (
	"fmt"
)

var a int
var b string = "programa"
var c bool
var d bool = true

func main() {
	e := 42
	f := "dice hola mundo."
	g := `El doctor dice que comer vegetales es
		 saludable.`
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println("mi", b, f)
	fmt.Println(g)
}
