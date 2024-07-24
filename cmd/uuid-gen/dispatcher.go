package main

import (
	"fmt"
	"strconv"
)

type WorkerID int

// Dispatcher has a buffer channel to
type Dispatcher struct {
	WorkerPool chan WorkerID
	jobCh      chan Job
	doneCh     chan struct{}
	WorkerCh   []chan Job
}

func NewDispatcher(jobCh chan Job, doneCh chan struct{}) *Dispatcher {
	workers, _ := strconv.Atoi(maxWorkers)
	d := &Dispatcher{
		//buffer channel to read the jobs from
		jobCh:      jobCh,
		WorkerPool: make(chan WorkerID, workers),
		doneCh:     doneCh,
		WorkerCh:   make([]chan Job, workers),
	}

	for i := 0; i < workers; i++ {
		d.WorkerCh[i] = make(chan Job)
	}

	return d
}

func (d *Dispatcher) Start() {
	go func() {
		for {
			fmt.Println("Starting dispatcher")
			select {
			case <-d.doneCh:
				for _, wCh := range d.WorkerCh {
					close(wCh)
				}
				return
			case job := <-d.jobCh:
				wid := <-d.WorkerPool
				d.WorkerCh[wid] <- job
			}
		}
	}()
}

func (d *Dispatcher) Stop() {
	d.doneCh <- struct{}{}
}
