/*
Extraer los valores del siguiente codigo usando select:

func main() {
	salir := make(chan int)
	c := gen(salir)

	recibir(c, salir)

	fmt.Println("A punto de finalizar.")
}

func gen(salir <-chan int) <-chan int {
	c := make(chan int)

	for i := 0; i < 100; i++ {
		c <- i
	}

	return c
}
*/
package main

import (
	"fmt"
)

func main() {
	salir := make(chan int)
	c := gen(salir)

	recibir(c, salir)

	fmt.Println("A punto de finalizar.")
}

func gen(salir chan<- int) <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		salir <- 1
		close(c)
	}()

	return c
}

func recibir(c, salir <-chan int) {

	for {
		select {
		case <-salir:
			return
		case v := <-c:
			fmt.Println("desde el canal ", v)
			return
		}
	}
}
