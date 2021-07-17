//es una herramienta para< cancelar gorutinas huerfanas o en fuga

package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background()) //devuele una funcion de cancelacion de un argumento con retorno contex background

	fmt.Println("chqueo de error 1:", ctx.Err()) //muestra el error
	fmt.Println("num gorutinas 1:", runtime.NumGoroutine())

	go func() {
		n := 0
		for {
			select {
			case <-ctx.Done(): //se usa en declaraciones select y solo tendra un valor cuando se cancele el contexto
				return
			default:
				n++
				time.Sleep(time.Millisecond * 200) //retarda 2 segundos
				fmt.Println("Trabajando", n)
			}
		}
	}()

	time.Sleep(time.Second * 2) // ejecta la funcion pricipal
	fmt.Println("chequeo de error:", ctx.Err())
	fmt.Println("num gorutinas 2:", runtime.NumGoroutine())

	fmt.Println("A punto de cancelar context.")
	cancel() //cancela el contexto
	fmt.Println("context cancelado.")

	time.Sleep(time.Second * 2)
	fmt.Println("chequeo de error 3:", ctx.Err())
	fmt.Println("num gorutinas 3:", runtime.NumGoroutine())
}
