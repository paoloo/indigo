package main

import (
	"log"
	"time"
)

type Rtimer struct {
	Done     chan bool
	Waiting  int
	Callback func()
}

func (r *Rtimer) Init(waiting_time int, callback func()) {
	r.Waiting = waiting_time
	r.Callback = callback
	r.Done = make(chan bool)
	ticker := time.NewTicker(time.Duration(r.Waiting) * time.Millisecond)
	go func() {
		for {
			select {
			case <-r.Done:
				ticker.Stop()
				return
			case <-ticker.C:
				log.Printf("boo")
				r.Callback()
			}
		}
	}()
}

func (r *Rtimer) Stop() {
	r.Done <- true
}
