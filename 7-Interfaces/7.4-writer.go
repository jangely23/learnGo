/*
Paquete io.Writer, es la interfaz que envuelve el método de escritura básico
*/

package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Hello, playground")
	fmt.Fprintln(os.Stdout, "Hello, playground")   //implementa la interface writer
	io.WriteString(os.Stdout, "Hello, playground") //implementa la interface writer
}
