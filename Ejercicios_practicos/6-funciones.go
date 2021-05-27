// crear una funcion anonima

package main

import "fmt"

func main() {
	nombre := "jessica"

	func(n string) {
		fmt.Println("Hola,", n)
	}(nombre)

}
