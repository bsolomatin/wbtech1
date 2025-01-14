package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	/*
		Реализовать постоянную запись данных в канал (главный поток). 
		Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout. 
		Необходима возможность выбора количества воркеров при старте.
		Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.
		go run main.go -workersCnt=5

		1)Минималистично и идиоматично просто закрываем канал -> не требуется дополнительных каналов или контекстов для завершения работы воркеров
		2)Сочетается с sync.WaitGroup (При закрытии канала делаем wg.Done())
	*/
	workersCntPtr := flag.Int("workersCnt", 3, "Number of workers") // Например: go run main.go -workersCnt=5
	flag.Parse()
	workersCnt := *workersCntPtr
	if workersCnt <= 0 {
		fmt.Println("Workers count should be positive")
		return
	}

	var wg sync.WaitGroup
	ch := make(chan string)
	graceSh := make(chan os.Signal, 1)
	signal.Notify(graceSh, os.Interrupt, syscall.SIGTERM) //subscription to interrupt signal sending (by ctrl+c) to start graceful shutdown
	for i := 0; i < workersCnt; i++ {
		wg.Add(1)
		go Process(&wg, ch, i)
	}

	go func() {
		defer func() {
			if rec := recover();  rec != nil {
				fmt.Println("Panic: ", rec)
			}
		}()

		for {
			select {
			case <-graceSh: //get interrupt signal by ctrl+c command
				fmt.Println("Interrupt signal has been received, graceful shutdown process starts")
				close(ch) //close channel to stop all workers
				return
			case ch <- fmt.Sprintf("%d", rand.Intn(15)+1): //success send into the channel
				time.Sleep(1 * time.Second)
			default:
				fmt.Println("Channel has been closed or overflowed, stop work")
				return
			}
		}
	}()

	wg.Wait() //waiting for ending all workers
	fmt.Println("Program has been ended sucessfully when all workers complete their tasks")
}

func Process(wg *sync.WaitGroup, ch <-chan string, id int) {
	defer wg.Done() //decrease counter after ending job
	for {
		select {
		case val, ok := <-ch: //fetch data from channel if possible
			if !ok {
				fmt.Println("Channel has been closed")
				return
			}
			fmt.Printf("Worker %d get %s\n", id, val)
		}
	}
}
