//canal con bufer con fallo

package main

import "fmt"

func main() {

	ca2 := make(chan int, 2)

	ca2 <- 63
	ca2 <- 42

	fmt.Println(<-ca2)
	fmt.Println(<-ca2)

}
