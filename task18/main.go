package main

import (
	"fmt"
	"math"
	"sync"
	"sync/atomic"
	"time"
)

type AtomicCounter struct {
	val int32 //will be enought for 2 billions iteration. If you need more you can use int64 instead, but it will be slower
}

func (ac *AtomicCounter) Add(increment int32) {
	atomic.AddInt32(&ac.val, int32(math.Max(0, float64(increment))))
}

func (ac *AtomicCounter) Value() int32 {
	return atomic.LoadInt32(&ac.val)
}



type Counter struct {
	val int
	m   sync.Mutex
}

func (c *Counter) Add(increment int) {
	c.m.Lock()
	defer c.m.Unlock()
	c.val += int(math.Max(0, float64(increment)))
}

func (c *Counter) Value() int {
	c.m.Lock()
	defer c.m.Unlock()
	return c.val
}

func main() {
	//Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде. По завершению программа должна выводить итоговое значение счетчика.
	counter := Counter{}
	atomicCounter := AtomicCounter{}

	var wg sync.WaitGroup
	for i := range 10 { //mutex counter
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			fmt.Println("Sleep in i goroutine by mutex counter", val)
			time.Sleep(1 * time.Second)
			counter.Add(val)
		}(i)
	}

	for i := range 10 { //atomic counter
		wg.Add(1)
		go func (val int32) {
			defer wg.Done()
			fmt.Println("Sleep in i goroutine by atomic counter", val)
			time.Sleep(1 * time.Second)
			atomicCounter.Add(val)
		}(int32(i))
	}

	wg.Wait()
	fmt.Println("Mutex counter ", counter.Value())
	fmt.Println("Atomic counter ", atomicCounter.Value())
}
