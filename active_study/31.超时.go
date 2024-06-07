package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

}
func ClassicContextTimeout(ctx context.Context) error {

	syncRet := make(chan bool)

	// here do bussiness process here
	go func() {
		time.Sleep(3 * time.Second) // wait 3 seconds
		syncRet <- true
	}()

	select {
	case <-ctx.Done(): // context canceled or timeout
		fmt.Println("context canceled and return.")
		return ctx.Err()
	case <-syncRet: // success return
		fmt.Println("process finished.")
		return nil
	}
}
