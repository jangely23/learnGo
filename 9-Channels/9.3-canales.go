package main

import "fmt"

func main() {
	c := make(chan int)

	go enviar(c)
	recibir(c) //No se usa una go rutina por que si o si requiere el valor del otro

	fmt.Println("Finalizado.")
}

func enviar(c chan<- int) {
	c <- 42
}

func recibir(c <-chan int) {
	fmt.Println(<-c)
}
