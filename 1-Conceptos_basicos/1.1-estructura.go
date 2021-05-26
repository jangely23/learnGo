/* La estructura de un programa:
- package main
- declaracion import
- funcion main
- punto  de entrada del programa
- punto de salida, final del programa.

Todos los programas deben tener al menos un paquete llamado main, ejemplo deun programa: */

package main //Paquete principal o biblioteca estandar de Go

import (
	"fmt" //paquetes importados para su uso
)

func main() { // funcion principal main, inicia la ejecucion del codigo
	fmt.Println("Hola, mundo") //usa el paquete importado para mostrar un mensaje
}
