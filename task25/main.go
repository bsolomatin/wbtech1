package main

import (
	"context"
	"fmt"
	"time"
)

type Sleeper interface {
	Sleep(duration time.Duration)
}

type TimeAfter struct{}

func (ta *TimeAfter) Sleep(duration time.Duration) {
	<-time.After(duration)
}

type TickerSleep struct{}

func (ts *TickerSleep) Sleep(duration time.Duration) {
	ticker := time.NewTicker(duration)
	defer ticker.Stop()
	<-ticker.C
}

type TimerSleep struct{}

func (ts *TimerSleep) Sleep(duration time.Duration) {
	timer := time.NewTimer(duration)
	<-timer.C
}

type ContextSleep struct{}

func (cs *ContextSleep) Sleep(duration time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()
	<-ctx.Done()
}

type BusyWaitSleep struct{}

func (bw *BusyWaitSleep) Sleep(duration time.Duration) {
	start := time.Now()
	for { //strongly do not recommended, due to useless highload resources
		if time.Since(start) >= duration {
			break
		}
	}
}

func main() {
	//Реализовать собственную функцию sleep.
	sleepers := []Sleeper{
		&TimeAfter{},
		&TickerSleep{},
		&TimerSleep{},
		&ContextSleep{},
		&BusyWaitSleep{},
	}

	for _, sleeper := range sleepers {
		fmt.Printf("Time now %s. Preparing for sleep. Planned sleeping time - 3 second\n",
			time.Now().Format("15:04:05 02-01-2006"))
		sleeper.Sleep(3 * time.Second)
		fmt.Printf("Time after sleeping %s\n", time.Now().Format("15:04:05 02-01-2006"))
	}
}
