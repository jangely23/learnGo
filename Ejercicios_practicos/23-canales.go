/*haz que este codigo funcione, con buffer y funciones literales:
func main() {
	c := make(chan int)

	c <- 42

	fmt.Println(<-c)
}
*/

package main

import (
	"fmt"
)

func main() {
	c := make(chan int)     //canal de tipo entero
	c2 := make(chan int, 1) //canal con 1 bufer

	go func() { // funcion anonima autoejecutable (funcion literal)
		c <- 42
	}()

	c2 <- 42 //envia valor al bufer

	fmt.Println(<-c)
	fmt.Println(<-c2)
}
