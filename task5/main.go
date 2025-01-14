package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	//Разработать программу, которая будет последовательно отправлять значения в канал, а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.
	//go run main.go -ttl=10
	ttl := flag.Int("ttl", 10, "Number of seconds to shutdown")
	flag.Parse()
	var wg sync.WaitGroup
	ch := make(chan string)
	graceSh := make(chan os.Signal, 1)
	signal.Notify(graceSh, os.Interrupt, syscall.SIGTERM)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(*ttl)*time.Second)
	defer cancel()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for msg := range ch {
			fmt.Println("Received:", msg)
		}
	}()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Context expired, sending termination signal...")
			close(ch)
			wg.Wait()
			return
		case <-graceSh:
			fmt.Println("Received termination signal, closing channel...")
			close(ch)
			wg.Wait()
			return
		default:
			ch <- "ping"
			time.Sleep(1 * time.Second)
		}
	}

}
