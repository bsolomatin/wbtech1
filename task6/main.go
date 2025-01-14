package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
  //Реализовать все возможные способы остановки выполнения горутины.
	var wg sync.WaitGroup
	//By chan
	wg.Add(1)
	stopByFlagCh := make(chan bool)
	go stopByFlagWorker(stopByFlagCh, &wg)
	time.Sleep(3 * time.Second) // let work
	stopByFlagCh <- true
	//By context cancel()
	wg.Add(1)
	ctx, cancel := context.WithCancel(context.Background()) //or any other suitable context.
	go stopByCtxWorker(ctx, &wg)
	time.Sleep(3 * time.Second) // let work
	cancel()
	//By timer
	wg.Add(1)
	timer := time.NewTimer(3 * time.Second)
	go stopByTimeWorker(timer, &wg)
	time.Sleep(3 * time.Second) // let work
	//By graceful shutdown
	wg.Add(1)
	graceSh := make(chan os.Signal, 1)
	signal.Notify(graceSh, os.Interrupt, syscall.SIGTERM)
	go stopByGraceShWorker(graceSh, &wg)
	//By panic
	wg.Add(1)
	go stopByPanicWorker(&wg)
	time.Sleep(1 * time.Second) //let work

	wg.Wait()
}

func stopByFlagWorker(ch chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ch:
			fmt.Println("Worker has been stopped by getting flag")
			close(ch)
			return
		default:
			fmt.Println("Worker continues work, flag still no received")
			time.Sleep(1 * time.Second)
		}
	}
}

func stopByCtxWorker(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Worker has been stopped by cancel context")
			return
		default:
			fmt.Println("Worker continues work, context still no cancelled")
			time.Sleep(1 * time.Second)
		}
	}
}

func stopByTimeWorker(timer *time.Timer, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-timer.C:
			fmt.Println("Worker has been stopped by timer ending")
			return
		default:
			fmt.Println("Worker continues work, timer still tick")
			time.Sleep(1 * time.Second)
		}
	}
}

func stopByGraceShWorker(ch chan os.Signal, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case sig := <-ch:
			fmt.Println("Received signal for interrupt/termination: ", sig)
			return
		default:
			fmt.Println("Worker continues work, waiting for interrupt/termination signal")
			time.Sleep(1 * time.Second)
		}
	}
}

func stopByPanicWorker(wg *sync.WaitGroup) {
	defer wg.Done()
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Panic %s happened, force recover for shutdown\n", r)
		}
	}()

	panic ("Stop goroutine by panic it is very, very bad idea. Do not reproduce this on prod")
}
