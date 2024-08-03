package main

import (
	"fmt"
	"sync"
	"time"
)

type (
	Chopsticks struct {
		sync.Mutex
	}

	Guest struct {
		LChops *Chopsticks
		RChops *Chopsticks
		ID     int
		Times  int
	}
)

const (
	MAX_GUESTS        = 10
	MAX_TIMES_EAT     = 2
	MAX_CONCUR_GUESTS = 5
)

var wg sync.WaitGroup

func (g *Guest) eat(pickupChopSticks chan *Guest) {
	defer wg.Done()
	fmt.Printf("Guest %d waiting\n", g.ID)
	pickupChopSticks <- g
	g.LChops.Lock()
	g.RChops.Lock()
	fmt.Printf("Guest %d eating\t*\n", g.ID)
	g.Times++
	eatingTime()
	g.LChops.Unlock()
	g.RChops.Unlock()
	fmt.Printf("Guest %d finishing eating %d times\n", g.ID, g.Times)
	<-pickupChopSticks
}

func eatingTime() {
	time.Sleep(2 * time.Second)
}

func main() {
	chopsticks := make([]*Chopsticks, MAX_GUESTS)
	for i := 0; i < cap(chopsticks); i++ {
		chopsticks[i] = new(Chopsticks)
	}

	guests := make([]*Guest, MAX_GUESTS)
	for i := 0; i < cap(guests); i++ {
		guests[i] = &Guest{
			ID:     i + 1,
			Times:  0,
			LChops: chopsticks[i],
			RChops: chopsticks[(i+1)%MAX_GUESTS],
		}
	}

	pickUpChopsticksPermission := make(chan *Guest, MAX_CONCUR_GUESTS)

	for j := 0; j < MAX_TIMES_EAT; j++ {
		for i := 0; i < cap(guests); i++ {
			wg.Add(1)
			go guests[i].eat(pickUpChopsticksPermission)
		}

	}

	wg.Wait()

}
