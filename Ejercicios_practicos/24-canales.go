/*
haz que este codigo fucione:

func main() {
	cs := make(chan<- int)

	go func() {
		cs <- 42
	}()
	fmt.Println(<-cs)

	fmt.Printf("------\n")
	fmt.Printf("cs\t%T\n", cs)
}

*/

package main

import (
	"fmt"
)

func main() {
	cs := make(chan int) // se vuelve bidireccional

	go func() {
		cs <- 42
	}()
	fmt.Println(<-cs)

	fmt.Printf("------\n")
	fmt.Printf("cs\t%T\n", cs)
}
