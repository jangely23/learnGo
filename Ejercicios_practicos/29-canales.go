// programa que lance 10 gorutinas, cada gorutina agrega 10 numeros a un canal
//extrae del canal e imprimelos

package main

import "fmt"

func main() {
	c := make(chan int)
	for j := 0; j < 10; j++ {
		go func() {
			for i := 0; i < 10; i++ {
				c <- i
			}
			//no se cierrra por que no requerimos cerrar el canal en cada iteracion, no se debe usar for range
		}()
	}

	for k := 0; k < 100; k++ { //se conoce el numero de iteracciones.
		fmt.Println(<-c)
	}
}
