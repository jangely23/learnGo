//es una herramienta para< cancelar gorutinas huerfanas o en fuga

package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	fmt.Println("chqueo de error 1:", ctx.Err())
	fmt.Println("num gorutinas 1:", runtime.NumGoroutine())

	go func() {
		n := 0
		for {
			select {
			case <-ctx.Done():
				return
			default:
				n++
				time.Sleep(time.Millisecond * 200)
				fmt.Println("Trabajando", n)
			}
		}
	}()

	time.Sleep(time.Second * 2)
	fmt.Println("chequeo de error:", ctx.Err())
	fmt.Println("num gorutinas 2:", runtime.NumGoroutine())

	fmt.Println("A punto de cancelar context.")
	cancel()
	fmt.Println("context cancelado.")

	time.Sleep(time.Second * 2)
	fmt.Println("chequeo de error 3:", ctx.Err())
	fmt.Println("num gorutinas 3:", runtime.NumGoroutine())
}
