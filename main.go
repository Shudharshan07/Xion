package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	data := NewData()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	for {
		select {
		case <-ctx.Done():
			return
		default:
		}
		time.Sleep(time.Second)
		data.UpdateData()
		data.Print()
		fmt.Print("\n")
	}
}
