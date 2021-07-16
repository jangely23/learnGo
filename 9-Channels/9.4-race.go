//usar  range para recorrer  valoresde un canal y cerrar el canal una vez se hace

package main

import "fmt"

func main() {

	c := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c) //cierra los canales
	}()

	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("Finalizado.")
}
